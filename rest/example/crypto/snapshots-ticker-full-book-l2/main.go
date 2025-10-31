// Crypto - Ticker Full Book (L2)
// https://massive.com/docs/crypto/get_v2_snapshot_locale_global_markets_crypto_tickers__ticker__book
// https://github.com/massive-com/client-go/v2/blob/master/rest/snapshot.go
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
