package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	// Init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// Set parameters
	params := models.ListAssetSnapshotsParams{}.WithTickerAnyOf("I:A1BSC,I:A1CYC,I:A1DOW")

	// Make request
	iter := c.ListAssetSnapshots(context.Background(), params)

	// do something with the result
	for iter.Next() {
		val, err := json.MarshalIndent(iter.Item(), "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(val))
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
