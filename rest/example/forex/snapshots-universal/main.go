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
		WithTickerAnyOf("C:USDCAD,C:USDEUR,C:USDAUD")

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
