package polygonws

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gorilla/websocket"
	"github.com/polygon-io/client-go/websocket/models"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"gopkg.in/tomb.v2"
)

// todo: in general, successful calls should be debug and unknown messages should be info
// todo: probably remove some junk logging before release too

// Client defines a client to the Polygon WebSocket API.
type Client struct {
	apiKey string
	feed   Feed
	market Market

	shouldClose bool
	backoff     backoff.BackOff

	mtx    sync.Mutex
	rwtomb tomb.Tomb
	ptomb  tomb.Tomb

	conn   *websocket.Conn
	rQueue chan json.RawMessage
	wQueue chan json.RawMessage
	subs   subscriptions

	parseData bool
	output    chan any
	// todo: maybe add an error channel to signal fatal errors

	log Logger
}

// New creates a client for the Polygon WebSocket API.
func New(config Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, errors.New("API key is required")
	}

	if config.Log == nil {
		config.Log = &nopLogger{}
	}

	c := &Client{
		apiKey:    config.APIKey,
		feed:      config.Feed,
		market:    config.Market,
		backoff:   backoff.NewExponentialBackOff(),
		rQueue:    make(chan json.RawMessage, 10000),
		wQueue:    make(chan json.RawMessage, 1000),
		subs:      make(subscriptions),
		parseData: config.ParseData,
		output:    make(chan any, 100000),
		log:       config.Log,
	}

	c.backoff = backoff.WithMaxRetries(c.backoff, 25) // todo: let user configure?

	return c, nil
}

func (c *Client) Connect() error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.conn != nil {
		return nil
	}

	notify := func(err error, _ time.Duration) {
		c.log.Errorf(err.Error())
	}
	if err := backoff.RetryNotify(c.connect(false), c.backoff, notify); err != nil {
		return err
	}

	return nil
}

func (c *Client) connect(reconnect bool) func() error {
	return func() error {
		url := fmt.Sprintf("wss://%v.polygon.io/%v", string(c.feed), string(c.market))
		conn, res, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			return fmt.Errorf("failed to dial server: %w", err)
		} else if res.StatusCode != 101 {
			return errors.New("server failed to switch protocols")
		}

		conn.SetReadLimit(maxMessageSize)
		if err := conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			return fmt.Errorf("failed to set read deadline: %w", err)
		}
		conn.SetPongHandler(func(string) error {
			return conn.SetReadDeadline(time.Now().Add(pongWait))
		})
		c.conn = conn

		// reset write queue and push auth/subscription messages
		c.wQueue = make(chan json.RawMessage, 1000)
		if err := c.authenticate(); err != nil {
			return err
		}
		subs := c.subs.get()
		for _, msg := range subs {
			c.wQueue <- msg
		}

		c.rwtomb = tomb.Tomb{}
		c.rwtomb.Go(c.read)
		c.rwtomb.Go(c.write)

		if !reconnect {
			c.ptomb = tomb.Tomb{}
			c.ptomb.Go(c.process)
		}

		return nil
	}
}

func (c *Client) reconnect() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.shouldClose {
		return
	}

	c.rwtomb.Kill(nil)
	if err := c.rwtomb.Wait(); err != nil {
		c.log.Errorf("r/w threads closed: %v", err)
	}

	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}

	c.log.Debugf("disconnected unexpectedly, reconnecting")

	notify := func(err error, _ time.Duration) {
		c.log.Errorf(err.Error())
	}
	err := backoff.RetryNotify(c.connect(true), c.backoff, notify)
	if err != nil {
		c.log.Errorf("error reconnecting, closing connection")

		c.ptomb.Kill(nil)
		if err := c.ptomb.Wait(); err != nil {
			c.log.Errorf("process thread closed: %v", err)
		}

		close(c.output)
	}
}

// Subscribe sends a subscription message for a topic and set of tickers. If no
// tickers are passed, it will subscribe to all tickers for a given topic.
func (c *Client) Subscribe(topic Topic, tickers ...string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if !c.market.supports(topic) {
		return fmt.Errorf("topic '%v' not supported for cluster '%v'", topic.prefix(), c.market)
	}

	if len(tickers) == 0 || slices.Contains(tickers, "*") {
		tickers = []string{"*"}
	}

	subscribe, err := getSub(models.Subscribe, topic, tickers...)
	if err != nil {
		return err
	}

	c.subs.add(topic, tickers...)
	c.wQueue <- subscribe

	return nil
}

// Unsubscribe sends a message to unsubscribe from a topic and set of tickers. If no
// tickers are passed, it will unsubscribe from all tickers for a given topic.
func (c *Client) Unsubscribe(topic Topic, tickers ...string) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if !c.market.supports(topic) {
		return fmt.Errorf("topic '%v' not supported for cluster '%v'", topic.prefix(), c.market)
	}

	if len(tickers) == 0 || slices.Contains(tickers, "*") {
		tickers = maps.Keys(c.subs[topic])
	}

	unsubscribe, err := getSub(models.Unsubscribe, topic, tickers...)
	if err != nil {
		return err
	}

	c.subs.delete(topic, tickers...)
	c.wQueue <- unsubscribe

	return nil
}

func (c *Client) Output() any {
	return <-c.output
}

func (c *Client) Close() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.conn == nil {
		return
	}

	c.shouldClose = true

	c.rwtomb.Kill(nil)
	if err := c.rwtomb.Wait(); err != nil {
		c.log.Errorf("r/w threads closed: %v", err)
	}

	c.ptomb.Kill(nil)
	if err := c.ptomb.Wait(); err != nil {
		c.log.Errorf("process thread closed: %v", err)
	}

	close(c.output)

	c.conn.Close()
	c.conn = nil
}

func (c *Client) authenticate() error {
	b, err := json.Marshal(models.ControlMessage{
		Action: "auth",
		Params: c.apiKey,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal auth message: %w", err)
	}

	c.wQueue <- b
	return nil
}

func (c *Client) read() error {
	defer func() {
		c.log.Debugf("read thread closed")
	}()

	for {
		select {
		case <-c.rwtomb.Dying():
			return nil
		default:
			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					return nil
				} else if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
					return fmt.Errorf("connection closed unexpectedly: %w", err)
				}
				return fmt.Errorf("failed to read message: %w", err)
			}
			c.rQueue <- msg
		}
	}
}

func (c *Client) write() error {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.log.Debugf("write thread closed")
		ticker.Stop()
		go c.reconnect()
	}()

	for {
		select {
		case <-c.rwtomb.Dying():
			if err := c.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(writeWait)); err != nil {
				return fmt.Errorf("failed to gracefully close: %w", err)
			}
			return nil
		case <-ticker.C:
			if err := c.conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeWait)); err != nil {
				return fmt.Errorf("failed to send ping message: %w", err)
			}
		case msg := <-c.wQueue:
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				return fmt.Errorf("failed to set write deadline: %w", err)
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return fmt.Errorf("failed to send message: %w", err)
			}
		default:
			continue
		}
	}
}

func (c *Client) process() (err error) {
	defer func() {
		c.log.Debugf("process thread closed")
		if err != nil {
			go c.Close() // this client should close if it hits a fatal error (e.g. auth failed)
		}
	}()

	for {
		select {
		case <-c.ptomb.Dying():
			return nil
		case data := <-c.rQueue:
			var msgs []json.RawMessage
			if err := json.Unmarshal(data, &msgs); err != nil {
				c.log.Errorf("failed to process raw messages: %v", err)
				continue
			}
			if err := c.route(msgs); err != nil {
				return err
			}
		default:
			continue
		}
	}
}

func (c *Client) route(msgs []json.RawMessage) error {
	for _, msg := range msgs {
		var ev models.EventType
		err := json.Unmarshal(msg, &ev)
		if err != nil {
			c.log.Errorf("failed to process message: %v", err)
			continue
		}

		switch ev.EventType { // todo: enum?
		case "status":
			if err := c.handleStatus(msg); err != nil {
				return err
			}
		default:
			c.handleData(ev.EventType, msg)
		}
	}

	return nil
}

func (c *Client) handleStatus(msg json.RawMessage) error {
	var cm models.ControlMessage
	if err := json.Unmarshal(msg, &cm); err != nil {
		c.log.Errorf("failed to unmarshal message: %v", err)
		return nil
	}

	switch cm.Status {
	case "connected":
		c.log.Debugf("connection successful")
	case "auth_success":
		c.log.Debugf("authentication successful")
	case "auth_failed":
		// this is a fatal error so need to close the connection
		return errors.New("authentication failed, closing connection")
	case "success":
		c.log.Debugf("received a successful status message: %v", cm.Message)
	case "error":
		c.log.Errorf("received an error status message: %v", cm.Message)
	default:
		c.log.Infof("unknown status message '%v': %v", cm.Status, cm.Message)
	}

	return nil
}

func (c *Client) handleData(eventType string, msg json.RawMessage) {
	if !c.parseData {
		c.output <- msg // push raw data to output channel
		return
	}

	switch eventType {
	case "A":
		var out models.EquityAgg
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "AM":
		var out models.EquityAgg
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "CA":
		var out models.CurrencyAgg
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "XA":
		var out models.CurrencyAgg
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "T":
		var out models.EquityTrade
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "XT":
		var out models.CryptoTrade
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "Q":
		var out models.EquityQuote
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "C":
		var out models.ForexQuote
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "XQ":
		var out models.CryptoQuote
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "NOI":
		var out models.Imbalance
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "LULD":
		var out models.LimitUpLimitDown
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	case "XL2":
		var out models.Level2Book
		if err := json.Unmarshal(msg, &out); err != nil {
			c.log.Errorf("failed to unmarshal message: %v", err)
			return
		}
		c.output <- out
	default:
		c.log.Infof("unknown message type '%v'", eventType)
	}
}
