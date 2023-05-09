package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	// Init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// Set parameters
	params := models.ListAssetSnapshotsParams{}.
		WithTickerAnyOf("AAPL")

	// Make request
	iter := c.ListAssetSnapshots(context.Background(), params)

	// do something with the result
	for iter.Next() {
		data, _ := json.MarshalIndent(iter.Item(), "", "    ")
		fmt.Println(string(data))
		//log.Println(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
