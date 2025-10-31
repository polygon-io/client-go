// Forex - Grouped Daily (Bars)
// https://massive.com/docs/forex/get_v2_aggs_grouped_locale_global_market_fx__date
// https://github.com/massive-com/client-go/v2/blob/master/rest/aggs.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	massive "github.com/massive-com/client-go/v2/rest"
	"github.com/massive-com/client-go/v2/rest/models"
)

func main() {
	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// set params
	params := models.GetGroupedDailyAggsParams{
		Locale:     models.Global,
		MarketType: "fx",
		Date:       models.Date(time.Date(2023, 3, 8, 0, 0, 0, 0, time.Local)),
	}.WithAdjusted(true)

	// make request
	res, err := c.GetGroupedDailyAggs(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
