package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	// Initialize client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// Set parameters
	params := models.ListShortVolumeParams{}.
		WithTicker(models.EQ, "A").
		WithDate(models.EQ, "2025-03-25").
		WithOrder(models.Asc).
		WithLimit(50000)

	// Make request
	iter := c.ListShortVolume(context.Background(), params)

	// Process results
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
