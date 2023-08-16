// Stocks - Stock Financials vX
// https://polygon.io/docs/stocks/get_vx_reference_financials
// https://github.com/polygon-io/client-go/blob/master/rest/vx.go
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
	params := models.ListStockFinancialsParams{}.
		WithTicker("AAPL")

	// make request
	iter := c.VX.ListStockFinancials(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
