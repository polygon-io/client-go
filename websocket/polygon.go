package polygonws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/polygon-io/client-go/websocket/models"
)

// todo: add reconnect logic
// todo: in general, successful calls should be debug and unknown messages should be info

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
		config.Log = &nopLogger{}
	}

	// todo: is this default dialer sufficient? might want to let user pass in a context so they can cancel the dial
	conn, res, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %w", err)
	} else if res.StatusCode != 101 {
		return nil, errors.New("server failed to switch protocols")
	}

	conn.SetReadLimit(maxMessageSize)
	if err := conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		return nil, fmt.Errorf("failed to set read deadline: %w", err)
	}
	conn.SetPongHandler(func(string) error {
		return conn.SetReadDeadline(time.Now().Add(pongWait))
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
	// todo: verify that this is thread-safe and potentially refactor to just push a message to the wQueue
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
		c.log.Errorf("failed to process message: %v", err)
	}
}

func (c *Client) handleStatus(msg json.RawMessage) {
	var cm models.ControlMessage
	if err := json.Unmarshal(msg, &cm); err != nil {
		c.log.Errorf("failed to unmarshal message")
		return
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
		return
	case "success":
		c.log.Infof("subscription successful") // todo: can subscriptions fail?
	default:
		c.log.Debugf("unknown status message '%v'", cm.Status)
	}
}
