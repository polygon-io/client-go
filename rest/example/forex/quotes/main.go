// Forex - Quotes (BBO)
// https://polygon.io/docs/forex/get_v3_quotes__fxticker
// https://github.com/polygon-io/client-go/blob/master/rest/quotes.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

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
