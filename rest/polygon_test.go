package polygon_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/aggs"
)

// todo: write some tests, just verifying that the client works for now

func TestAggs(t *testing.T) {
	c := polygon.New(os.Getenv("API_KEY"))
	res, err := c.Aggs.Get(context.Background(), aggs.GetParams{
		Ticker:     "AAPL",
		Multiplier: 1,
		Resolution: "day",
		From:       time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local),
		To:         time.Date(2021, 8, 22, 0, 0, 0, 0, time.Local),
		QueryParams: aggs.GetQueryParams{
			Adjusted: polygon.Bool(true),
			Sort:     polygon.AggsSort(aggs.Desc),
			Limit:    polygon.Int32(10),
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
	c := polygon.New(os.Getenv("API_KEY"))
	res, err := c.Aggs.GetPreviousClose(context.Background(), aggs.GetPreviousCloseParams{
		Ticker: "AAPL",
		QueryParams: aggs.GetPreviousCloseQueryParams{
			Adjusted: polygon.Bool(true),
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
	c := polygon.New(os.Getenv("API_KEY"))
	res, err := c.Aggs.GetGroupedDaily(context.Background(), aggs.GetGroupedDailyParams{
		Locale:     "global",
		MarketType: aggs.Crypto,
		Date:       time.Date(2021, 7, 21, 0, 0, 0, 0, time.Local),
		QueryParams: aggs.GetGroupedDailyQueryParams{
			Adjusted: polygon.Bool(true),
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
	c := polygon.New(os.Getenv("API_KEY"))
	res, err := c.Aggs.GetDailyOpenClose(context.Background(), aggs.GetDailyOpenCloseParams{
		Ticker: "AAPL",
		Date:   time.Date(2021, 7, 21, 0, 0, 0, 0, time.Local),
		QueryParams: aggs.GetDailyOpenCloseQueryParams{
			Adjusted: polygon.Bool(true),
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
