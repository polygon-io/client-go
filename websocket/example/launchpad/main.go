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
		Market: polygonws.Stocks, // Change Market to match examples below (Stocks, Options, Forex, or Crypto)
		Log:    log,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// Stocks - make sure you change "Market: polygonws.Stocks" above in the polygonws.Config
	err = c.Subscribe(polygonws.StocksLaunchpadMinAggs, "*")
	//err = c.Subscribe(polygonws.StocksLaunchpadValue, "*")

	// Options - make sure you change "Market: polygonws.Options" above in the polygonws.Config
	//err = c.Subscribe(polygonws.OptionsLaunchpadMinAggs, "O:A230616C00070000")
	//err = c.Subscribe(polygonws.OptionsLaunchpadValue, "O:A230616C00070000")

	// Forex - make sure you change "Market: polygonws.Forex" above in the polygonws.Config
	//err = c.Subscribe(polygonws.ForexLaunchpadMinAggs, "*")
	//err = c.Subscribe(polygonws.ForexLaunchpadValue, "*")

	// Crypto - make sure you change "Market: polygonws.Crypto" above in the polygonws.Config
	//err = c.Subscribe(polygonws.CryptoLaunchpadMinAggs, "*")
	//err = c.Subscribe(polygonws.CryptoLaunchpadValue, "*")

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
