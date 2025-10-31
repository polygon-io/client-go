// Stocks - Market Status
// https://massive.com/docs/stocks/get_v1_marketstatus_now
// https://github.com/massive-com/client-go/v2/blob/master/rest/reference.go
package main

import (
	"context"
	"log"
	"os"

	massive "github.com/massive-com/client-go/v2/rest"
)

func main() {

	// init client
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	// make request
	res, err := c.GetMarketStatus(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)

}
