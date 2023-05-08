package main

import (
	"context"
	"log"
	"os"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	// Init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// Set parameters
	params := models.ListAssetSnapshotsParams{}.
		WithTickerAnyOf("AAPL,META,F").
		WithTimestamp(models.EQ, models.Nanos(time.Date(2023, 05, 8, 0, 0, 0, 0, time.UTC)))

	// Make request
	iter := c.ListAssetSnapshots(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Println(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
