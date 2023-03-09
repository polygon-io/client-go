// options - Moving Average Convergence/Divergence (MACD)
// https://polygon.io/docs/options/get_v1_indicators_macd__optionsticker
// https://github.com/polygon-io/client-go/blob/master/rest/indicators.go
package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {

	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// set params
	params := &models.GetMACDParams{
		Ticker: "O:SPY241220P00720000",
	}

	// make request
	res, err := c.GetMACD(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res) 

}