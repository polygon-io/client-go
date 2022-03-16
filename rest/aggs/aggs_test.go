package aggs_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/aggs"
	"github.com/polygon-io/client-go/rest/client"
	"github.com/stretchr/testify/assert"
)

var expectedAggregate = aggs.Aggregate{
	Volume:       77287356,
	VWAP:         146.991,
	Open:         145.935,
	Close:        146.8,
	High:         148.195,
	Low:          145.81,
	Timestamp:    1626926400000,
	Transactions: 480209,
}

var expectedBaseResponse = client.BaseResponse{
	Status:       "OK",
	RequestID:    "cffb2db04ed53d1fdf2547f15c1ca14e",
	Count:        1,
	Message:      "",
	ErrorMessage: "",
}

var expectedResponse = aggs.AggsResponse{
	Ticker:       "AAPL",
	BaseResponse: expectedBaseResponse,
	QueryCount:   1,
	ResultsCount: 1,
	Adjusted:     true,
	Aggs:         []aggs.Aggregate{expectedAggregate},
}

func TestGet(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/1626926400000/1629604800000?adjusted=true&explain=false&limit=1&sort=desc",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Aggs.Get(context.Background(), aggs.GetParams{
		Ticker:     "AAPL",
		Multiplier: 1,
		Resolution: "day",
		From:       time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local),
		To:         time.Date(2021, 8, 22, 0, 0, 0, 0, time.Local),
		QueryParams: aggs.GetQueryParams{
			Adjusted: polygon.Bool(true),
			Sort:     polygon.AggsSort(aggs.Desc),
			Limit:    polygon.Int32(1),
			Explain:  polygon.Bool(false),
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}

func TestGetPreviousClose(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/ticker/AAPL/prev",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Aggs.GetPreviousClose(context.Background(), aggs.GetPreviousCloseParams{
		Ticker: "AAPL",
		QueryParams: aggs.GetPreviousCloseQueryParams{
			Adjusted: polygon.Bool(true),
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}

func TestGetGroupedDaily(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/grouped/locale/us/market/stocks/2021-07-22",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Aggs.GetGroupedDaily(context.Background(), aggs.GetGroupedDailyParams{
		Locale:     "us",
		MarketType: "stocks",
		Date:       time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local),
		QueryParams: aggs.GetGroupedDailyQueryParams{
			Adjusted: polygon.Bool(true),
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}

func TestGetDailyOpenClose(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedBaseResponse := client.BaseResponse{
		Status: "OK",
	}

	expectedResponse := aggs.DailyOpenCloseResponse{
		BaseResponse: expectedBaseResponse,
		Symbol:       "AAPL",
		From:         "2020-10-14",
		Open:         121,
		High:         123.03,
		Low:          119.62,
		Close:        121.19,
		Volume:       151057198,
		AfterHours:   120.81,
		PreMarket:    121.55,
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v1/open-close/AAPL/2020-10-14?adjusted=true",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Aggs.GetDailyOpenClose(context.Background(), aggs.GetDailyOpenCloseParams{
		Ticker: "AAPL",
		Date:   time.Date(2020, 10, 14, 0, 0, 0, 0, time.Local),
		QueryParams: aggs.GetDailyOpenCloseQueryParams{
			Adjusted: polygon.Bool(true),
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}
