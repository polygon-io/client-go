package main

import (
	"os"
	"os/signal"

	polygonws "github.com/polygon-io/client-go/websocket"
	"github.com/polygon-io/client-go/websocket/models"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})
	c, err := polygonws.New(polygonws.Config{
		APIKey: os.Getenv("POLYGON_API_KEY"),
		Feed:   polygonws.LaunchpadFeed,
		Market: polygonws.Stocks,
		Log:    log,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	err = c.Subscribe(polygonws.StocksLaunchpadMinAggs, "MSFT, META")
	// err = c.Subscribe(polygonws.StocksLaunchpadValue, "*")
	if err != nil {
		log.Error(err)
		return
	}

	if err := c.Connect(); err != nil {
		log.Error(err)
		return
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	for {
		select {
		case <-sigint:
			return
		case <-c.Error():
			return
		case out, more := <-c.Output():
			if !more {
				return
			}
			switch out.(type) {
			case models.EquityAgg:
				log.WithFields(logrus.Fields{"aggregate": out}).Info()
			case models.EquityTrade:
				log.WithFields(logrus.Fields{"trade": out}).Info()
			case models.EquityQuote:
				log.WithFields(logrus.Fields{"quote": out}).Info()
			}
		}
	}
}
