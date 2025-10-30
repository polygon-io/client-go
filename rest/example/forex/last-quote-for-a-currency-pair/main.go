// Forex - Last Quote for a Currency Pair
// https://massive.com/docs/forex/get_v1_last_quote_currencies__from___to
// https://github.com/massive-com/client-go/blob/master/rest/quotes.go
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
	params := &models.GetLastForexQuoteParams{
		From: "AUD",
		To:   "USD",
	}

	// make request
	res, err := c.GetLastForexQuote(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
