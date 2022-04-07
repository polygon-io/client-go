package main

import (
	"fmt"

	websocket "github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const APIKEY = "YOUR_API_KEY"
const CHANNELS = "T.SPY,Q.SPY"

func main() {
	c, _, err := websocket.DefaultDialer.Dial("wss://socket.polygon.io/stocks", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_ = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"action\":\"auth\",\"params\":\"%s\"}", APIKEY)))
	_ = c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("{\"action\":\"subscribe\",\"params\":\"%s\"}", CHANNELS)))

	// Buffered channel to account for bursts or spikes in data:
	chanMessages := make(chan interface{}, 10000)

	// Read messages off the buffered queue:
	go func() {
		for msgBytes := range chanMessages {
			logrus.Info("Message Bytes: ", msgBytes)
		}
	}()

	// As little logic as possible in the reader loop:
	for {
		var msg interface{}
		err := c.ReadJSON(&msg)
		// Ideally use c.ReadMessage instead of ReadJSON so you can parse the JSON data in a
		// separate go routine. Any processing done in this loop increases the chances of disconnects
		// due to not consuming the data fast enough.
		if err != nil {
			panic(err)
		}
		chanMessages <- msg
	}
}
