// Options - Previous Close
// https://massive.com/docs/options/get_v2_aggs_ticker__optionsticker__prev
// https://github.com/massive-com/client-go/v2/blob/master/rest/aggs.go
package main

import (
	"context"
	"log"
	"os"

	massive "github.com/massive-com/client-go/v2/rest"
	"github.com/massive-com/client-go/v2/rest/models"
)

func main() {

	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// set params
	params := models.GetPreviousCloseAggParams{
		Ticker: "O:SPY251219C00650000",
	}.WithAdjusted(true)

	// make request
	res, err := c.GetPreviousCloseAgg(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
