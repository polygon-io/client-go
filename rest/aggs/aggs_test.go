package aggs_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAggs(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := models.GetAggsResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     1,
		},
		Results: []models.Agg{{Volume: 77287356}},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/2021-07-22/2021-08-22?adjusted=true&explain=false&limit=1&sort=desc",
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
		Adjusted:   models.Ptr(true),
		Sort:       models.Ptr(models.Desc),
		Limit:      models.Ptr(1),
		Explain:    models.Ptr(false),
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}

func TestGetGroupedDailyAggs(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := models.GetGroupedDailyAggsResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     1,
		},
		Results: []models.Agg{{Volume: 77287356}},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/grouped/locale/us/market/stocks/2021-07-22",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Aggs.GetGroupedDailyAggs(context.Background(), models.GetGroupedDailyAggsParams{
		Locale:     models.US,
		MarketType: models.Stocks,
		Date:       time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local),
		Adjusted:   models.Ptr(true),
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}

func TestGetPreviousCloseAgg(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := models.GetPreviousCloseAggResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     1,
		},
		Results: []models.Agg{{Volume: 77287356}},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/ticker/AAPL/prev",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Aggs.GetPreviousCloseAgg(context.Background(), models.GetPreviousCloseAggParams{
		Ticker:   "AAPL",
		Adjusted: models.Ptr(true),
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}

func TestGetDailyOpenCloseAgg(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := models.GetDailyOpenCloseAggResponse{
		BaseResponse: models.BaseResponse{
			Status: "OK",
		},
		Symbol:     "AAPL",
		From:       "2020-10-14",
		Open:       121,
		High:       123.03,
		Low:        119.62,
		Close:      121.19,
		Volume:     151057198,
		AfterHours: 120.81,
		PreMarket:  121.55,
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

	res, err := c.Aggs.GetDailyOpenCloseAgg(context.Background(), models.GetDailyOpenCloseAggParams{
		Ticker:   "AAPL",
		Date:     time.Date(2020, 10, 14, 0, 0, 0, 0, time.Local),
		Adjusted: models.Ptr(true),
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}
