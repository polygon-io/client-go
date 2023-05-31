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
		Market: polygonws.Stocks, // Change the Market to match when running other examples
		Log:    log,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	err = c.Subscribe(polygonws.StocksLaunchpadMinAggs, "*")
	// _ = c.Subscribe(polygonws.StocksLaunchpadValue, "*")
	// _ = c.Subscribe(polygonws.OptionsLaunchpadMinAggs, "O:A230616C00070000")
	// _ = c.Subscribe(polygonws.OptionsLaunchpadValue, "O:A230616C00070000")
	// _ = c.Subscribe(polygonws.ForexLaunchpadMinAggs, "*")
	// _ = c.Subscribe(polygonws.ForexLaunchpadValue, "*")
	// _ = c.Subscribe(polygonws.CryptoLaunchpadMinAggs, "*")
	// _ = c.Subscribe(polygonws.CryptoLaunchpadValue, "*")

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
			case models.CurrencyAgg:
				log.WithFields(logrus.Fields{"currency aggregate": out}).Info()
			case models.EquityAgg:
				log.WithFields(logrus.Fields{"equity aggregate": out}).Info()
			case models.LaunchpadValue:
				log.WithFields(logrus.Fields{"value": out}).Info()
			}
		}
	}
}
