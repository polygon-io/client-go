// Crypto - Ticker Full Book (L2)
// https://polygon.io/docs/crypto/get_v2_snapshot_locale_global_markets_crypto_tickers__ticker__book
// https://github.com/polygon-io/client-go/blob/master/rest/snapshot.go
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
	params := &models.GetCryptoFullBookSnapshotParams{
		Ticker: "X:BTCUSD",
	}

	// make request
	res, err := c.GetCryptoFullBookSnapshot(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
