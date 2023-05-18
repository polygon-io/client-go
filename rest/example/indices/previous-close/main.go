// Indices - Previous Close
// https://polygon.io/docs/indices/get_v2_aggs_ticker__indicesticker__prev
// https://github.com/polygon-io/client-go/blob/master/rest/aggs.go
package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := models.GetPreviousCloseAggParams{
		Ticker: "I:SPX",
	}.WithAdjusted(true)

	// make request
	res, err := c.GetPreviousCloseAgg(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
