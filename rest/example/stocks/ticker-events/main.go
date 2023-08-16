// Stocks - Ticker Events
// https://polygon.io/docs/stocks/get_vx_reference_tickers__id__events
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
