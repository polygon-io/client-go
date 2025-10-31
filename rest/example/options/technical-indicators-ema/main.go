// Options - Exponential Moving Average (EMA)
// https://massive.com/docs/options/get_v1_indicators_ema__optionsticker
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
	params := &models.GetEMAParams{
		Ticker: "O:SPY241220P00720000",
	}

	// make request
	res, err := c.GetEMA(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
