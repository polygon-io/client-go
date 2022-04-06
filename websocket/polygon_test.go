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
		Market: polygonws.Stocks,
		Log:    log,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
