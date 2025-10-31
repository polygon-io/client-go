// Indices - Market Holidays
// https://massive.com/docs/indices/get_v1_marketstatus_upcoming
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
	res, err := c.GetMarketHolidays(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// do something with the result
	log.Print(res)
}
