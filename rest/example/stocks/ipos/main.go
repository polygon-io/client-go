// Stocks - IPOs
// https://polygon.io/docs/stocks/get_v1_reference_ipos
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
	params := models.ListIPOsParams{}.WithTicker("RDDT")

	// make request
	iter := c.ListIPOs(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
