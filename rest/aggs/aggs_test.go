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

	expectedResponse := `{
	"status": "OK",
	"request_id": "6a7e466379af0a71039d60cc78e72282",
	"ticker": "AAPL",
	"queryCount": 2,
	"resultsCount": 2,
	"adjusted": true,
	"results": [
		{
			"v": 135647456,
			"vw": 74.6099,
			"o": 74.06,
			"c": 75.0875,
			"h": 75.15,
			"l": 73.7975,
			"t": 1577941200000,
			"n": 1
		},
		{
			"v": 146535512,
			"vw": 74.7026,
			"o": 74.2875,
			"c": 74.3575,
			"h": 75.145,
			"l": 74.125,
			"t": 1578027600000,
			"n": 1
		}
	]
}`

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/2021-07-22/2021-08-22?adjusted=true&explain=false&limit=1&sort=desc",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, expectedResponse)
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
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}

func TestGetGroupedDailyAggs(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"queryCount": 3,
	"resultsCount": 3,
	"adjusted": true,
	"results": [
		{
			"T": "KIMpL",
			"v": 4369,
			"vw": 26.0407,
			"o": 26.07,
			"c": 25.9102,
			"h": 26.25,
			"l": 25.91,
			"t": 1602705600000,
			"n": 74
		},
		{
			"T": "TANH",
			"v": 25933.6,
			"vw": 23.493,
			"o": 24.5,
			"c": 23.4,
			"h": 24.763,
			"l": 22.65,
			"t": 1602705600000,
			"n": 1096
		},
		{
			"T": "VSAT",
			"v": 312583,
			"vw": 34.4736,
			"o": 34.9,
			"c": 34.24,
			"h": 35.47,
			"l": 34.21,
			"t": 1602705600000,
			"n": 4966
		}
	]
}`

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/grouped/locale/us/market/stocks/2021-07-22",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, expectedResponse)
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
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}

func TestGetDailyOpenCloseAgg(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"symbol": "AAPL",
	"from": "2020-10-14T00:00:00.000Z",
	"open": 324.66,
	"high": 326.2,
	"low": 322.3,
	"close": 325.12,
	"volume": 26122646,
	"afterHours": 322.1,
	"preMarket": 324.5
}`

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v1/open-close/AAPL/2020-10-14?adjusted=true",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, expectedResponse)
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
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}

func TestGetPreviousCloseAgg(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"request_id": "6a7e466379af0a71039d60cc78e72282",
	"ticker": "AAPL",
	"queryCount": 1,
	"resultsCount": 1,
	"adjusted": true,
	"results": [
		{
			"T": "AAPL",
			"v": 131704427,
			"vw": 116.3058,
			"o": 115.55,
			"c": 115.97,
			"h": 117.59,
			"l": 114.13,
			"t": 1605042000000
		}
	]
}`

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/aggs/ticker/AAPL/prev",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, expectedResponse)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Aggs.GetPreviousCloseAgg(context.Background(), models.GetPreviousCloseAggParams{
		Ticker:   "AAPL",
		Adjusted: models.Ptr(true),
	})

	assert.Nil(t, err)
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}
