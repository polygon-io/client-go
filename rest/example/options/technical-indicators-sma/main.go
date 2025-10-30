// Options - Simple Moving Average (SMA)
// https://massive.com/docs/options/get_v1_indicators_sma__optionsticker
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
	params := &models.GetSMAParams{
		Ticker: "O:SPY241220P00720000",
	}

	// make request
	res, err := c.GetSMA(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
