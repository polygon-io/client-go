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

	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	c, err := polygonws.New(polygonws.Config{
		APIKey: apiKey,
		Feed:   polygonws.RealTime,
		Market: polygonws.Stocks,
		// RawData: true, // uncomment for raw data handling
		Log: log,
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

	go printOutput(ctx, c) // comment for raw data handling
	// go printRawOutput(ctx, c) // uncomment for raw data handling

	time.Sleep(10 * time.Second)
	if err := c.Subscribe(polygonws.StocksTrades, "*"); err != nil {
		log.Error(err)
	}

	time.Sleep(250 * time.Millisecond)
	if err := c.Unsubscribe(polygonws.StocksTrades); err != nil {
		log.Error(err)
	}

	time.Sleep(5 * time.Second)
	if err := c.Unsubscribe(polygonws.StocksSecAggs, "MSFT"); err != nil {
		log.Error(err)
	}

	time.Sleep(10 * time.Second)
	if err := c.Subscribe(polygonws.StocksSecAggs, "SNAP", "IBM", "LPL"); err != nil {
		log.Error(err)
	}
	time.Sleep(5 * time.Second)
	if err := c.Unsubscribe(polygonws.StocksSecAggs, "SNAP", "*"); err != nil {
		log.Error(err)
	}

	time.Sleep(15 * time.Second)
	if err := c.Subscribe(polygonws.StocksSecAggs, "AAPL", "MSFT"); err != nil {
		log.Error(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}

func printOutput(ctx context.Context, client *polygonws.Client) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			out := client.Output()
			if out == nil {
				continue
			}
			fmt.Println(out)
		}
	}
}

//nolint:deadcode
func printRawOutput(ctx context.Context, client *polygonws.Client) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			out := client.Output()
			if out == nil {
				continue
			}
			if b, ok := out.(json.RawMessage); ok {
				fmt.Println(string(b))
			}
		}
	}
}
