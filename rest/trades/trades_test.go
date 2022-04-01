package trades_test

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

func TestListTrades(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Trades.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	trade1 := models.Trade{Price: 1.23}
	trade2 := models.Trade{Price: 1.5}
	expectedResponse := models.ListTradesResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     2,
			PaginationHooks: models.PaginationHooks{
				NextURL: "https://api.polygon.io/v3/trades/AAPL?cursor=YXA9OT",
			},
		},
		Results: []models.Trade{trade1, trade2},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/trades/AAPL?limit=2&order=asc&sort=timestamp&timestamp.gte=1626912000000000000",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	expectedNextResponse := models.ListTradesResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req2",
			Count:     0,
		},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/trades/AAPL?cursor=YXA9OT",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedNextResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	iter, err := c.Trades.ListTrades(context.Background(), models.ListTradesParams{
		Ticker:       "AAPL",
		TimestampGTE: models.Ptr(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC)),
		Order:        models.Ptr(models.Asc),
		Limit:        models.Ptr(2),
		Sort:         models.Ptr(models.Timestamp),
	})

	// verify the first page
	assert.Nil(t, err)
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Trade())
	// verify the first and second trades
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, trade1, iter.Trade())
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, trade2, iter.Trade())

	// verify the end of the list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetLastTrade(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Trades.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := models.GetLastTradeResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     1,
		},
		Results: models.LastTrade{Price: 1.23},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/last/trade/AAPL",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Trades.GetLastTrade(context.Background(), models.GetLastTradeParams{
		Ticker: "AAPL",
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}

func TestGetLastCryptoTrade(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := models.GetLastCryptoTradeResponse{
		BaseResponse: models.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
		},
		Last: models.CryptoTrade{
			Price:      26049.42,
			Size:       0.0449,
			Exchange:   4,
			Conditions: []int{1},
			Timestamp:  1605560463099,
		},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v1/last/crypto/BTC/USD",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Trades.GetLastCryptoTrade(context.Background(), models.GetLastCryptoTradeParams{
		From: "BTC",
		To:   "USD",
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedResponse, res)
}
