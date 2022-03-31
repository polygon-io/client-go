package quotes_test

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

func TestListQuotes(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	quote1 := models.Quote{AskPrice: 1.23}
	quote2 := models.Quote{AskPrice: 1.5}
	expectedResponse := models.QuotesResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     2,
			PaginationHooks: client.PaginationHooks{
				NextURL: "https://api.polygon.io/v3/quotes/AAPL?cursor=YXA9OT",
			},
		},
		Results: []models.Quote{quote1, quote2},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/quotes/AAPL?limit=2&order=asc&sort=timestamp&timestamp.lte=1626912000000000000",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	quote3 := models.Quote{AskPrice: 1.40}
	expectedNextResponse := models.QuotesResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req2",
			Count:     1,
		},
		Results: []models.Quote{quote3},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/quotes/AAPL?cursor=YXA9OT",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedNextResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	iter, err := c.Quotes.ListQuotes(context.Background(), models.ListQuotesParams{
		Ticker:       "AAPL",
		TimestampLTE: models.Ptr(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC)),
		Order:        models.Ptr(models.Asc),
		Limit:        models.Ptr(2),
		Sort:         models.Ptr(models.Timestamp),
	})

	// verify the first page
	assert.Nil(t, err)
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Quote())
	// verify the first and second quotes
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, quote1, iter.Quote())
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, quote2, iter.Quote())

	// verify the second page
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	// verify the third quote
	assert.Equal(t, quote3, iter.Quote())

	// verify the end of the list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetLastQuote(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := models.LastQuoteResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     1,
		},
		Results: models.LastQuote{AskPrice: 1.23},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/last/nbbo/AAPL",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Quotes.GetLastQuote(context.Background(), models.GetLastQuoteParams{
		Ticker: "AAPL",
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}

func TestGetLastForexQuote(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := models.LastForexQuoteResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
		},
		Last: models.ForexQuote{
			Ask:       1.23,
			Bid:       1.24,
			Exchange:  5,
			Timestamp: 1626912000000000000,
		},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v1/last_quote/currencies/USD/GBP",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Quotes.GetLastForexQuote(context.Background(), models.LastForexQuoteParams{
		From: "USD",
		To:   "GBP",
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}
