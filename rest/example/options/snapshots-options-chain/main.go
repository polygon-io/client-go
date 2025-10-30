// Options - Options Chain
// https://massive.com/docs/options/get_v3_snapshot_options__underlyingasset
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
	params := models.ListOptionsChainParams{
		UnderlyingAsset: "SPY",
		StrikePriceGTE:  new(float64),
		StrikePriceLTE:  new(float64),
		Limit:           new(int),
	}.WithStrikePrice("gte", 500.00).WithStrikePrice("lte", 600.00).WithLimit(250)

	// make the request
	iter := c.ListOptionsChainSnapshot(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
