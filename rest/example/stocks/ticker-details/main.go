// Stocks - Ticker Details v3
// https://massive.com/docs/stocks/get_v3_reference_tickers__ticker
// https://github.com/massive-com/client-go/blob/master/rest/reference.go
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
	params := models.GetTickerDetailsParams{
		Ticker: "AAPL",
	}.WithDate(models.Date(time.Date(2023, 3, 9, 0, 0, 0, 0, time.UTC)))

	// make request
	res, err := c.GetTickerDetails(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
