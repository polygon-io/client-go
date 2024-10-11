// Stocks - Short Interest
// https://polygon.io/docs/stocks/get_v1_reference_short-interest__identifierType___identifier
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

	// Initialize client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// Set parameters
	params := models.ListShortInterestParams{
		IdentifierType: "ticker",
		Identifier:     "AMD",
	}.WithDay(models.GTE, 2024, time.October, 1).
		WithDay(models.LTE, 2024, time.October, 10).
		WithLimit(10)

	// Make request
	iter := c.ListShortInterest(context.Background(), params)

	// Do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
