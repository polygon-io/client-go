// Indices - Simple Moving Average (SMA)
// https://massive.com/docs/indices/get_v1_indicators_sma__indicesticker
// https://github.com/massive-com/client-go/blob/master/rest/indicators.go
package main

import (
	"context"
	"log"
	"os"

	massive "github.com/massive-com/client-go/rest"
	"github.com/massive-com/client-go/rest/models"
)

func main() {
	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// set params
	params := &models.GetSMAParams{
		Ticker: "I:SPX",
	}

	// make request
	res, err := c.GetSMA(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
