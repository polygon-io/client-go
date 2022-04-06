package polygonws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/polygon-io/client-go/websocket/models"
	"go.uber.org/zap"
)

// todo: add reconnect logic

type Client struct {
	apiKey string

	ctx    context.Context
	cancel context.CancelFunc

	conn   *websocket.Conn
	rQueue chan []byte
	wQueue chan []byte

	log Logger
}

// todo: might want to separate Connect logic out from New
func New(config Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, errors.New("API key is required")
	}

	url := fmt.Sprintf("wss://%v.polygon.io/%v", string(config.Feed), string(config.Market))

	if config.Log == nil {
		config.Log = zap.NewNop().Sugar()
	}

	// todo: is this default dialer sufficient? might want to let user pass in a context so they can cancel the dial
	conn, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 101 {
		return nil, errors.New("server failed to switch protocols")
	}
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	ctx, cancel := context.WithCancel(context.Background())

	c := &Client{
		apiKey: config.APIKey,
		ctx:    ctx,
		cancel: cancel,
		conn:   conn,
		rQueue: make(chan []byte, 10000),
		wQueue: make(chan []byte, 100),
		log:    config.Log,
	}

	go c.read()
	go c.write()
	go c.process()

	return c, nil
}

// todo: Subscribe, Unsubscribe, etc

func (c *Client) Close() error {
	c.cancel()
	err := c.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(writeWait))
	if err != nil {
		c.log.Errorf("failed to gracefully close: %v", err)
		return err
	}
	c.log.Infof("connection closed successfully")
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
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.log.Errorf("failed to send ping message: %v", err)
				return
			}
		case msg := <-c.wQueue:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
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
	}()

	for {
		select {
		case <-c.ctx.Done():
			return
		case data := <-c.rQueue:
			c.route(data) // todo: this might merit a "data router" type
		}
	}
}

func (c *Client) route(data []byte) {
	var msgs []json.RawMessage
	if err := json.Unmarshal(data, &msgs); err != nil {
		c.log.Errorf("failed to process raw messages: %v", err)
		return
	}

	for _, msg := range msgs {
		if err := c.handle(msg); err != nil {
			c.log.Errorf("failed to process message: %v", err)
		}
	}
}

func (c *Client) handle(msg json.RawMessage) error {
	var ev models.EventType
	err := json.Unmarshal(msg, &ev)
	if err != nil {
		return err
	}

	switch ev.EventType {
	case "status":
		return c.handleStatus(msg)
	default:
		c.log.Debugf("unknown message type '%v'", ev.EventType)
	}

	return nil
}

func (c *Client) handleStatus(msg json.RawMessage) error {
	var cm models.ControlMessage
	if err := json.Unmarshal(msg, &cm); err != nil {
		return err
	}

	switch cm.Status {
	case "connected":
		c.log.Infof("connection successful")
		b, err := json.Marshal(models.ControlMessage{
			Action: "auth",
			Params: c.apiKey,
		})
		if err != nil {
			c.log.Errorf("authentication failed, closing connection")
			c.Close() // fatal errors should close the connection
		}
		c.wQueue <- b
	case "auth_success":
		c.log.Infof("authentication successful")
	case "auth_failed":
		c.log.Errorf("authentication failed, closing connection")
		c.Close()
		return nil
	case "success":
		c.log.Infof("subscription successful") // todo: can subscriptions fail?
	default:
		c.log.Debugf("unknown status message '%v'", cm.Status)
	}

	return nil
}
