// Indices - Moving Average Convergence/Divergence (MACD)
// https://massive.com/docs/indices/get_v1_indicators_macd__indicesticker
// https://github.com/massive-com/client-go/blob/master/rest/indicators.go
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
	params := models.GetMACDParams{
		Ticker: "I:SPX",
	}.WithShortWindow(12).
		WithLongWindow(26).
		WithSignalWindow(9).
		WithOrder(models.Desc)

	// make request
	res, err := c.GetMACD(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
