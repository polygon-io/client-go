// Crypto - Last Trade for a Crypto Pair
// https://massive.com/docs/crypto/get_v1_last_crypto__from___to
// https://github.com/massive-com/client-go/blob/master/rest/trades.go
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
