// Stocks - Snapshot Gainers/Losers
// https://massive.com/docs/stocks/get_v2_snapshot_locale_us_markets_stocks__direction
// https://github.com/massive-com/client-go/blob/master/rest/snapshot.go
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
