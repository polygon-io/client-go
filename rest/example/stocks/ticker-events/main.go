// Stocks - Ticker Events
// https://massive.com/docs/stocks/get_vx_reference_tickers__id__events
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
	params := &models.GetTickerEventsParams{
		ID: "TSLA",
	}

	// make request
	res, err := c.VX.GetTickerEvents(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
