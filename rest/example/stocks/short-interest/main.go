// Stocks - Short Interest
// https://polygon.io/docs/stocks/get_v1_reference_short-interest__identifierType___identifier
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
	params := models.GetShortInterestParams{
		IdentifierType: "ticker",
		Identifier:     "AAPL",
	}

	// Make request
	res, err := c.GetShortInterest(context.Background(), &params)
	if err != nil {
		log.Fatal(err)
	}

	// Handle the result
	log.Print(res)
}
