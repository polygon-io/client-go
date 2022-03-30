package reference_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListTickers(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	ticker1 := models.TickerDetails{Ticker: "A", Name: "Agilent Technologies Inc."}
	ticker2 := models.TickerDetails{Ticker: "AA", Name: "Alcoa Corporation"}
	expectedResponse := models.TickersResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     2,
			PaginationHooks: client.PaginationHooks{
				NextURL: "https://api.polygon.io/v3/reference/tickers?cursor=YXA9OT",
			},
		},
		Results: []*models.TickerDetails{&ticker1, &ticker2},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/reference/tickers?active=true&cik=5&cusip=10&date=2021-07-22&exchange=4&limit=2&market=stocks&order=asc&sort=ticker&type=CS",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	expectedNextResponse := models.TickersResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req2",
			Count:     0,
		},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/reference/tickers?cursor=YXA9OT",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedNextResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	iter := c.Reference.ListTickers(context.Background(), models.ListTickersParams{
		QueryParams: models.ListTickersQueryParams{
			Type:     models.Ptr("CS"),
			Market:   models.Ptr(models.Stocks),
			Exchange: models.Ptr(strconv.FormatInt(4, 10)),
			CUSIP:    models.Ptr(strconv.FormatInt(10, 10)),
			CIK:      models.Ptr(strconv.FormatInt(5, 10)),
			Date:     models.Ptr(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC)),
			Active:   models.Ptr(true),
			Sort:     models.Ptr(models.Ticker),
			Order:    models.Ptr(models.Asc),
			Limit:    models.Ptr(2),
		},
	})

	// verify the first page
	assert.Nil(t, iter.Err())
	assert.Nil(t, iter.Ticker())
	// verify the first and second trades
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, &ticker1, iter.Ticker())
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, &ticker2, iter.Ticker())

	// verify the end of the list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetTickerDetails(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	ticker1 := models.TickerDetails{Ticker: "A", Name: "Agilent Technologies Inc."}
	expectedResponse := models.TickersResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     1,
		},
		Results: []*models.TickerDetails{&ticker1},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/reference/tickers/A?date=2021-07-22",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Reference.GetTickerDetails(context.Background(), models.GetTickerDetailsParams{
		Ticker: "A",
		QueryParams: models.GetTickerDetailsQueryParams{
			Date: models.Ptr(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC)),
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}

func TestGetTickerTypes(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := models.TickerTypesResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     1,
		},
		Results: []*models.TickerType{
			{
				AssetClass:  "stocks",
				Code:        "CS",
				Description: "Common Stock",
				Locale:      "us",
			},
		},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/reference/tickers/types?asset_class=stocks&locale=us",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Reference.GetTickerTypes(context.Background(), models.GetTickerTypesParams{
		QueryParams: models.GetTickerTypesQueryParams{
			AssetClass: models.Ptr("stocks"),
			Locale:     models.Ptr(models.US),
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}
