// Options - Options Contracts
// https://massive.com/docs/options/get_v3_reference_options_contracts
// https://github.com/massive-com/client-go/blob/master/rest/reference.go
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
