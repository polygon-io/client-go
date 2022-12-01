package main

import (
	"context"
	"log"
	"os"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	// By default, the common use case example will run.
	// If you want to run the Launchpad example then from the root dir run `POLYGON_API_KEY=... go run rest/example/main.go launchpad`
	exampleToRun := "common"
	if len(os.Args) > 1 && os.Args[1] == "launchpad" {
		exampleToRun = "launchpad"
	}

	switch exampleToRun {
	case "common":
		getAllTickersSnapshot()
		listTrades()
	case "launchpad":
		getAggregateBarsLaunchpad()
	}
}

func getAllTickersSnapshot() {
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	params := models.GetAllTickersSnapshotParams{
		Locale:     models.US,
		MarketType: models.Stocks,
	}.WithTickers("AAPL,MSFT")

	res, err := c.GetAllTickersSnapshot(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res) // do something with the result
}

func listTrades() {
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	params2 := models.ListTradesParams{
		Ticker: "CORN",
	}.WithDay(2021, 7, 22).WithLimit(50000).WithOrder(models.Asc)

	iter := c.ListTrades(context.Background(), params2)
	for iter.Next() {
		log.Print(iter.Item()) // do something with the result
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
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
