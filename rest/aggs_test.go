package polygon_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

const (
	expectedAggsResponseURL = "https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/1626912000000/1629590400000?adjusted=true&limit=2&sort=desc"

	agg1 = `{
	"v": 135647456,
	"vw": 74.6099,
	"o": 74.06,
	"c": 75.0875,
	"h": 75.15,
	"l": 73.7975,
	"t": 1577941200000,
	"n": 1
}`

	agg2 = `{
	"v": 146535512,
	"vw": 74.7026,
	"o": 74.2875,
	"c": 74.3575,
	"h": 75.145,
	"l": 74.125,
	"t": 1578027600000,
	"n": 1
}`
)

var (
	expectedAggsResponse = `{
	"status": "OK",
	"request_id": "6a7e466379af0a71039d60cc78e72282",
	"next_url": "https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/1626912000000/1629590400000?cursor=AGGSCURSOR",
	"ticker": "AAPL",
	"queryCount": 2,
	"resultsCount": 2,
	"adjusted": true,
	"results": [
` + indent(true, agg1, "\t\t") + `,
` + indent(true, agg2, "\t\t") + `
	]
}`
)

func TestListAggs(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()
	registerResponder(expectedAggsResponseURL, expectedAggsResponse)
	registerResponder("https://api.polygon.io/v2/aggs/ticker/AAPL/range/1/day/1626912000000/1629590400000?cursor=AGGSCURSOR", "{}")

	iter := c.ListAggs(context.Background(), models.ListAggsParams{
		Ticker:     "AAPL",
		Multiplier: 1,
		Timespan:   "day",
		From:       models.Millis(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC)),
		To:         models.Millis(time.Date(2021, 8, 22, 0, 0, 0, 0, time.UTC)),
	}.WithOrder(models.Desc).WithLimit(2).WithAdjusted(true))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect1 models.Agg
	err := json.Unmarshal([]byte(agg1), &expect1)
	assert.Nil(t, err)
	assert.Equal(t, expect1, iter.Item())

	// second item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect2 models.Agg
	err = json.Unmarshal([]byte(agg2), &expect2)
	assert.Nil(t, err)
	assert.Equal(t, expect2, iter.Item())

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetAggs(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()
	registerResponder(expectedAggsResponseURL, expectedAggsResponse)

	res, err := c.GetAggs(context.Background(), models.GetAggsParams{
		Ticker:     "AAPL",
		Multiplier: 1,
		Timespan:   "day",
		From:       models.Millis(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC)),
		To:         models.Millis(time.Date(2021, 8, 22, 0, 0, 0, 0, time.UTC)),
	}.WithOrder(models.Desc).WithLimit(2).WithAdjusted(true))
	assert.Nil(t, err)

	var expect models.GetAggsResponse
	err = json.Unmarshal([]byte(expectedAggsResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetAggsWithQueryParam(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()
	registerResponder(expectedAggsResponseURL, expectedAggsResponse)

	res, err := c.GetAggs(context.Background(), models.GetAggsParams{
		Ticker:     "AAPL",
		Multiplier: 1,
		Timespan:   "day",
		From:       models.Millis(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC)),
		To:         models.Millis(time.Date(2021, 8, 22, 0, 0, 0, 0, time.UTC)),
	}.WithOrder(models.Desc).WithLimit(2), models.QueryParam("adjusted", "true"))
	assert.Nil(t, err)

	var expect models.GetAggsResponse
	err = json.Unmarshal([]byte(expectedAggsResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetGroupedDailyAggs(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
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

	registerResponder("https://api.polygon.io/v2/aggs/grouped/locale/us/market/stocks/2021-07-22", expectedResponse)
	res, err := c.GetGroupedDailyAggs(context.Background(), models.GetGroupedDailyAggsParams{
		Locale:     models.US,
		MarketType: models.Stocks,
		Date:       models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.Local)),
	}.WithAdjusted(true))
	assert.Nil(t, err)

	var expect models.GetGroupedDailyAggsResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetDailyOpenCloseAgg(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
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

	registerResponder("https://api.polygon.io/v1/open-close/AAPL/2020-10-14?adjusted=true", expectedResponse)
	res, err := c.GetDailyOpenCloseAgg(context.Background(), models.GetDailyOpenCloseAggParams{
		Ticker: "AAPL",
		Date:   models.Date(time.Date(2020, 10, 14, 0, 0, 0, 0, time.Local)),
	}.WithAdjusted(true))
	assert.Nil(t, err)

	var expect models.GetDailyOpenCloseAggResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetPreviousCloseAgg(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
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

	registerResponder("https://api.polygon.io/v2/aggs/ticker/AAPL/prev", expectedResponse)
	res, err := c.GetPreviousCloseAgg(context.Background(), models.GetPreviousCloseAggParams{
		Ticker: "AAPL",
	}.WithAdjusted(true))
	assert.Nil(t, err)

	var expect models.GetPreviousCloseAggResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func registerResponder(url, body string) {
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, body)
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)
}

func indent(first bool, data, indent string) string {
	lines := strings.Split(data, "\n")
	for i := range lines {
		if i == 0 && !first {
			continue
		}
		lines[i] = indent + lines[i]
	}
	return strings.Join(lines, "\n")
}
