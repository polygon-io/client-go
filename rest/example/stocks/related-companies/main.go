// Stocks - Related Companies
// https://polygon.io/docs/stocks/get_v1_related-companies__ticker
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
	params := models.GetTickerRelatedCompaniesParams{
		Ticker: "AAPL",
	}

	// make request
	res, err := c.GetTickerRelatedCompanies(context.Background(), &params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
