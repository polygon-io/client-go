package polygonws_test

import (
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	c, err := polygonws.New(polygonws.Config{
		APIKey: apiKey,
		Feed:   polygonws.RealTime,
		Market: polygonws.Crypto,
		Log:    log,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	c.Close() // this shouldn't panic
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	// calling connect again shouldn't panic or data race
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}

	if err := c.Subscribe(polygonws.CryptoTrades, "BTC-USD", "ETH-USD"); err != nil {
		log.Error(err)
	}
	time.Sleep(1 * time.Second)
	if err := c.Unsubscribe(polygonws.CryptoTrades, "BTC-USD"); err != nil {
		log.Error(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
