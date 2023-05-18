// Forex - Snapshot All Tickers
// https://polygon.io/docs/forex/get_v2_snapshot_locale_global_markets_forex_tickers
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
	params := &models.GetAllTickersSnapshotParams{
		Locale:     models.Global,
		MarketType: models.Forex,
	}

	// make request
	res, err := c.GetAllTickersSnapshot(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
