package main

import (
	"context"
	"log"
	"os"

	massive "github.com/massive-com/client-go/rest"
	"github.com/massive-com/client-go/rest/models"
)

func main() {
	// Init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// Set parameters
	params := models.ListUniversalSnapshotsParams{}.
		WithTickerAnyOf("O:AAPL230512C00050000,O:META230512C00020000,O:F230512C00005000")

	// Make request
	iter := c.ListUniversalSnapshots(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Println(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
