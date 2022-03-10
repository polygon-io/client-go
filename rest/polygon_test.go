package polygon_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	polygon "github.com/polygon-io/client-golang/rest"
	"github.com/polygon-io/client-golang/rest/aggregates"
	"github.com/polygon-io/client-golang/rest/client"
)

// todo: write some tests, just verifying that the client works for now

func TestAggs(t *testing.T) {
	c := polygon.New(client.HTTPBaseConfig{
		URL:        "https://api.polygon.io",
		Key:        os.Getenv("API_KEY"),
		MaxRetries: 3,
	})

	pathParams := aggregates.GetPathParams{
		Ticker:     "AAPL",
		Multiplier: 1,
		Resolution: "day",
		From:       time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local),
		To:         time.Date(2021, 8, 22, 0, 0, 0, 0, time.Local),
	}
	queryParams := &aggregates.GetQueryParams{
		Adjusted: true,
		Sort:     "asc",
		Limit:    10,
	}

	res, err := c.Aggregates.Get(context.Background(), pathParams, queryParams)
	if err != nil {
		t.Fatal(err)
	}

	b, err := res.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}

func TestAggsPreviousClose(t *testing.T) {
	c := polygon.New(client.HTTPBaseConfig{
		URL:        "https://api.polygon.io",
		Key:        os.Getenv("API_KEY"),
		MaxRetries: 3,
	})

	pathParams := aggregates.GetPreviousClosePathParams{
		Ticker: "AAPL",
	}
	queryParams := &aggregates.GetPreviousCloseQueryParams{
		Adjusted: true,
	}

	res, err := c.Aggregates.GetPreviousClose(context.Background(), pathParams, queryParams)
	if err != nil {
		t.Fatal(err)
	}

	b, err := res.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
