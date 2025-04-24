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
	params := models.ListIPOsParams{}.
		WithTicker("RAPP").
		WithListingDate(models.EQ, "2024-06-07").
		WithOrder(models.Desc).
		WithLimit(1000)

	// Make request
	iter := c.VX.ListIPOs(context.Background(), params)

	// Process results
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
