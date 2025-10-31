// Forex - Relative Strength Index (RSI)
// https://massive.com/docs/forex/get_v1_indicators_rsi__fxticker
// https://github.com/massive-com/client-go/v2/blob/master/rest/indicators.go
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
	params := &models.GetRSIParams{
		Ticker: "C:EURUSD",
	}

	// make request
	res, err := c.GetRSI(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
