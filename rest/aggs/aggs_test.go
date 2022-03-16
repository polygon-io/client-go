package aggs_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/aggs"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/1626926400000/1629604800000?adjusted=true&limit=1&sort=desc",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
			{
				"ticker": "AAPL",
				"queryCount": 1,
				"resultsCount": 1,
				"adjusted": true,
				"results": [
				 {
				  "v": 77287356,
				  "vw": 146.991,
				  "o": 145.935,
				  "c": 146.8,
				  "h": 148.195,
				  "l": 145.81,
				  "t": 1626926400000,
				  "n": 480209
				 }
				],
				"status": "OK",
				"request_id": "cffb2db04ed53d1fdf2547f15c1ca14e",
				"count": 1
			}
`,
			)
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
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, "AAPL", res.Ticker)
}

func TestGetPreviousClose(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/ticker/AAPL/prev",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
			{
				"ticker": "AAPL",
				"queryCount": 1,
				"resultsCount": 1,
				"adjusted": true,
				"results": [
				  {
					"T": "AAPL",
					"v": 9.4493006e+07,
					"vw": 153.473,
					"o": 150.9,
					"c": 155.09,
					"h": 155.57,
					"l": 150.38,
					"t": 1647374400000,
					"n": 735965
				  }
				],
				"status": "OK",
				"request_id": "7ab4157627a1486ab072fe45f31ed808",
				"count": 1
			  }
`,
			)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)
	adjusted := true
	res, err := c.Aggs.GetPreviousClose(context.Background(), aggs.GetPreviousCloseParams{
		Ticker: "AAPL",
		QueryParams: aggs.GetPreviousCloseQueryParams{
			Adjusted: &adjusted,
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, "AAPL", res.Ticker)
}

func TestGetGroupedDaily(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/grouped/locale/us/market/stocks/2021-07-22",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
			{
				"queryCount": 1,
				"resultsCount": 1,
				"adjusted": true,
				"results": [
				  {
					"T": "CORN",
					"v": 368616,
					"vw": 13.407,
					"o": 13.35,
					"c": 13.43,
					"h": 13.46,
					"l": 13.34,
					"t": 1602705600000,
					"n": 758
					}
				  ],
				  "status": "OK",
				  "request_id": "f3c9b3358637c9a4a1308d57c2f164e3",
				  "count": 1
				}
`,
			)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)
	adjusted := true
	res, err := c.Aggs.GetGroupedDaily(context.Background(), aggs.GetGroupedDailyParams{
		Locale:     "us",
		MarketType: "stocks",
		Date:       time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local),
		QueryParams: aggs.GetGroupedDailyQueryParams{
			Adjusted: &adjusted,
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, res.QueryCount)
}

func TestGetDailyOpenClose(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v1/open-close/AAPL/2020-10-14?adjusted=true",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, `
			{
				"status": "OK",
				"from": "2020-10-14",
				"symbol": "AAPL",
				"open": 121,
				"high": 123.03,
				"low": 119.62,
				"close": 121.19,
				"volume": 151057198,
				"afterHours": 120.81,
				"preMarket": 121.55
			  }
`,
			)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)
	adjusted := true
	res, err := c.Aggs.GetDailyOpenClose(context.Background(), aggs.GetDailyOpenCloseParams{
		Ticker: "AAPL",
		Date:   time.Date(2020, 10, 14, 0, 0, 0, 0, time.Local),
		QueryParams: aggs.GetDailyOpenCloseQueryParams{
			Adjusted: &adjusted,
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, "AAPL", res.Symbol)
}
