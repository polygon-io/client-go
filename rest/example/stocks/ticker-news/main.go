// Stocks - Ticker News
// https://polygon.io/docs/stocks/get_v2_reference_news
// https://github.com/polygon-io/client-go/blob/master/rest/reference.go
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
	params := models.ListTickerNewsParams{}.
		WithTicker(models.LTE, "AAPL").
		WithPublishedUTC(models.LT, models.Millis(time.Date(2023, 3, 9, 0, 0, 0, 0, time.UTC))).
		WithSort(models.PublishedUTC).
		WithOrder(models.Asc).
		WithLimit(1000)

	// make request
	iter := c.ListTickerNews(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item()) 
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
