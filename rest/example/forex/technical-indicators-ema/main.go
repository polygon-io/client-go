// Forex - Exponential Moving Average (EMA)
// https://polygon.io/docs/forex/get_v1_indicators_ema__fxticker
// https://github.com/polygon-io/client-go/blob/master/rest/indicators.go
package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := &models.GetEMAParams{
		Ticker: "C:EURUSD",
	}

	// make request
	res, err := c.GetEMA(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
