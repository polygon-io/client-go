// Forex - Snapshot All Tickers
// https://massive.com/docs/forex/get_v2_snapshot_locale_global_markets_forex_tickers
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
