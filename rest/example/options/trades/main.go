// Options - Trades
// https://massive.com/docs/options/get_v3_trades__optionsticker
// https://github.com/massive-com/client-go/blob/master/rest/trades.go
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
