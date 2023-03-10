// Stocks - Dividends v3
// https://polygon.io/docs/stocks/get_v3_reference_dividends
// https://github.com/polygon-io/client-go/blob/master/rest/reference.go
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
