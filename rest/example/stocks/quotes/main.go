// Stocks - Quotes (NBBO)
// https://massive.com/docs/stocks/get_v3_quotes__stockticker
// https://github.com/massive-com/client-go/blob/master/rest/quotes.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	massive "github.com/massive-com/client-go/rest"
	"github.com/massive-com/client-go/rest/models"
)

func main() {

	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// set params
	params := models.ListQuotesParams{
		Ticker: "AAPL",
	}.WithTimestamp(models.EQ, models.Nanos(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).
		WithSort(models.Timestamp).
		WithOrder(models.Asc).
		WithLimit(50000)

	// make request
	iter := c.ListQuotes(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
