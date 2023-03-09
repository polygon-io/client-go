// options - Aggregates (Bars)
// https://polygon.io/docs/options/get_v2_aggs_ticker__optionsticker__range__multiplier___timespan___from___to
// https://github.com/polygon-io/client-go/blob/master/rest/aggs.go
package main

import (
	"context"
	"time"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {

	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := models.GetAggsParams{
		Ticker:     "O:SPY251219C00650000",
		Multiplier: 1,
		Timespan:   "day",
		From:       models.Millis(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
		To:         models.Millis(time.Date(2023, 3, 9, 0, 0, 0, 0, time.UTC)),
	}.WithOrder(models.Desc).WithLimit(2).WithAdjusted(true)

	// make request
	res, err := c.GetAggs(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res) 

}