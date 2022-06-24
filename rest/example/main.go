package main

import (
	"context"
	"log"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {
	getAllTickersSnapshot()
	listTrades()
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
