// Indices - Snapshot
// https://polygon.io/docs/indices/get_v3_snapshot_indices
// https://github.com/polygon-io/client-go/blob/master/rest/snapshot.go
package main

import (
	"context"
	"log"
	"os"
	"strings"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

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
