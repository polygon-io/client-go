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

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
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
	"last_updated_utc": "2021-04-25T00:00:00.000Z"
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

	registerResponder("https://api.polygon.io/v3/reference/tickers?active=true&cik=5&cusip=10&date=2021-07-22&exchange=XNAS&limit=2&market=stocks&order=asc&sort=ticker&type=CS", expectedResponse)
	registerResponder("https://api.polygon.io/v3/reference/tickers?cursor=YWN0aXZlPXRydWUmZGF0ZT0yMDIxLTA0LTI1JmxpbWl0PTEmb3JkZXI9YXNjJnBhZ2VfbWFya2VyPUElN0M5YWRjMjY0ZTgyM2E1ZjBiOGUyNDc5YmZiOGE1YmYwNDVkYzU0YjgwMDcyMWE2YmI1ZjBjMjQwMjU4MjFmNGZiJnNvcnQ9dGlja2Vy", "{}")
	iter := c.ListTickers(context.Background(), models.ListTickersParams{}.
		WithType("CS").WithMarket(models.AssetStocks).
		WithExchange("XNAS").WithCUSIP(10).WithCIK(5).
		WithDate(models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).WithActive(true).
		WithSort(models.TickerSymbol).WithOrder(models.Asc).WithLimit(2))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect models.Ticker
	err := json.Unmarshal([]byte(ticker1), &expect)
	assert.Nil(t, err)
	assert.Equal(t, expect, iter.Item())

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetTickerDetails(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
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
		"last_updated_utc": "2021-04-25T00:00:00.000Z",
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
	res, err := c.GetTickerDetails(context.Background(), models.GetTickerDetailsParams{
		Ticker: "A",
	}.WithDate(models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))))
	assert.Nil(t, err)

	var expect models.GetTickerDetailsResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestListTickerNews(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	news1 := `{
		"id": "1bb0692621da5ab737f4b97ccbef9478ed6820be86474abf8d3b413f8d7d2419",
		"publisher": {
			"name": "The Motley Fool",
			"homepage_url": "https://www.fool.com/",
			"logo_url": "https://s3.polygon.io/public/assets/news/logos/themotleyfool.svg",
			"favicon_url": "https://s3.polygon.io/public/assets/news/favicons/themotleyfool.ico"
		},
		"title": "Should NVIDIA Investors Be Concerned About This Massive Risk? - The Motley Fool",
		"author": "Travis Hoium",
		"published_utc": "2024-07-17T16:08:34Z",
		"article_url": "https://www.fool.com/investing/2024/07/17/should-nvidia-investors-be-concerned-about-this-ma/",
		"tickers": [
			"NVDA",
			"GOOG",
			"GOOGL",
			"META",
			"MSFT"
		],
		"image_url": "https://g.foolcdn.com/editorial/images/783626/ai-chip-black.jpg",
		"description": "NVIDIA's revenue is heavily concentrated among a few customers, which poses a significant risk to its long-term profitability. The company's growth has been driven by a small number of companies ordering AI chips, and this customer concentration could lead to disappointment for investors.",
		"keywords": [
			"NVIDIA",
			"customer concentration",
			"risk",
			"long-term profitability",
			"AI chips"
		],
		"insights": [
			{
				"ticker": "NVDA",
				"sentiment": "negative",
				"sentiment_reasoning": "The article highlights that a few customers make up a huge amount of NVIDIA's revenue, which is a huge risk for the company's long-term profitability. This customer concentration could lead to disappointment for investors."
			},
			{
				"ticker": "GOOG",
				"sentiment": "neutral",
				"sentiment_reasoning": "Suzanne Frey, an executive at Alphabet, is a member of The Motley Fool's board of directors, but the article does not mention Alphabet in the context of NVIDIA's customer concentration risk."
			},
			{
				"ticker": "GOOGL",
				"sentiment": "neutral",
				"sentiment_reasoning": "Suzanne Frey, an executive at Alphabet, is a member of The Motley Fool's board of directors, but the article does not mention Alphabet in the context of NVIDIA's customer concentration risk."
			},
			{
				"ticker": "META",
				"sentiment": "neutral",
				"sentiment_reasoning": "Randi Zuckerberg, a former director of market development and spokeswoman for Facebook (now Meta Platforms), is a member of The Motley Fool's board of directors, but the article does not mention Meta Platforms in the context of NVIDIA's customer concentration risk."
			},
			{
				"ticker": "MSFT",
				"sentiment": "neutral",
				"sentiment_reasoning": "The article does not mention Microsoft in the context of NVIDIA's customer concentration risk."
			}
		]
}`

	expectedResponse := `{
	"status": "OK",
	"count": 1,
	"next_url": "https://api.polygon.io/v2/reference/news?cursor=YXA9MjAyNC0wNy0xN1QxNiUzQTA4JTNBMzRaJmFzPTFiYjA2OTI2MjFkYTVhYjczN2Y0Yjk3Y2NiZWY5NDc4ZWQ2ODIwYmU4NjQ3NGFiZjhkM2I0MTNmOGQ3ZDI0MTkmbGltaXQ9MSZvcmRlcj1kZXNjZW5kaW5nJnRpY2tlcj1NU0ZU",
	"request_id": "2eff2a5bc193b01555b0d4ec91c8fdbf",
	"results": [
` + indent(true, news1, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v2/reference/news?limit=2&order=asc&published_utc.lt=1626912000000&sort=published_utc&ticker.lte=MSFT", expectedResponse)
	registerResponder("https://api.polygon.io/v2/reference/news?cursor=YXA9MjAyNC0wNy0xN1QxNiUzQTA4JTNBMzRaJmFzPTFiYjA2OTI2MjFkYTVhYjczN2Y0Yjk3Y2NiZWY5NDc4ZWQ2ODIwYmU4NjQ3NGFiZjhkM2I0MTNmOGQ3ZDI0MTkmbGltaXQ9MSZvcmRlcj1kZXNjZW5kaW5nJnRpY2tlcj1NU0ZU", "{}")
	iter := c.ListTickerNews(context.Background(), models.ListTickerNewsParams{}.
		WithTicker(models.LTE, "MSFT").WithPublishedUTC(models.LT, models.Millis(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).
		WithSort(models.PublishedUTC).WithOrder(models.Asc).WithLimit(2))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect models.TickerNews
	err := json.Unmarshal([]byte(news1), &expect)
	assert.Nil(t, err)
	assert.Equal(t, expect, iter.Item())

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetTickerRelatedCompanies(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
    "request_id": "0f1dbb2f2781b7d043553bfa400fdfc5",
    "results": [
        {
            "ticker": "MSFT"
        },
        {
            "ticker": "GOOGL"
        },
        {
            "ticker": "AMZN"
        },
        {
            "ticker": "GOOG"
        },
        {
            "ticker": "TSLA"
        },
        {
            "ticker": "NVDA"
        },
        {
            "ticker": "META"
        },
        {
            "ticker": "NFLX"
        },
        {
            "ticker": "DIS"
        },
        {
            "ticker": "BRK.B"
        }
    ],
    "status": "OK",
    "ticker": "AAPL"
}`

	registerResponder("https://api.polygon.io/v1/related-companies/AAPL", expectedResponse)
	params := models.GetTickerRelatedCompaniesParams{
		Ticker: "AAPL",
	}
	res, err := c.GetTickerRelatedCompanies(context.Background(), &params)
	assert.Nil(t, err)

	var expect models.GetTickerRelatedCompaniesResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetTickerTypes(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
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
	res, err := c.GetTickerTypes(context.Background(), models.GetTickerTypesParams{}.WithAssetClass("stocks").WithLocale(models.US))
	assert.Nil(t, err)

	var expect models.GetTickerTypesResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetMarketHolidays(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
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
	res, err := c.GetMarketHolidays(context.Background())
	assert.Nil(t, err)

	var expect models.GetMarketHolidaysResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetMarketStatus(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
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
	"serverTime": "2022-05-17T10:26:06-04:00"
}`

	registerResponder("https://api.polygon.io/v1/marketstatus/now", expectedResponse)
	res, err := c.GetMarketStatus(context.Background())
	assert.Nil(t, err)

	var expect models.GetMarketStatusResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestListSplits(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	split := `{
	"execution_date": "2020-08-31",
	"split_from": 1.0,
	"split_to": 4,
	"ticker": "AAPL"
}`

	expectedResponse := `{
	"status": "OK",
	"request_id": "2b539ae65c1478dee109b7397bd591b2",
	"results": [
` + indent(true, split, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/splits?execution_date=2021-07-22&limit=2&order=asc&reverse_split=false&sort=ticker&ticker=AAPL", expectedResponse)
	iter := c.ListSplits(context.Background(), models.ListSplitsParams{}.
		WithTicker(models.EQ, "AAPL").WithExecutionDate(models.EQ, models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).WithReverseSplit(false).
		WithSort(models.TickerSymbol).WithOrder(models.Asc).WithLimit(2))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect models.Split
	err := json.Unmarshal([]byte(split), &expect)
	assert.Nil(t, err)
	assert.Equal(t, expect, iter.Item())

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestListDividends(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	dividend := `{
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
` + indent(true, dividend, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/dividends?dividend_type=CD&ticker=CSSEN", expectedResponse)
	registerResponder("https://api.polygon.io/v3/reference/dividends?cursor=YXA9MjUmYXM9JmxpbWl0PTEwJm9yZGVyPWRlc2Mmc29ydD1leF9kaXZpZGVuZF9kYXRlJnRpY2tlcj1DU1NFTg", "{}")
	iter := c.ListDividends(context.Background(), models.ListDividendsParams{}.WithTicker(models.EQ, "CSSEN").WithDividendType(models.DividendCD))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect models.Dividend
	err := json.Unmarshal([]byte(dividend), &expect)
	assert.Nil(t, err)
	assert.Equal(t, expect, iter.Item())

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestListConditions(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
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
	iter := c.ListConditions(context.Background(), models.ListConditionsParams{}.WithAssetClass(models.AssetStocks).WithDataType(models.DataTrade).WithLimit(1))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect models.Condition
	err := json.Unmarshal([]byte(condition), &expect)
	assert.Nil(t, err)
	assert.Equal(t, expect, iter.Item())

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetExchanges(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
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
	res, err := c.GetExchanges(context.Background(), models.GetExchangesParams{}.WithAssetClass(models.AssetStocks).WithLocale(models.US))
	assert.Nil(t, err)

	var expect models.GetExchangesResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetOptionsContract(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"results": {
		"cfi": "OCASPS",
		"contract_type": "call",
		"exercise_style": "american",
		"expiration_date": "2024-01-19",
		"primary_exchange": "BATO",
		"shares_per_contract": 100,
		"strike_price": 2.5,
		"ticker": "O:EVRI240119C00002500",
		"underlying_ticker": "EVRI"
	},
	"status": "OK",
	"request_id": "52fccf652441fc4d4fd35e2d0d2dd1f2"
}`

	registerResponder("https://api.polygon.io/v3/reference/options/contracts/O:EVRI240119C00002500", expectedResponse)
	res, err := c.GetOptionsContract(context.Background(), models.GetOptionsContractParams{
		Ticker: "O:EVRI240119C00002500",
	}.WithAsOf(models.Date(time.Date(2022, 5, 16, 0, 0, 0, 0, time.Local))))
	assert.Nil(t, err)

	var expect models.GetOptionsContractResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestListOptionsContracts(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	contract1 := `{
	"cfi": "OCASPS",
	"contract_type": "call",
	"exercise_style": "american",
	"expiration_date": "2022-05-20",
	"primary_exchange": "BATO",
	"shares_per_contract": 100,
	"strike_price": 65,
	"ticker": "O:A220520C00065000",
	"underlying_ticker": "A"
}`

	contract2 := `{
	"cfi": "OCASPS",
	"contract_type": "call",
	"exercise_style": "american",
	"expiration_date": "2022-05-20",
	"primary_exchange": "BATO",
	"shares_per_contract": 100,
	"strike_price": 70,
	"ticker": "O:A220520C00070000",
	"underlying_ticker": "A"
}`

	expectedResponse := `{
	"status": "OK",
	"request_id": "975d5e1aacc6c147c94934b016b8d1a7",
	"next_url": "https://api.polygon.io/v3/reference/options/contracts?cursor=YXA9JTdCJTIySUQlMjIlM0ElMjIyNzI2MTY3OTUxMjgzOTI1NTI5JTIyJTJDJTIyU3RhcnREYXRlVXRjJTIyJTNBJTdCJTIyVGltZSUyMiUzQSUyMjIwMjEtMDktMjFUMDAlM0EwMCUzQTAwWiUyMiUyQyUyMlZhbGlkJTIyJTNBdHJ1ZSU3RCUyQyUyMkVuZERhdGVVdGMlMjIlM0ElN0IlMjJUaW1lJTIyJTNBJTIyMDAwMS0wMS0wMVQwMCUzQTAwJTNBMDBaJTIyJTJDJTIyVmFsaWQlMjIlM0FmYWxzZSU3RCUyQyUyMnVuZGVybHlpbmdfdGlja2VyJTIyJTNBJTIyQSUyMiUyQyUyMnRpY2tlciUyMiUzQSUyMk8lM0FBMjIwNTIwQzAwMTEwMDAwJTIyJTJDJTIyZXhwaXJhdGlvbl9kYXRlJTIyJTNBJTIyMjAyMi0wNS0yMFQwMCUzQTAwJTNBMDBaJTIyJTJDJTIyc3RyaWtlX3ByaWNlJTIyJTNBMTEwJTJDJTIyY2ZpJTIyJTNBJTIyT0NBU1BTJTIyJTJDJTIyY29udHJhY3RfdHlwZSUyMiUzQSUyMmNhbGwlMjIlMkMlMjJleGVyY2lzZV9zdHlsZSUyMiUzQSUyMmFtZXJpY2FuJTIyJTJDJTIycHJpbWFyeV9leGNoYW5nZSUyMiUzQSU3QiUyMlN0cmluZyUyMiUzQSUyMkJBVE8lMjIlMkMlMjJWYWxpZCUyMiUzQXRydWUlN0QlMkMlMjJzaGFyZXNfcGVyX2NvbnRyYWN0JTIyJTNBMTAwJTJDJTIyYWRkaXRpb25hbF91bmRlcmx5aW5ncyUyMiUzQSUyMm51bGwlMjIlN0QmYXM9JmNvbnRyYWN0X3R5cGU9Y2FsbCZsaW1pdD0xMCZzb3J0PXRpY2tlcg",
	"results": [
` + indent(true, contract1, "\t\t") + `,
` + indent(true, contract2, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/options/contracts?contract_type=call", expectedResponse)
	registerResponder("https://api.polygon.io/v3/reference/options/contracts?cursor=YXA9JTdCJTIySUQlMjIlM0ElMjIyNzI2MTY3OTUxMjgzOTI1NTI5JTIyJTJDJTIyU3RhcnREYXRlVXRjJTIyJTNBJTdCJTIyVGltZSUyMiUzQSUyMjIwMjEtMDktMjFUMDAlM0EwMCUzQTAwWiUyMiUyQyUyMlZhbGlkJTIyJTNBdHJ1ZSU3RCUyQyUyMkVuZERhdGVVdGMlMjIlM0ElN0IlMjJUaW1lJTIyJTNBJTIyMDAwMS0wMS0wMVQwMCUzQTAwJTNBMDBaJTIyJTJDJTIyVmFsaWQlMjIlM0FmYWxzZSU3RCUyQyUyMnVuZGVybHlpbmdfdGlja2VyJTIyJTNBJTIyQSUyMiUyQyUyMnRpY2tlciUyMiUzQSUyMk8lM0FBMjIwNTIwQzAwMTEwMDAwJTIyJTJDJTIyZXhwaXJhdGlvbl9kYXRlJTIyJTNBJTIyMjAyMi0wNS0yMFQwMCUzQTAwJTNBMDBaJTIyJTJDJTIyc3RyaWtlX3ByaWNlJTIyJTNBMTEwJTJDJTIyY2ZpJTIyJTNBJTIyT0NBU1BTJTIyJTJDJTIyY29udHJhY3RfdHlwZSUyMiUzQSUyMmNhbGwlMjIlMkMlMjJleGVyY2lzZV9zdHlsZSUyMiUzQSUyMmFtZXJpY2FuJTIyJTJDJTIycHJpbWFyeV9leGNoYW5nZSUyMiUzQSU3QiUyMlN0cmluZyUyMiUzQSUyMkJBVE8lMjIlMkMlMjJWYWxpZCUyMiUzQXRydWUlN0QlMkMlMjJzaGFyZXNfcGVyX2NvbnRyYWN0JTIyJTNBMTAwJTJDJTIyYWRkaXRpb25hbF91bmRlcmx5aW5ncyUyMiUzQSUyMm51bGwlMjIlN0QmYXM9JmNvbnRyYWN0X3R5cGU9Y2FsbCZsaW1pdD0xMCZzb3J0PXRpY2tlcg", "{}")
	iter := c.ListOptionsContracts(context.Background(), models.ListOptionsContractsParams{}.WithContractType("call"))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect1 models.OptionsContract
	err := json.Unmarshal([]byte(contract1), &expect1)
	assert.Nil(t, err)
	assert.Equal(t, expect1, iter.Item())

	// second item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect2 models.OptionsContract
	err = json.Unmarshal([]byte(contract2), &expect2)
	assert.Nil(t, err)
	assert.Equal(t, expect2, iter.Item())

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetShortInterest(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
		"results": [{
			"currency_code": "USD",
			"date": "2023-12-31",
			"isin": "US0378331005",
			"name": "Apple Inc.",
			"security_description": "Common Stock",
			"short_volume": 2006566,
			"short_volume_exempt": 3000,
			"ticker": "AAPL",
			"us_code": "378331005"
		}],
		"status": "OK",
		"request_id": "some-request-id",
		"next_url": null
	}`

	// Mock the API response
	registerResponder("https://api.polygon.io/v1/reference/short-interest/ticker/AAPL", expectedResponse)

	params := models.GetShortInterestParams{
		IdentifierType: "ticker",
		Identifier:     "AAPL",
	}
	res, err := c.GetShortInterest(context.Background(), &params)
	assert.Nil(t, err)

	var expect models.GetShortInterestResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestListIPOs(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	ipo1 := `{
        "currency_code": "USD",
        "final_issue_price": 17,
        "highest_offer_price": 17,
        "ipo_status": "HISTORY",
        "isin": "US75383L1026",
        "issue_end_date": "2024-06-06",
        "issue_start_date": "2024-06-01",
        "issuer_name": "Rapport Therapeutics Inc.",
        "last_updated": "2024-06-27",
        "listing_date": "2024-06-07",
        "listing_price": null,
        "lot_size": 100,
        "lowest_offer_price": 17,
        "max_shares_offered": 8000000,
        "min_shares_offered": 1000000,
        "primary_exchange": "XNAS",
        "security_description": "Ordinary Shares",
        "security_type": "CS",
        "shares_outstanding": 35376457,
        "ticker": "RAPP",
        "total_offer_size": 136000000,
        "us_code": "75383L102"
    }`

	// Construct the expected API response with the IPO listing
	expectedResponse := `{
        "status": "OK",
        "count": 1,
        "next_url": "https://api.polygon.io/v1/reference/ipos?cursor=nextCursorValue",
        "request_id": "6a7e466379af0a71039d60cc78e72282",
        "results": [
    ` + indent(true, ipo1, "\t\t") + `
        ]
    }`

	registerResponder("https://api.polygon.io/v1/reference/ipos?limit=10&order=asc&sort=listing_date", expectedResponse)
	registerResponder("https://api.polygon.io/v1/reference/ipos?cursor=nextCursorValue", "{}")

	iter := c.ListIPOs(context.Background(), models.ListIPOsParams{}.
		WithLimit(10).
		WithOrder(models.Asc).
		WithSort(models.IPOsSortListingDate))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect models.IPOListing
	err := json.Unmarshal([]byte(ipo1), &expect)
	assert.Nil(t, err)
	assert.Equal(t, expect, iter.Item())

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}
