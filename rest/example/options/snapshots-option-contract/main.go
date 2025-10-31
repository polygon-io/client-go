// Options - Option Contract
// https://massive.com/docs/options/get_v3_snapshot_options__underlyingasset___optioncontract
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
	params := &models.GetOptionContractSnapshotParams{
		UnderlyingAsset: "AAPL",
		OptionContract:  "O:AAPL230616C00150000",
	}

	// make request
	res, err := c.GetOptionContractSnapshot(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
