// Crypto - Tickers
// https://massive.com/docs/crypto/get_v3_reference_tickers
// https://github.com/massive-com/client-go/v2/blob/master/rest/reference.go
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
	params := models.ListTickersParams{}.
		WithMarket(models.AssetCrypto).
		WithActive(true).
		WithSort(models.TickerSymbol).
		WithOrder(models.Asc).
		WithLimit(1000)

	// make request
	iter := c.ListTickers(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
