// Options - Options Contracts
// https://polygon.io/docs/options/get_v3_reference_options_contracts
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
	params := models.ListOptionsContractsParams{}.
		WithUnderlyingTicker(models.EQ, "HCP").
		WithContractType("call").
		WithLimit(1000)

	// make request
	iter := c.ListOptionsContracts(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
