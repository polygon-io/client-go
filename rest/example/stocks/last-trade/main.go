// Stocks - Last Trade
// https://massive.com/docs/stocks/get_v2_last_trade__stocksticker
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
	params := &models.GetLastTradeParams{
		Ticker: "AAPL",
	}

	// make request
	res, err := c.GetLastTrade(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
