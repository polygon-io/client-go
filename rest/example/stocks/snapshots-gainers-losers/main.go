// Stocks - Snapshot Gainers/Losers
// https://polygon.io/docs/stocks/get_v2_snapshot_locale_us_markets_stocks__direction
// https://github.com/polygon-io/client-go/blob/master/rest/snapshot.go
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
	params := &models.GetGainersLosersSnapshotParams{
		Locale:     "us",
		MarketType: "stocks",
		Direction:  "gainers", // or "losers"
	}

	// make request
	res, err := c.GetGainersLosersSnapshot(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
