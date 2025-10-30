// Indices - Previous Close
// https://massive.com/docs/indices/get_v2_aggs_ticker__indicesticker__prev
// https://github.com/massive-com/client-go/blob/master/rest/aggs.go
package main

import (
	"context"
	"log"
	"os"

	massive "github.com/massive-com/client-go/rest"
	"github.com/massive-com/client-go/rest/models"
)

func main() {
	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

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
