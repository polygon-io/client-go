// Forex - Snapshot Gainers/Losers
// https://massive.com/docs/forex/get_v2_snapshot_locale_global_markets_forex__direction
// https://github.com/massive-com/client-go/v2/blob/master/rest/snapshot.go
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
	params := &models.GetGainersLosersSnapshotParams{
		Locale:     models.Global,
		MarketType: models.Forex,
		Direction:  models.Gainers, // or models.Losers
	}

	// make request
	res, err := c.GetGainersLosersSnapshot(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
