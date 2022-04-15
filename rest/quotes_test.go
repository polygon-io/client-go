package polygon_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListQuotes(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	quote1 := `{
	"ask_exchange": 10,
	"ask_price": 103.3,
	"ask_size": 60,
	"bid_exchange": 11,
	"bid_price": 102.7,
	"bid_size": 60,
	"conditions": [
		1
	],
	"participant_timestamp": 1517562000065321200,
	"sequence_number": 2060,
	"sip_timestamp": 1517562000065700400,
	"tape": 3
}`

	quote2 := `{
	"ask_exchange": 10,
	"ask_price": 180,
	"ask_size": 2,
	"bid_exchange": 11,
	"bid_price": 170,
	"bid_size": 2,
	"conditions": [
		1
	],
	"participant_timestamp": 1517562000065408300,
	"sequence_number": 2061,
	"sip_timestamp": 1517562000065791500,
	"tape": 3
}`

	expectedResponse := `{
	"status": "OK",
	"request_id": "a47d1beb8c11b6ae897ab76cdbbf35a3",
	"next_url": "https://api.polygon.io/v3/quotes/AAPL?cursor=YWN0aXZlPXRydWUmZGF0ZT0yMDIxLTA0LTI1JmxpbWl0PTEmb3JkZXI9YXNjJnBhZ2VfbWFya2VyPUElN0M5YWRjMjY0ZTgyM2E1ZjBiOGUyNDc5YmZiOGE1YmYwNDVkYzU0YjgwMDcyMWE2YmI1ZjBjMjQwMjU4MjFmNGZiJnNvcnQ9dGlja2Vy",
	"results": [
` + indent(true, quote1, "\t\t") + `,
` + indent(true, quote2, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/quotes/AAPL?limit=2&order=asc&sort=timestamp&timestamp=1626912000000000000", expectedResponse)
	registerResponder("https://api.polygon.io/v3/quotes/AAPL?cursor=YWN0aXZlPXRydWUmZGF0ZT0yMDIxLTA0LTI1JmxpbWl0PTEmb3JkZXI9YXNjJnBhZ2VfbWFya2VyPUElN0M5YWRjMjY0ZTgyM2E1ZjBiOGUyNDc5YmZiOGE1YmYwNDVkYzU0YjgwMDcyMWE2YmI1ZjBjMjQwMjU4MjFmNGZiJnNvcnQ9dGlja2Vy", "{}")
	iter := c.Quotes.ListQuotes(context.Background(), models.ListQuotesParams{Ticker: "AAPL"}.
		WithTimestamp(models.EQ, models.Nanos(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).
		WithSort(models.Timestamp).WithOrder(models.Asc).WithLimit(2))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err := json.MarshalIndent(iter.Item(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, quote1, string(b))

	// second item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err = json.MarshalIndent(iter.Item(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, quote2, string(b))

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetLastQuote(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"request_id": "b84e24636301f19f88e0dfbf9a45ed5c",
	"results": {
		"T": "AAPL",
		"X": 19,
		"P": 127.98,
		"S": 7,
		"x": 11,
		"p": 127.96,
		"s": 1,
		"y": 1617827221349366000,
		"q": 83480742,
		"t": 1617827221349730300,
		"z": 3
	}
}`

	registerResponder("https://api.polygon.io/v2/last/nbbo/AAPL", expectedResponse)
	res, err := c.Quotes.GetLastQuote(context.Background(), &models.GetLastQuoteParams{
		Ticker: "AAPL",
	})

	assert.Nil(t, err)
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}

func TestGetLastForexQuote(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "success",
	"request_id": "a73a29dbcab4613eeaf48583d3baacf0",
	"symbol": "AUD/USD",
	"last": {
		"ask": 0.73124,
		"bid": 0.73122,
		"exchange": 48,
		"timestamp": 1605557756000
	}
}`

	registerResponder("https://api.polygon.io/v1/last_quote/currencies/USD/GBP", expectedResponse)
	res, err := c.Quotes.GetLastForexQuote(context.Background(), &models.GetLastForexQuoteParams{
		From: "USD",
		To:   "GBP",
	})

	assert.Nil(t, err)
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}
