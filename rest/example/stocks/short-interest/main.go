package main

import (
	"context"
	"log"
	"os"

	massive "github.com/massive-com/client-go/v2/rest"
	"github.com/massive-com/client-go/v2/rest/models"
)

func main() {
	// Initialize client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// Set parameters
	params := models.ListShortInterestParams{}.
		WithTicker(models.EQ, "A").
		WithSettlementDate(models.EQ, "2025-03-14").
		WithOrder(models.Asc).
		WithLimit(50000)

	// Make request
	iter := c.ListShortInterest(context.Background(), params)

	// Process results
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
