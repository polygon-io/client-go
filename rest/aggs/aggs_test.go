package aggs_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

var expectedResponse = models.AggsResponse{
	Ticker: "AAPL",
	BaseResponse: client.BaseResponse{
		Status:       "OK",
		RequestID:    "cffb2db04ed53d1fdf2547f15c1ca14e",
		Count:        1,
		Message:      "",
		ErrorMessage: "",
	},
	QueryCount:   1,
	ResultsCount: 1,
	Adjusted:     true,
	Aggs: []models.Aggregate{
		{
			Volume:       77287356,
			VWAP:         146.991,
			Open:         145.935,
			Close:        146.8,
			High:         148.195,
			Low:          145.81,
			Timestamp:    1626926400000,
			Transactions: 480209,
		},
	},
}

func TestGetAggs(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/1626912000000/1629590400000?adjusted=true&explain=false&limit=1&sort=desc",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Aggs.GetAggs(context.Background(), models.GetAggsParams{
		Ticker:     "AAPL",
		Multiplier: 1,
		Resolution: "day",
		From:       time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC),
		To:         time.Date(2021, 8, 22, 0, 0, 0, 0, time.UTC),
		QueryParams: models.GetAggsQueryParams{
			Adjusted: models.Bool(true),
			Sort:     models.SortOrder(models.Desc),
			Limit:    models.Int(1),
			Explain:  models.Bool(false),
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

	res, err := c.Aggs.GetPreviousClose(context.Background(), models.GetPreviousCloseParams{
		Ticker: "AAPL",
		QueryParams: models.GetPreviousCloseQueryParams{
			Adjusted: models.Bool(true),
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

	res, err := c.Aggs.GetGroupedDaily(context.Background(), models.GetGroupedDailyParams{
		Locale:     models.US,
		MarketType: models.Stocks,
		Date:       time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local),
		QueryParams: models.GetGroupedDailyQueryParams{
			Adjusted: models.Bool(true),
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

	expectedResponse := models.DailyOpenCloseResponse{
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

	res, err := c.Aggs.GetDailyOpenClose(context.Background(), models.GetDailyOpenCloseParams{
		Ticker: "AAPL",
		Date:   time.Date(2020, 10, 14, 0, 0, 0, 0, time.Local),
		QueryParams: models.GetDailyOpenCloseQueryParams{
			Adjusted: models.Bool(true),
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}
