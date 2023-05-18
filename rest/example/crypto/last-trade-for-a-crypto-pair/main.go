// Crypto - Last Trade for a Crypto Pair
// https://polygon.io/docs/crypto/get_v1_last_crypto__from___to
// https://github.com/polygon-io/client-go/blob/master/rest/trades.go
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
	params := &models.GetLastCryptoTradeParams{
		From: "BTC",
		To:   "USD",
	}

	// make request
	res, err := c.GetLastCryptoTrade(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
