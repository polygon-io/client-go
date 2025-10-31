// Stocks - Dividends v3
// https://massive.com/docs/stocks/get_v3_reference_dividends
// https://github.com/massive-com/client-go/v2/blob/master/rest/reference.go
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
	params := models.ListDividendsParams{}.
		WithTicker(models.EQ, "MSFT").
		WithDividendType(models.DividendCD).
		WithLimit(1000).
		WithOrder(models.Desc)

	// make request
	iter := c.ListDividends(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
