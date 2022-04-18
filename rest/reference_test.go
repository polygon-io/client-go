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

func TestListTickers(t *testing.T) {
	c := polygon.New("API_KEY")

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
	iter := c.Reference.ListTickers(context.Background(), models.ListTickersParams{}.
		WithType("CS").WithMarket(models.AssetStocks).
		WithExchange(4).WithCUSIP(10).WithCIK(5).
		WithDate(models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).WithActive(true).
		WithSort(models.TickerSymbol).WithOrder(models.Asc).WithLimit(2))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err := json.MarshalIndent(iter.Item(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, ticker1, string(b))

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetTickerDetails(t *testing.T) {
	c := polygon.New("API_KEY")

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
	c := polygon.New("API_KEY")

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

func TestGetMarketHolidays(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `[
	{
		"exchange": "NYSE",
		"name": "Thanksgiving",
		"date": "2020-11-26",
		"status": "closed"
	},
	{
		"exchange": "NASDAQ",
		"name": "Thanksgiving",
		"date": "2020-11-26",
		"status": "closed"
	},
	{
		"exchange": "OTC",
		"name": "Thanksgiving",
		"date": "2020-11-26",
		"status": "closed"
	},
	{
		"exchange": "NASDAQ",
		"name": "Thanksgiving",
		"date": "2020-11-27",
		"status": "early-close",
		"open": "2020-11-27T14:30:00.000Z",
		"close": "2020-11-27T18:00:00.000Z"
	},
	{
		"exchange": "NYSE",
		"name": "Thanksgiving",
		"date": "2020-11-27",
		"status": "early-close",
		"open": "2020-11-27T14:30:00.000Z",
		"close": "2020-11-27T18:00:00.000Z"
	}
]`

	registerResponder("https://api.polygon.io/v1/marketstatus/upcoming", expectedResponse)
	res, err := c.Reference.GetMarketHolidays(context.Background())

	assert.Nil(t, err)
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}

func TestGetMarketStatus(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"afterHours": true,
	"currencies": {
		"crypto": "open",
		"fx": "open"
	},
	"earlyHours": false,
	"exchanges": {
		"nasdaq": "extended-hours",
		"nyse": "extended-hours",
		"otc": "closed"
	},
	"market": "extended-hours",
	"serverTime": "2020-11-10T22:37:37.000Z"
}`

	registerResponder("https://api.polygon.io/v1/marketstatus/now", expectedResponse)
	res, err := c.Reference.GetMarketStatus(context.Background())
	assert.Nil(t, err)
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}

func TestListSplits(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	split1 := `{
	"execution_date": "2020-08-31",
	"split_from": 1,
	"split_to": 4,
	"ticker": "AAPL"
}`

	expectedResponse := `{
	"status": "OK",
	"request_id": "2b539ae65c1478dee109b7397bd591b2",
	"results": [
` + indent(true, split1, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/splits?execution_date=2021-07-22&limit=2&order=asc&reverse_split=false&sort=ticker&ticker=AAPL", expectedResponse)
	iter := c.Reference.ListSplits(context.Background(), models.ListSplitsParams{}.
		WithTicker(models.EQ, "AAPL").WithExecutionDate(models.EQ, models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).WithReverseSplit(false).
		WithSort(models.TickerSymbol).WithOrder(models.Asc).WithLimit(2))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err := json.MarshalIndent(iter.Item(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, split1, string(b))

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestListDividends(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	dividend1 := `{
	"cash_amount": 0.59375,
	"declaration_date": "2020-09-09",
	"dividend_type": "CD",
	"ex_dividend_date": "2025-06-12",
	"frequency": 4,
	"pay_date": "2025-06-30",
	"record_date": "2025-06-15",
	"ticker": "CSSEN"
}`

	expectedResponse := `{
	"status": "OK",
	"request_id": "eca6d9a0d8dc1cd1b29d2d3112fe938e",
	"next_url": "https://api.polygon.io/v3/reference/dividends?cursor=YXA9MjUmYXM9JmxpbWl0PTEwJm9yZGVyPWRlc2Mmc29ydD1leF9kaXZpZGVuZF9kYXRlJnRpY2tlcj1DU1NFTg",
	"results": [
` + indent(true, dividend1, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/dividends?dividend_type=CD&ticker=CSSEN", expectedResponse)
	registerResponder("https://api.polygon.io/v3/reference/dividends?cursor=YXA9MjUmYXM9JmxpbWl0PTEwJm9yZGVyPWRlc2Mmc29ydD1leF9kaXZpZGVuZF9kYXRlJnRpY2tlcj1DU1NFTg", "{}")
	iter := c.Reference.ListDividends(context.Background(), models.ListDividendsParams{}.WithTicker(models.EQ, "CSSEN").WithDividendType(models.DividendCD))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err := json.MarshalIndent(iter.Item(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, dividend1, string(b))

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestListConditions(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	condition := `{
	"asset_class": "stocks",
	"data_types": [
		"trade"
	],
	"id": 1,
	"legacy": false,
	"name": "Acquisition",
	"sip_mapping": {
		"UTP": "A"
	},
	"type": "sale_condition",
	"update_rules": {
		"consolidated": {
			"updates_high_low": true,
			"updates_open_close": true,
			"updates_volume": true
		},
		"market_center": {
			"updates_high_low": true,
			"updates_open_close": true,
			"updates_volume": true
		}
	}
}`

	expectedResponse := `{
	"status": "OK",
	"request_id": "4599a4e2ba5e19e2e732f711e97b0d84",
	"count": 1,
	"next_url": "https://api.polygon.io/v3/reference/conditions?cursor=YXA9MiZhcz0mYXNzZXRfY2xhc3M9c3RvY2tzJmRhdGFfdHlwZT10cmFkZSZsaW1pdD0yJnNvcnQ9YXNzZXRfY2xhc3M",
	"results": [
` + indent(true, condition, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/conditions?asset_class=stocks&data_type=trade&limit=1", expectedResponse)
	registerResponder("https://api.polygon.io/v3/reference/conditions?cursor=YXA9MiZhcz0mYXNzZXRfY2xhc3M9c3RvY2tzJmRhdGFfdHlwZT10cmFkZSZsaW1pdD0yJnNvcnQ9YXNzZXRfY2xhc3M", "{}")
	iter := c.Reference.ListConditions(context.Background(), models.ListConditionsParams{}.WithAssetClass(models.AssetStocks).WithDataType(models.DataTrade).WithLimit(1))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err := json.MarshalIndent(iter.Item(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, condition, string(b))

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestListExchanges(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"request_id": "c784b78622b5a68c932af78a68b5907c",
	"count": 1,
	"results": [
		{
			"acronym": "AMEX",
			"asset_class": "stocks",
			"id": 1,
			"locale": "us",
			"mic": "XASE",
			"name": "NYSE American, LLC",
			"operating_mic": "XNYS",
			"participant_id": "A",
			"type": "exchange",
			"url": "https://www.nyse.com/markets/nyse-american"
		}
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/exchanges?asset_class=stocks&locale=us", expectedResponse)
	res, err := c.Reference.GetExchanges(context.Background(), models.GetExchangesParams{}.WithAssetClass(models.AssetStocks).WithLocale(models.US))

	assert.Nil(t, err)
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}
