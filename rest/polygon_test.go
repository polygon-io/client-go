package polygon_test

import (
	"context"
	"encoding/json"
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

	res, err := c.Aggregates.Get(context.Background(), aggregates.GetParams{
		Ticker:     "AAPL",
		Multiplier: 1,
		Resolution: "day",
		From:       time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local),
		To:         time.Date(2021, 8, 22, 0, 0, 0, 0, time.Local),
		QueryParams: &aggregates.GetQueryParams{
			Adjusted: true,
			Sort:     "asc",
			Limit:    10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(res)
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

	res, err := c.Aggregates.GetPreviousClose(context.Background(), aggregates.GetPreviousCloseParams{
		Ticker: "AAPL",
		QueryParams: &aggregates.GetPreviousCloseQueryParams{
			Adjusted: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}

func TestAggsGroupedDaily(t *testing.T) {
	c := polygon.New(client.HTTPBaseConfig{
		URL:        "https://api.polygon.io",
		Key:        os.Getenv("API_KEY"),
		MaxRetries: 3,
	})

	res, err := c.Aggregates.GetGroupedDaily(context.Background(), aggregates.GetGroupedDailyParams{
		Locale:     "global",
		MarketType: "crypto",
		Date:       time.Date(2021, 7, 21, 0, 0, 0, 0, time.Local),
		QueryParams: &aggregates.GetGroupedDailyQueryParams{
			Adjusted: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}

func TestAggsDailyOpenClose(t *testing.T) {
	c := polygon.New(client.HTTPBaseConfig{
		URL:        "https://api.polygon.io",
		Key:        os.Getenv("API_KEY"),
		MaxRetries: 3,
	})

	res, err := c.Aggregates.GetDailyOpenClose(context.Background(), aggregates.GetDailyOpenCloseParams{
		Ticker: "AAPL",
		Date:   time.Date(2021, 7, 21, 0, 0, 0, 0, time.Local),
		QueryParams: &aggregates.GetDailyOpenCloseQueryParams{
			Adjusted: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
