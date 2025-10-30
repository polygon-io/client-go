// Stocks - Stock Financials vX
// https://massive.com/docs/stocks/get_vx_reference_financials
// https://github.com/massive-com/client-go/blob/master/rest/vx.go
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
