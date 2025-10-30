// Stocks - Last Quote (NBBO)
// https://massive.com/docs/stocks/get_v2_last_nbbo__stocksticker
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
	params := &models.GetLastQuoteParams{
		Ticker: "AAPL",
	}

	// make request
	res, err := c.GetLastQuote(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
