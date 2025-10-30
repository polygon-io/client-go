package main

import (
	"context"
	"log"
	"os"
	"time"

	massive "github.com/massive-com/client-go/rest"
	"github.com/massive-com/client-go/rest/models"
)

func main() {
	getAggregateBarsLaunchpad()
}

func getAggregateBarsLaunchpad() {
	c := massive.New(os.Getenv("MASSIVE_API_KEY"))

	params := &models.ListAggsParams{
		Ticker:     "CORN",
		Multiplier: 1,
		Timespan:   models.Day,
		From:       models.Millis(time.Now().AddDate(0, 0, -7)),
		To:         models.Millis(time.Now()),
	}

	iter := c.ListAggs(context.TODO(), params)
	for iter.Next() {
		log.Print(iter.Item()) // do something with the current value
	}
	if iter.Err() != nil {
		log.Fatal(iter.Err())
	}
}
