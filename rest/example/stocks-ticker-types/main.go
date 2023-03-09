// stocks - ticker types
// https://polygon.io/docs/stocks/get_v3_reference_tickers_types
// https://github.com/polygon-io/client-go/blob/master/rest/reference.go
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
	params := models.GetTickerTypesParams{}.
		WithAssetClass("stocks").
		WithLocale(models.US)

	// make request
	res, err := c.GetTickerTypes(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res) 

}