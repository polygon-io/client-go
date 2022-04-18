package polygonws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/polygon-io/client-go/websocket/models"
)

// todo: add reconnect logic
// todo: in general, successful calls should be debug and unknown messages should be info

type set map[string]struct{}

type Client struct {
	apiKey        string
	feed          Feed
	market        Market
	subscriptions map[string]set

	ctx    context.Context
	cancel context.CancelFunc

	conn   *websocket.Conn
	rQueue chan []byte
	wQueue chan []byte

	log Logger
}

func New(config Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, errors.New("API key is required")
	}

	if config.Log == nil {
		config.Log = &nopLogger{}
	}

	ctx, cancel := context.WithCancel(context.Background())

	c := &Client{
		apiKey:        config.APIKey,
		feed:          config.Feed,
		market:        config.Market,
		subscriptions: make(map[string]set),
		ctx:           ctx,
		cancel:        cancel,
		rQueue:        make(chan []byte, 10000),
		wQueue:        make(chan []byte, 100),
		log:           config.Log,
	}

	// push an auth message to the write queue
	if err := c.authenticate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Connect() error {
	if c.conn != nil {
		return nil
	}

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

	// todo: on reconnect, need to clear the write queue and push an auth message to the front
	//       this is a potential data race, might need to stop and restart write thread beforehand

	go c.read()
	go c.write()
	go c.process()

	return nil
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

// todo: maybe strip "." from tickers?
func getParams(market Market, topic Topic, tickers ...string) (string, error) {
	if !supportsTopic(market, topic) {
		return "", fmt.Errorf("topic '%v' not supported for feed '%v'", topic, market)
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

func (c *Client) setSubscription(prefix string, ticker string) {
	_, exists := c.subscriptions[prefix]
	if !exists || ticker == "*" {
		c.subscriptions[prefix] = make(set)
	}
	c.subscriptions[prefix][ticker] = struct{}{}
}

func (c *Client) deleteSubscription(prefix string, ticker string) {
	if _, exists := c.subscriptions[prefix]; !exists {
		c.subscriptions[prefix] = make(set)
	}
	if _, exists := c.subscriptions[prefix][ticker]; !exists {
		c.log.Infof("already unsubscribed to this ticker")
	}
	delete(c.subscriptions[prefix], ticker)
}

func (c *Client) Subscribe(topic Topic, tickers ...string) error {
	params, err := getParams(c.market, topic, tickers...)
	if err != nil {
		return err
	}

	for _, t := range tickers {
		c.setSubscription(topic.prefix(), t)
	}

	subscribe, err := json.Marshal(&models.ControlMessage{
		Action: models.Subscribe,
		Params: params,
	})
	if err != nil {
		return err
	}

	c.log.Debugf("subscribing to '%v'", params) // todo: remove before release
	c.wQueue <- subscribe
	return nil
}

func (c *Client) Unsubscribe(topic Topic, tickers ...string) error {
	params, err := getParams(c.market, topic, tickers...)
	if err != nil {
		return err
	}

	for _, t := range tickers {
		c.deleteSubscription(topic.prefix(), t)
	}

	unsubscribe, err := json.Marshal(&models.ControlMessage{
		Action: models.Unsubscribe,
		Params: params,
	})
	if err != nil {
		return err
	}

	c.log.Debugf("unsubscribing from '%v'", params) // todo: remove before release
	c.wQueue <- unsubscribe
	return nil
}

func (c *Client) Close() error {
	if c.conn == nil {
		return nil
	}

	c.cancel()
	// todo: verify that this is thread-safe and potentially refactor to just push a message to the wQueue
	err := c.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(writeWait))
	if err != nil {
		c.log.Errorf("failed to gracefully close: %v", err)
		return err
	}
	c.log.Infof("connection closed successfully")
	return nil
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
		c.conn.Close() // todo: should this force close?
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
		c.conn.Close() // todo: should this force close?
	}()

	for {
		select {
		case <-c.ctx.Done():
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
				return // todo: should this return?
			}
			if err := c.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				c.log.Errorf("failed to send message: %v", err)
				return
			}
		}
	}
}

// todo: add config option to skip message processing
func (c *Client) process() {
	defer func() {
		c.log.Debugf("closing process thread")
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

// todo: this might merit a "data router" type
func (c *Client) route(msgs []json.RawMessage) {
	for _, msg := range msgs {
		var ev models.EventType
		err := json.Unmarshal(msg, &ev)
		if err != nil {
			c.log.Errorf("failed to process message: %v", err)
			return
		}

		switch ev.EventType {
		case "status":
			c.handleStatus(msg)
		default:
			c.log.Debugf("unknown message type '%v'", ev.EventType)
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
		c.Close()
		return
	case "success":
		c.log.Infof("subscription successful") // todo: improve this
	case "error":
		c.log.Errorf("received an error: %v", cm.Message) // todo: improve this
	default:
		c.log.Infof("unknown status message '%v'", cm.Status)
	}
}
