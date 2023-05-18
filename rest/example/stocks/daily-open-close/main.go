// Stocks - Daily Open/Close
// https://polygon.io/docs/stocks/get_v1_open-close__stocksticker___date
// https://github.com/polygon-io/client-go/blob/master/rest/aggs.go
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

	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := models.GetDailyOpenCloseAggParams{
		Ticker: "AAPL",
		Date:   models.Date(time.Date(2023, 3, 8, 0, 0, 0, 0, time.Local)),
	}.WithAdjusted(true)

	// make request
	res, err := c.GetDailyOpenCloseAgg(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
