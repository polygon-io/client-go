// options - Trades
// https://polygon.io/docs/options/get_v3_trades__optionsticker
// https://github.com/polygon-io/client-go/blob/master/rest/trades.go
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
	params := models.ListTradesParams{
		Ticker: "O:TSLA210903C00700000",
	}.WithLimit(50000).WithOrder(models.Asc)

	// make request
	iter := c.ListTrades(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item()) 
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
