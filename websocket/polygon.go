package polygonws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gorilla/websocket"
	"github.com/polygon-io/client-go/websocket/models"
)

// todo: in general, successful calls should be debug and unknown messages should be info
// todo: probably remove some junk logging before release too

type set map[string]struct{}

type Client struct {
	apiKey        string
	feed          Feed
	market        Market
	subscriptions map[string]set

	backoff backoff.BackOff

	ctx    context.Context
	cancel context.CancelFunc

	conn   *websocket.Conn
	rQueue chan []byte
	wQueue chan []byte

	parseData bool
	output    chan any

	log Logger
}

func New(ctx context.Context, config Config) (*Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if config.APIKey == "" {
		return nil, errors.New("API key is required")
	}

	if config.Log == nil {
		config.Log = &nopLogger{}
	}

	tctx, tcancel := context.WithCancel(context.Background())

	c := &Client{
		apiKey:        config.APIKey,
		feed:          config.Feed,
		market:        config.Market,
		subscriptions: make(map[string]set),
		backoff:       backoff.NewExponentialBackOff(),
		ctx:           tctx,
		cancel:        tcancel,
		rQueue:        make(chan []byte, 10000),
		wQueue:        make(chan []byte, 100),
		parseData:     config.ParseData,
		output:        make(chan any, 100000),
		log:           config.Log,
	}

	c.backoff = backoff.WithContext(backoff.NewExponentialBackOff(), ctx)
	c.backoff = backoff.WithMaxRetries(c.backoff, 25) // todo: let user configure?

	return c, nil
}

func (c *Client) Connect() error {
	if c.conn != nil {
		return nil
	}

	notify := func(err error, _ time.Duration) {
		c.log.Errorf(err.Error())
	}
	if err := backoff.RetryNotify(c.connect, c.backoff, notify); err != nil {
		return err
	}

	return nil
}

func (c *Client) connect() error {
	// todo: is this default dialer sufficient? might want to let user pass in a context so they can cancel the dial
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

	c.wQueue = make(chan []byte, 100)
	if err := c.authenticate(); err != nil {
		return fmt.Errorf("failed to write auth message: %w", err)
	}
	c.pushSubscriptions()

	go c.read()
	go c.write()
	go c.process()

	c.log.Debugf("connected successfully")

	return nil
}

func (c *Client) Subscribe(topic Topic, tickers ...string) error {
	params, err := getParams(c.market, topic, tickers...)
	if err != nil {
		return err
	}

	c.setSubscriptions(topic, tickers...)

	subscribe, err := json.Marshal(&models.ControlMessage{
		Action: models.Subscribe,
		Params: params,
	})
	if err != nil {
		return err
	}

	c.wQueue <- subscribe
	return nil
}

func (c *Client) Unsubscribe(topic Topic, tickers ...string) error {
	params, err := getParams(c.market, topic, tickers...)
	if err != nil {
		return err
	}

	c.deleteSubscriptions(topic, tickers...)

	unsubscribe, err := json.Marshal(&models.ControlMessage{
		Action: models.Unsubscribe,
		Params: params,
	})
	if err != nil {
		return err
	}

	c.wQueue <- unsubscribe
	return nil
}

func (c *Client) Output() any {
	return <-c.output
}

func (c *Client) Close() {
	if c.conn == nil {
		return
	}
	c.cancel()
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

func (c *Client) read() {
	defer func() {
		c.log.Debugf("closing read thread")
	}()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					break
				} else if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
					c.log.Errorf("connection closed unexpectedly: %v", err)
					break
				}
				c.log.Errorf("failed to read message: %v", err)
				break
			}
			c.rQueue <- msg
		}
	}
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.log.Debugf("closing write thread")
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case <-c.ctx.Done():
			err := c.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(writeWait))
			if err != nil {
				c.log.Errorf("failed to gracefully close: %v", err)
				return
			}
			c.log.Debugf("connection closed successfully")
			return
		case <-ticker.C:
			err := c.conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(writeWait))
			if err != nil {
				c.log.Errorf("failed to send ping message: %v", err)
				return
			}
		case msg := <-c.wQueue:
			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				c.log.Errorf("failed to set write deadline: %v", err)
				return
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				c.log.Errorf("failed to send message: %v", err)
				return
			}
		}
	}
}

func (c *Client) process() {
	defer func() {
		c.log.Debugf("closing process thread")
		close(c.output)
	}()

	for {
		select {
		case <-c.ctx.Done():
			return
		case data := <-c.rQueue:
			var msgs []json.RawMessage
			if err := json.Unmarshal(data, &msgs); err != nil {
				c.log.Errorf("failed to process raw messages: %v", err)
				continue
			}
			c.route(msgs)
		}
	}
}

func (c *Client) route(msgs []json.RawMessage) {
	for _, msg := range msgs {
		var ev models.EventType
		err := json.Unmarshal(msg, &ev)
		if err != nil {
			c.log.Errorf("failed to process message: %v", err)
			return
		}

		switch ev.EventType { // todo: enum?
		case "status":
			c.handleStatus(msg)
		default:
			c.handleData(ev.EventType, msg)
		}
	}
}

func (c *Client) handleStatus(msg json.RawMessage) {
	var cm models.ControlMessage
	if err := json.Unmarshal(msg, &cm); err != nil {
		c.log.Errorf("failed to unmarshal message: %v", err)
		return
	}

	switch cm.Status {
	case "connected":
		c.log.Debugf("connection successful")
	case "auth_success":
		c.log.Debugf("authentication successful")
	case "auth_failed":
		c.log.Errorf("authentication failed, closing connection")
		// todo: this is a fatal error so need to cancel any reconnects
		c.cancel()
		return
	case "success":
		c.log.Debugf("received a successful status message: %v", cm.Message)
	case "error":
		c.log.Errorf("received an error status message: %v", cm.Message)
	default:
		c.log.Infof("unknown status message '%v': %v", cm.Status, cm.Message)
	}
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

func supportsTopic(market Market, topic Topic) bool {
	switch market {
	case Stocks:
		return topic > stocksMin && topic < stocksMax
	case Options:
		return topic > optionsMin && topic < optionsMax
	case Forex:
		return topic > forexMin && topic < forexMax
	case Crypto:
		return topic > cryptoMin && topic < cryptoMax
	}
	return false
}

func getParams(market Market, topic Topic, tickers ...string) (string, error) {
	if !supportsTopic(market, topic) {
		return "", fmt.Errorf("topic '%v' not supported for feed '%v'", topic.prefix(), market)
	}

	if len(tickers) == 0 {
		return topic.prefix() + ".*", nil
	}

	var params []string
	for _, ticker := range tickers {
		params = append(params, topic.prefix()+"."+ticker)
	}

	return strings.Join(params, ","), nil
}

func (c *Client) setSubscriptions(topic Topic, tickers ...string) {
	for _, t := range tickers {
		_, exists := c.subscriptions[topic.prefix()]
		if !exists || t == "*" {
			c.subscriptions[topic.prefix()] = make(set)
		}
		c.subscriptions[topic.prefix()][t] = struct{}{}
	}
}

func (c *Client) pushSubscriptions() {
	for prefix, tickers := range c.subscriptions {
		var params []string
		for ticker, _ := range tickers {
			params = append(params, prefix+"."+ticker)
		}

		subscribe, err := json.Marshal(&models.ControlMessage{
			Action: models.Subscribe,
			Params: strings.Join(params, ","),
		})
		if err != nil {
			c.log.Errorf("failed to build subscription: %v", err)
			continue
		}

		c.wQueue <- subscribe
	}
}

func (c *Client) deleteSubscriptions(topic Topic, tickers ...string) {
	for _, t := range tickers {
		if _, prefixExists := c.subscriptions[topic.prefix()]; !prefixExists {
			c.subscriptions[topic.prefix()] = make(set)
		}
		if _, tickerExists := c.subscriptions[topic.prefix()][t]; !tickerExists {
			c.log.Infof("already unsubscribed to this ticker")
		}
		delete(c.subscriptions[topic.prefix()], t)
	}
}
