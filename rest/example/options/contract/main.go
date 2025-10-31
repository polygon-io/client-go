// Options - Options Contract
// https://massive.com/docs/options/get_v3_reference_options_contracts__options_ticker
// https://github.com/massive-com/client-go/v2/blob/master/rest/reference.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	massive "github.com/massive-com/client-go/v2/rest"
	"github.com/massive-com/client-go/v2/rest/models"
)

func main() {

	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// set params
	params := models.GetOptionsContractParams{
		Ticker: "O:EVRI240119C00002500",
	}.WithAsOf(models.Date(time.Date(2022, 5, 16, 0, 0, 0, 0, time.Local)))

	// make request
	res, err := c.GetOptionsContract(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
