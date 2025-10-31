// Forex - Quotes (BBO)
// https://massive.com/docs/forex/get_v3_quotes__fxticker
// https://github.com/massive-com/client-go/v2/blob/master/rest/quotes.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	massive "github.com/massive-com/client-go/v2/rest"
	"github.com/massive-com/client-go/v2/rest/models"
)

func main() {
	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// set params
	params := models.ListQuotesParams{
		Ticker: "C:EUR-USD",
	}.WithTimestamp(models.EQ, models.Nanos(time.Date(2023, 4, 13, 0, 0, 0, 0, time.UTC))).
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
