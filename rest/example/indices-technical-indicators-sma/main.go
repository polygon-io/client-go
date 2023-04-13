// Indices - Simple Moving Average (SMA)
// https://polygon.io/docs/indices/get_v1_indicators_sma__indicesticker
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
