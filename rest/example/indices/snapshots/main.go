// Indices - Snapshot
// https://massive.com/docs/indices/get_v3_snapshot_indices
// https://github.com/massive-com/client-go/blob/master/rest/snapshot.go
package main

import (
	"context"
	"log"
	"os"
	"strings"

	massive "github.com/massive-com/client-go/rest"
	"github.com/massive-com/client-go/rest/models"
)

func main() {
	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// define tickers
	tickers := []string{"I:SPX", "I:DJI"}
	tickerAnyOf := strings.Join(tickers, ",")

	// set params
	params := &models.GetIndicesSnapshotParams{
		TickerAnyOf: &tickerAnyOf,
	}

	// make request
	res, err := c.GetIndicesSnapshot(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
