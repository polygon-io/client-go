// Stocks - Trades
// https://massive.com/docs/stocks/get_v3_trades__stockticker
// https://github.com/massive-com/client-go/v2/blob/master/rest/trades.go
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
	params := models.ListTradesParams{
		Ticker: "IBIO",
	}.WithDay(2023, 2, 1).WithLimit(50000).WithOrder(models.Asc)

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
