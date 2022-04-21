package polygonws_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	polygonws "github.com/polygon-io/client-go/websocket"
	"github.com/sirupsen/logrus"
)

func TestMain(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return // skip in CI for now
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	c, err := polygonws.New(polygonws.Config{
		APIKey:    apiKey,
		Feed:      polygonws.RealTime,
		Market:    polygonws.Stocks,
		ParseData: true, // comment for raw data handling
		Log:       log,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// test subscribing before connecting
	if err := c.Subscribe(polygonws.StocksSecAggs, "AAPL", "MSFT"); err != nil {
		log.Error(err)
	}

	c.Close() // this shouldn't panic
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	// calling connect again shouldn't panic or data race
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	go printOutput(c) // comment for raw data handling
	// go printRawOutput(c) // uncomment for raw data handling

	time.Sleep(5 * time.Second)
	if err := c.Subscribe(polygonws.StocksTrades, "*"); err != nil {
		log.Error(err)
	}

	time.Sleep(250 * time.Millisecond)
	if err := c.Unsubscribe(polygonws.StocksTrades, "*"); err != nil {
		log.Error(err)
	}

	time.Sleep(5 * time.Second)
	if err := c.Unsubscribe(polygonws.StocksSecAggs, "MSFT"); err != nil {
		log.Error(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func printOutput(c *polygonws.Client) {
	for {
		out := c.Output()
		if out == nil {
			break
		}
		fmt.Println(out)
	}
}

//nolint:deadcode
func printRawOutput(c *polygonws.Client) {
	for {
		out := c.Output()
		if out == nil {
			break
		}
		if b, ok := out.(json.RawMessage); ok {
			fmt.Println(string(b))
		}
	}
}
