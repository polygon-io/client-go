// Options - Options Chain
// https://polygon.io/docs/options/get_v3_snapshot_options__underlyingasset
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
	params := models.ListOptionsChainParams{
		UnderlyingAsset: "SPY",
		StrikePriceGTE:  new(float64),
		StrikePriceLTE:  new(float64),
		Limit:           new(int),
	}.WithStrikePrice("gte", 500.00).WithStrikePrice("lte", 600.00).WithLimit(250)

	// make request
	iter := c.ListOptionsChainSnapshot(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
