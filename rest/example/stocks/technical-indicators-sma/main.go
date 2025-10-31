// Stocks - Simple Moving Average (SMA)
// https://massive.com/docs/stocks/get_v1_indicators_sma__stockticker
// https://github.com/massive-com/client-go/v2/blob/master/rest/indicators.go
package main

import (
	"context"
	"log"
	"os"

	massive "github.com/massive-com/client-go/v2/rest"
	"github.com/massive-com/client-go/v2/rest/models"
)

func main() {

	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// set params
	params := models.GetSMAParams{
		Ticker: "AAPL",
	}.WithLimit(1000)

	// make request
	res, err := c.GetSMA(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
