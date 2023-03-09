// options - options contract
// https://polygon.io/docs/options/get_v3_reference_options_contracts__options_ticker
// https://github.com/polygon-io/client-go/blob/master/rest/reference.go
package main

import (
	"context"
	"time"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {

	// init client
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

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