package main

import (
	"context"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"log"
	"os"
	"time"
)

func main() {
	getAggregateBarsLaunchpad()
}

func getAggregateBarsLaunchpad() {
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	params3 := &models.GetAggsParams{
		Ticker:     "CORN",
		Multiplier: 1,
		Timespan:   models.Day,
		From:       models.Millis(time.Now().AddDate(0, 0, -7)),
		To:         models.Millis(time.Now()),
	}

	res, err := c.GetAggs(context.Background(), params3,
		models.RequiredEdgeHeaders("EDGE_USER_ID", "EDGE_USER_IP_ADDRESS"),
		models.EdgeUserAgent("EDGE_USER_AGENT"),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res) // do something with the result
}
