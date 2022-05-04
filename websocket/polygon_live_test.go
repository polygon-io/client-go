package polygonws_test

import (
	"context"
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

	rawData := true
	var maxRetries uint64 = 5
	c, err := polygonws.New(polygonws.Config{
		APIKey:     apiKey,
		Feed:       polygonws.RealTime,
		Market:     polygonws.Stocks,
		MaxRetries: &maxRetries,
		RawData:    rawData,
		Log:        log,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go subscribe(c, log)
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.Error():
			return
		case out, more := <-c.Output():
			if !more {
				return
			}
			if rawData {
				out = fmt.Sprintf("%s", out)
			}
			fmt.Println(out)
		}
	}
}

func subscribe(c *polygonws.Client, log *logrus.Logger) {
	if err := c.Subscribe(polygonws.StocksSecAggs, "AAPL", "MSFT"); err != nil {
		log.Error(err)
	}

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
}
