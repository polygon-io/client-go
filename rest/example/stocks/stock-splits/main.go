// Stocks - Stock Splits v3
// https://massive.com/docs/stocks/get_v3_reference_splits
// https://github.com/massive-com/client-go/blob/master/rest/reference.go
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
	params := models.ListSplitsParams{}.
		WithTicker(models.EQ, "AAPL").
		WithExecutionDate(models.EQ, models.Date(time.Date(2020, 8, 31, 0, 0, 0, 0, time.UTC))).
		WithReverseSplit(false).
		WithSort(models.TickerSymbol).
		WithOrder(models.Asc).
		WithLimit(1000)

	// make request
	iter := c.ListSplits(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
