// Crypto - Exchanges
// https://polygon.io/docs/crypto/get_v3_reference_exchanges
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
	params := models.GetExchangesParams{}.
		WithAssetClass(models.AssetCrypto)

	// make request
	res, err := c.GetExchanges(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res) 
}
