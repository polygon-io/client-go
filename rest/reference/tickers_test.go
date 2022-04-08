package reference_test

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListTickers(t *testing.T) {
	c := polygon.NewClient("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	ticker1 := `{
	"ticker": "A",
	"name": "Agilent Technologies Inc.",
	"market": "stocks",
	"locale": "us",
	"primary_exchange": "XNYS",
	"type": "CS",
	"active": true,
	"currency_name": "usd",
	"cik": "0001090872",
	"composite_figi": "BBG000BWQYZ5",
	"share_class_figi": "BBG001SCTQY4",
	"last_updated_utc": "2021-04-25T00:00:00Z"
}`

	expectedResponse := `{
	"status": "OK",
	"count": 1,
	"next_url": "https://api.polygon.io/v3/reference/tickers?cursor=YWN0aXZlPXRydWUmZGF0ZT0yMDIxLTA0LTI1JmxpbWl0PTEmb3JkZXI9YXNjJnBhZ2VfbWFya2VyPUElN0M5YWRjMjY0ZTgyM2E1ZjBiOGUyNDc5YmZiOGE1YmYwNDVkYzU0YjgwMDcyMWE2YmI1ZjBjMjQwMjU4MjFmNGZiJnNvcnQ9dGlja2Vy",
	"request_id": "e70013d92930de90e089dc8fa098888e",
	"results": [
` + indent(true, ticker1, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/tickers?active=true&cik=5&cusip=10&date=2021-07-22&exchange=4&limit=2&market=stocks&order=asc&sort=ticker&type=CS", expectedResponse)
	registerResponder("https://api.polygon.io/v3/reference/tickers?cursor=YWN0aXZlPXRydWUmZGF0ZT0yMDIxLTA0LTI1JmxpbWl0PTEmb3JkZXI9YXNjJnBhZ2VfbWFya2VyPUElN0M5YWRjMjY0ZTgyM2E1ZjBiOGUyNDc5YmZiOGE1YmYwNDVkYzU0YjgwMDcyMWE2YmI1ZjBjMjQwMjU4MjFmNGZiJnNvcnQ9dGlja2Vy", "{}")
	iter, err := c.Reference.ListTickers(context.Background(), models.ListTickersParams{}.
		WithType("CS").WithMarket(models.AssetStocks).
		WithExchange(4).WithCUSIP(10).WithCIK(5).
		WithDate(models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).WithActive(true).
		WithSort(models.TickerSymbol).WithOrder(models.Asc).WithLimit(2))

	// iter creation
	assert.Nil(t, err)
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Ticker())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err := json.MarshalIndent(iter.Ticker(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, ticker1, string(b))

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetTickerDetails(t *testing.T) {
	c := polygon.NewClient("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"request_id": "31d59dda-80e5-4721-8496-d0d32a654afe",
	"results": {
		"ticker": "AAPL",
		"name": "Apple Inc.",
		"market": "stocks",
		"locale": "us",
		"primary_exchange": "XNAS",
		"type": "CS",
		"active": true,
		"currency_name": "usd",
		"cik": "0000320193",
		"composite_figi": "BBG000B9XRY4",
		"share_class_figi": "BBG001S5N8V8",
		"last_updated_utc": "2021-04-25T00:00:00Z",
		"market_cap": 2771126040150,
		"phone_number": "(408) 996-1010",
		"address": {
			"address1": "One Apple Park Way",
			"city": "Cupertino",
			"state": "CA"
		},
		"description": "Apple designs a wide variety of consumer electronic devices, including smartphones (iPhone), tablets (iPad), PCs (Mac), smartwatches (Apple Watch), AirPods, and TV boxes (Apple TV), among others. The iPhone makes up the majority of Apple's total revenue. In addition, Apple offers its customers a variety of services such as Apple Music, iCloud, Apple Care, Apple TV+, Apple Arcade, Apple Card, and Apple Pay, among others. Apple's products run internally developed software and semiconductors, and the firm is well known for its integration of hardware, software and services. Apple's products are distributed online as well as through company-owned stores and third-party retailers. The company generates roughly 40% of its revenue from the Americas, with the remainder earned internationally.",
		"sic_code": "3571",
		"sic_description": "ELECTRONIC COMPUTERS",
		"homepage_url": "https://www.apple.com",
		"total_employees": 154000,
		"list_date": "1980-12-12",
		"branding": {
			"logo_url": "https://api.polygon.io/v1/reference/company-branding/d3d3LmFwcGxlLmNvbQ/images/2022-01-10_logo.svg",
			"icon_url": "https://api.polygon.io/v1/reference/company-branding/d3d3LmFwcGxlLmNvbQ/images/2022-01-10_icon.png"
		},
		"share_class_shares_outstanding": 16406400000,
		"weighted_shares_outstanding": 16334371000
	}
}`

	registerResponder("https://api.polygon.io/v3/reference/tickers/A?date=2021-07-22", expectedResponse)
	res, err := c.Reference.GetTickerDetails(context.Background(), models.GetTickerDetailsParams{
		Ticker: "A",
	}.WithDate(models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))))

	assert.Nil(t, err)
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}

func TestGetTickerTypes(t *testing.T) {
	c := polygon.NewClient("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"request_id": "31d59dda-80e5-4721-8496-d0d32a654afe",
	"count": 1,
	"results": [
		{
			"asset_class": "stocks",
			"code": "CS",
			"description": "Common Stock",
			"locale": "us"
		}
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/tickers/types?asset_class=stocks&locale=us", expectedResponse)
	res, err := c.Reference.GetTickerTypes(context.Background(), models.GetTickerTypesParams{}.WithAssetClass("stocks").WithLocale(models.US))

	assert.Nil(t, err)
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
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
