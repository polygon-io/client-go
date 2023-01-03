package polygon_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

var snapshot1 = `{
	"day": {
		"c": 20.506,
		"h": 20.64,
		"l": 20.506,
		"o": 20.64,
		"v": 37216,
		"vw": 20.616
	},
	"lastQuote": {
		"P": 20.6,
		"p": 20.5,
		"S": 22,
		"s": 13,
		"t": 1605192959994246100
	},
	"lastTrade": {
		"c": [
			14,
			41
		],
		"i": "71675577320245",
		"p": 20.506,
		"s": 2416,
		"t": 1605192894630916600,
		"x": 4
	},
	"min": {
		"av": 37216,
		"c": 20.506,
		"h": 20.506,
		"l": 20.506,
		"o": 20.506,
		"v": 5000,
		"vw": 20.5105
	},
	"prevDay": {
		"c": 20.63,
		"h": 21,
		"l": 20.5,
		"o": 20.79,
		"v": 292738,
		"vw": 20.6939
	},
	"ticker": "BCAT",
	"todaysChange": -0.124,
	"todaysChangePerc": -0.601,
	"updated": 1605192894630916600
}`

var snapshot2 = `{
	"day": {
		"c": 313.225,
		"h": 314.35,
		"l": 309.71,
		"o": 310.09,
		"v": 6322693,
		"vw": 312.6791
	},
	"lastQuote": {
		"P": 313.13,
		"p": 313.11,
		"S": 4,
		"s": 2,
		"t": 1649083047683654000
	},
	"lastTrade": {
		"i": "23432",
		"p": 313.1296,
		"s": 100,
		"t": 1649083047682204000,
		"x": 4
	},
	"min": {
		"av": 6321712,
		"c": 313.1826,
		"h": 313.19,
		"l": 312.66,
		"o": 312.78,
		"v": 54315,
		"vw": 312.9441
	},
	"prevDay": {
		"c": 309.42,
		"h": 310.13,
		"l": 305.54,
		"o": 309.37,
		"v": 27101029,
		"vw": 308.0485
	},
	"ticker": "MSFT",
	"todaysChange": 3.71,
	"todaysChangePerc": 1.199,
	"updated": 1649083047682204000
}`

func TestListSnapshotAllTickers(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"count": 2,
	"tickers": [
` + indent(true, snapshot1, "\t\t") + `,
` + indent(true, snapshot2, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v2/snapshot/locale/us/markets/stocks/tickers?tickers=AAPL%2CMSFT", expectedResponse)
	res, err := c.GetAllTickersSnapshot(context.Background(), models.GetAllTickersSnapshotParams{
		Locale:     "us",
		MarketType: "stocks",
	}.WithTickers("AAPL,MSFT"))
	assert.Nil(t, err)

	var expect models.GetAllTickersSnapshotResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetTickerSnapshot(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"count": 2,
	"ticker": ` + indent(false, snapshot1, "\t") + `
}`

	registerResponder("https://api.polygon.io/v2/snapshot/locale/us/markets/stocks/tickers/AAPL", expectedResponse)
	res, err := c.GetTickerSnapshot(context.Background(), &models.GetTickerSnapshotParams{
		Ticker:     "AAPL",
		Locale:     "us",
		MarketType: "stocks",
	})
	assert.Nil(t, err)

	var expect models.GetTickerSnapshotResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetGainersLosersSnapshot(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"count": 2,
	"tickers": [
` + indent(true, snapshot1, "\t\t") + `,
` + indent(true, snapshot2, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v2/snapshot/locale/us/markets/stocks/gainers", expectedResponse)
	res, err := c.GetGainersLosersSnapshot(context.Background(), &models.GetGainersLosersSnapshotParams{
		Locale:     "us",
		MarketType: "stocks",
		Direction:  "gainers",
	})
	assert.Nil(t, err)

	var expect models.GetGainersLosersSnapshotResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetOptionContractSnapshot(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"request_id": "d9ff18dac69f55c218f69e4753706acd",
	"results": {
		"break_even_price": 171.075,
		"day": {
			"change": -1.05,
			"change_percent": -4.67,
			"close": 21.4,
			"high": 22.49,
			"last_updated": 1636520400000000000,
			"low": 21.35,
			"open": 22.49,
			"previous_close": 22.45,
			"volume": 37,
			"vwap": 21.6741
		},
		"details": {
			"contract_type": "call",
			"exercise_style": "american",
			"expiration_date": "2023-06-16",
			"shares_per_contract": 100,
			"strike_price": 150,
			"ticker": "O:AAPL230616C00150000"
		},
		"greeks": {
			"delta": 0.5520187372272933,
			"gamma": 0.00706756515659829,
			"theta": -0.018532772783847958,
			"vega": 0.7274811132998142
		},
		"implied_volatility": 0.3048997097864957,
		"last_quote": {
			"ask": 21.25,
			"ask_size": 110,
			"bid": 20.9,
			"bid_size": 172,
			"last_updated": 1636573458756383500,
			"midpoint": 21.075,
			"timeframe": "REAL-TIME"
		},
		"open_interest": 8921,
		"underlying_asset": {
			"change_to_break_even": 23.123999999999995,
			"last_updated": 1636573459862384600,
			"price": 147.951,
			"ticker": "AAPL",
			"timeframe": "REAL-TIME"
		}
	}
}`

	registerResponder("https://api.polygon.io/v3/snapshot/options/AAPL/O:AAPL230616C00150000", expectedResponse)
	res, err := c.GetOptionContractSnapshot(context.Background(), &models.GetOptionContractSnapshotParams{
		UnderlyingAsset: "AAPL",
		OptionContract:  "O:AAPL230616C00150000",
	})
	assert.Nil(t, err)

	var expect models.GetOptionContractSnapshotResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetOptionsChainSnapshot(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	  "results": [
		{
		  "break_even_price": 162.375,
		  "day": {
			"change": 0,
			"change_percent": 0,
			"close": 79.35,
			"high": 79.35,
			"last_updated": 1672434000000,
			"low": 79.3,
			"open": 79.3,
			"previous_close": 79.35,
			"volume": 22,
			"vwap": 79.325
		  },
		  "details": {
			"contract_type": "call",
			"exercise_style": "american",
			"expiration_date": "2023-01-06",
			"shares_per_contract": 100,
			"strike_price": 50,
			"ticker": "O:AAPL230106C00050000"
		  },
		  "greeks": {},
		  "last_quote": {
			"ask": 75.05,
			"ask_size": 48,
			"bid": 74.85,
			"bid_size": 43,
			"last_updated": 1672775256862312000,
			"midpoint": 112.375,
			"timeframe": "DELAYED"
		  },
		  "open_interest": 5,
		  "underlying_asset": {
			"change_to_break_even": 37.435,
			"last_updated": 1672775257417223400,
			"price": 124.94,
			"ticker": "AAPL",
			"timeframe": "DELAYED"
		  }
		}
	  ],
	  "status": "OK",
	  "request_id": "0d350849-a2a8-43c5-8445-9c6f55d371e6",
	  "next_url": "http://localhost:8080/v3/snapshot/options/AAPL?cursor=YXA9MSZhcz0mbGltaXQ9MSZzb3J0PXRpY2tlcg"
	}`

	registerResponder("https://api.polygon.io/v3/snapshot/options/AAPL", expectedResponse)

	chainParams := models.NewGetOptionsChainParams("AAPL").WithLimit(1)

	res, err := c.GetOptionsChainSnapshot(context.Background(), chainParams)
	assert.Nil(t, err)

	var expect models.GetOptionsChainSnapshotResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)

	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestGetCryptoFullBookSnapshot(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"data": {
		"askCount": 593.1412981600005,
		"asks": [
			{
				"p": 11454,
				"x": {
					"2": 1
				}
			},
			{
				"p": 11455,
				"x": {
					"2": 1
				}
			}
		],
		"bidCount": 694.951789670001,
		"bids": [
			{
				"p": 16303.17,
				"x": {
					"1": 2
				}
			},
			{
				"p": 16302.94,
				"x": {
					"1": 0.02859424,
					"6": 0.023455
				}
			}
		],
		"spread": -4849.17,
		"ticker": "X:BTCUSD",
		"updated": 1605295074162
	}
}`

	registerResponder("https://api.polygon.io/v2/snapshot/locale/global/markets/crypto/tickers/X:BTCUSD/book", expectedResponse)
	res, err := c.GetCryptoFullBookSnapshot(context.Background(), &models.GetCryptoFullBookSnapshotParams{
		Ticker: "X:BTCUSD",
	})
	assert.Nil(t, err)

	var expect models.GetCryptoFullBookSnapshotResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}
