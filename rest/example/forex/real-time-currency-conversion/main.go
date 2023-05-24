// Forex - Real-time Currency Conversion
// https://polygon.io/docs/forex/get_v1_conversion__from___to
// https://github.com/polygon-io/client-go/blob/master/rest/quotes.go
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
	params := &models.GetRealTimeCurrencyConversionParams{
		From: "AUD",
		To:   "USD",
	}

	// make request
	res, err := c.GetRealTimeCurrencyConversion(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
