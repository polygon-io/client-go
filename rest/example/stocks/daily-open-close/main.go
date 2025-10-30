// Stocks - Daily Open/Close
// https://massive.com/docs/stocks/get_v1_open-close__stocksticker___date
// https://github.com/massive-com/client-go/blob/master/rest/aggs.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	massive "github.com/massive-com/client-go/rest"
	"github.com/massive-com/client-go/rest/models"
)

func main() {

	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

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
