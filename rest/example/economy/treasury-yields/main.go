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
	params := models.ListTreasuryYieldsParams{}.
		WithDate(models.GTE, "2024-12-15").
		WithDate(models.LTE, "2024-12-31").
		WithOrder(models.Asc).
		WithLimit(50000)

	// Make request
	iter := c.VX.ListTreasuryYields(context.Background(), params)

	// Process results
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
