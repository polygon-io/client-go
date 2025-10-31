// Options - Quotes
// https://massive.com/docs/options/get_v3_quotes__optionsticker
// https://github.com/massive-com/client-go/v2/blob/master/rest/quotes.go
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
	params := models.ListQuotesParams{
		Ticker: "O:SPY241220P00720000",
	}.WithLimit(50000).WithOrder(models.Asc)

	// make request
	iter := c.ListQuotes(context.Background(), params)

	// do something with the result
	for iter.Next() {
		log.Print(iter.Item())
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}

}
