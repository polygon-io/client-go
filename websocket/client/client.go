package client

import (
	"errors"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/polygon-io/client-go/websocket/models"
)

const APIURL = "wss://socket.polygon.io"

// Client defines a dialer that knows how to connect to the Polygon WebSocket server.
type Client struct {
	APIKey string
	Dialer *websocket.Dialer
}

// New creates a client for the Polygon WebSocket API.
func New(apiKey string) Client {
	return Client{
		APIKey: apiKey,
		Dialer: websocket.DefaultDialer, // todo: is this a sufficient default?
	}
}

// Connect dials the WebSocket server, authenticates, and subcribes to a data feed.
func (c *Client) Connect(market models.MarketType, params string) (*Conn, error) {
	conn, err := c.dial(market)
	if err != nil {
		return nil, fmt.Errorf("failed to dial the websocket server: %w", err)
	}

	if err := conn.authenticate(c.APIKey); err != nil {
		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}

	if err := conn.subscribe(params); err != nil {
		return nil, fmt.Errorf("failed to subscribe to feed '%v': %w", params, err)
	}

	return conn, nil
}

func (c *Client) dial(market models.MarketType) (*Conn, error) {
	url := APIURL + "/" + string(market)
	conn, _, err := c.Dialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	var res []models.ControlResponse
	err = conn.ReadJSON(&res)
	if err != nil {
		return nil, err
	} else if len(res) == 0 || res[0].Status != "connected" {
		return nil, errors.New("unable to connect")
	}

	return &Conn{
		conn: conn,
	}, nil
}

// Conn defines a connection to a WebSocket server.
type Conn struct {
	conn *websocket.Conn
}

// Collect reads from the connection and pushes raw data to a channel.
func (c *Conn) Collect(data chan []byte) {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return // an error here likely means the connection has closed
		}
		data <- msg
	}
}

// Close attempts to gracefully close the connection to the server.
func (c *Conn) Close() error {
	if err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		c.conn.Close()
		return fmt.Errorf("failed to gracefully disconnect: %w", err)
	}
	return nil
}

func (c *Conn) authenticate(apiKey string) error {
	err := c.conn.WriteJSON(models.ControlRequest{
		Action: "auth",
		Params: apiKey,
	})
	if err != nil {
		return err
	}

	var res []models.ControlResponse
	err = c.conn.ReadJSON(&res)
	if err != nil {
		return err
	} else if len(res) == 0 || res[0].Status != "auth_success" {
		return errors.New("unauthorized")
	}

	return nil
}

func (c *Conn) subscribe(params string) error {
	err := c.conn.WriteJSON(models.ControlRequest{
		Action: "subscribe",
		Params: params,
	})
	if err != nil {
		return err
	}

	var res []models.ControlResponse
	err = c.conn.ReadJSON(&res)
	if err != nil {
		return err
	} else if len(res) == 0 || res[0].Status != "success" {
		return errors.New("subscribe unsuccessful")
	}

	return nil
}
