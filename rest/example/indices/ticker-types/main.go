// Indices - Ticker Types
// https://massive.com/docs/indices/get_v3_reference_tickers_types
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
	params := models.GetTickerTypesParams{}.WithAssetClass("indices")

	// make request
	res, err := c.GetTickerTypes(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
