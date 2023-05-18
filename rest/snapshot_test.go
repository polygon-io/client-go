package polygon_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		"av": 37216.0,
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
		"last_trade": {
			"sip_timestamp": 1676573362154648300,
			"conditions": [
				209
			],
			"price": 110.9,
			"size": 10,
			"exchange": 308,
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

func TestListOptionsChainSnapshot(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	chain1 := `{
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
		"last_trade": {
			"sip_timestamp": 1676573362154648300,
			"conditions": [
				209
			],
			"price": 110.9,
			"size": 10,
			"exchange": 308,
			"timeframe": "REAL-TIME"
		},
		"open_interest": 5,
		"underlying_asset": {
			"change_to_break_even": 37.435,
			"last_updated": 1672775257417223400,
			"price": 124.94,
			"ticker": "AAPL",
			"timeframe": "DELAYED"
		}
	}`
	chain2 := `{
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
		"last_trade": {
			"sip_timestamp": 1676573362154648300,
			"conditions": [
				209
			],
			"price": 110.9,
			"size": 10,
			"exchange": 308,
			"timeframe": "REAL-TIME"
		},
		"open_interest": 5,
		"underlying_asset": {
			"change_to_break_even": 37.435,
			"last_updated": 1672775257417223400,
			"price": 124.94,
			"ticker": "AAPL",
			"timeframe": "DELAYED"
		}
	}`
	chain3 := `{
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
		"last_trade": {
			"sip_timestamp": 1676573362154648300,
			"conditions": [
				209
			],
			"price": 110.9,
			"size": 10,
			"exchange": 308,
			"timeframe": "REAL-TIME"
		},
		"open_interest": 5,
		"underlying_asset": {
			"change_to_break_even": 37.435,
			"last_updated": 1672775257417223400,
			"price": 124.94,
			"ticker": "AAPL",
			"timeframe": "DELAYED"
		}
	}`

	expectedResponse := `{
	  "results": [
		` + indent(true, chain1, "\t\t") + `,
		` + indent(true, chain2, "\t\t") + `,
		` + indent(true, chain3, "\t\t") + `
	  ],
	  "status": "OK",
	  "request_id": "0d350849-a2a8-43c5-8445-9c6f55d371e6",
	  "next_url": "https://api.polygon.io/v3/snapshot/options/AAPL?cursor=YXA9MSZhcz0mbGltaXQ9MSZzb3J0PXRpY2tlcg"
	}`

	registerResponder("https://api.polygon.io/v3/snapshot/options/AAPL", expectedResponse)
	registerResponder("https://api.polygon.io/v3/snapshot/options/AAPL?cursor=YXA9MSZhcz0mbGltaXQ9MSZzb3J0PXRpY2tlcg", "{}")

	iter := c.ListOptionsChainSnapshot(context.Background(), &models.ListOptionsChainParams{UnderlyingAsset: "AAPL"})

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect1 models.OptionContractSnapshot
	err := json.Unmarshal([]byte(chain1), &expect1)
	assert.Nil(t, err)
	assert.Equal(t, expect1, iter.Item())

	// second item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect2 models.OptionContractSnapshot
	err = json.Unmarshal([]byte(chain2), &expect2)
	assert.Nil(t, err)
	assert.Equal(t, expect2, iter.Item())

	// third item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect3 models.OptionContractSnapshot
	err = json.Unmarshal([]byte(chain3), &expect3)
	assert.Nil(t, err)
	assert.Equal(t, expect3, iter.Item())

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
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

func TestGetIndicesSnapshot(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()
	expectedIndicesSnapshotResponse := `{
  "results": [
    {
      "value": 1326.17,
      "name": "Dow Jones Americas Health Care Index",
      "ticker": "I:A1HCR",
      "market_status": "open",
      "type": "indices",
      "session": {
        "change": 47.07,
        "change_percent": 3.68,
        "close": 1282.67,
        "high": 1288.89,
        "low": 1282.25,
        "open": 1283.33,
        "previous_close": 1279.1000000000001
      }
    },
    {
      "value": 3918.32,
      "name": "Standard & Poor's 500",
      "ticker": "I:SPX",
      "market_status": "open",
      "type": "indices",
      "session": {
        "change": 5.56,
        "change_percent": 0.142,
        "close": 3926.36,
        "high": 3927.38,
        "low": 3878.1,
        "open": 3914.13,
        "previous_close": 3912.76
      }
    }
  ],
  "status": "OK",
  "request_id": "5ad18f153c5aa4a543cc10aeb9245622"
}

`

	expectedGetIndicesSnapshotUrl := "https://api.polygon.io/v3/snapshot/indices?ticker.any_of=I%3AA1HCR%2CI%3ASPX"
	registerResponder(expectedGetIndicesSnapshotUrl, expectedIndicesSnapshotResponse)
	tickerAnyOf := []string{"I:A1HCR", "I:SPX"}

	res, err := c.GetIndicesSnapshot(context.Background(), models.GetIndicesSnapshotParams{}.WithTickerAnyOf(tickerAnyOf...))
	assert.Nil(t, err)

	var expect models.GetIndicesSnapshotResponse
	err = json.Unmarshal([]byte(expectedIndicesSnapshotResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}

func TestListAssetSnapshots(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	tt := []struct {
		name           string
		haveParams     *models.ListAssetSnapshotsParams
		haveRequestURL string
		wantResponse   string
		testData       []string
		wantErr        bool
	}{
		{
			name:           "Stock Snapshot",
			haveParams:     models.ListAssetSnapshotsParams{}.WithTickerAnyOf("AAPL,META,F"),
			haveRequestURL: "https://api.polygon.io/v3/snapshot?ticker.any_of=AAPL%2CMETA%2CF",
			wantResponse: `{
				"results": [
					` + indent(true, stockSnapshotsTestData[0], "\t\t") + `,
					` + indent(true, stockSnapshotsTestData[1], "\t\t") + `,
					` + indent(true, stockSnapshotsTestData[2], "\t\t") + `
					],
					"status": "OK",
					"request_id": "0d350849-a2a8-43c5-8445-9c6f55d371e6",
					"next_url": "https://api.polygon.io/v3/snapshot/cursor=YXA9MSZhcz0mbGltaXQ9MSZzb3J0PXRpY2tlcg"
				}`,
			testData: stockSnapshotsTestData,
			wantErr:  false,
		},
		{
			name:           "Options Snapshot",
			haveParams:     models.ListAssetSnapshotsParams{}.WithTickerAnyOf("O:AAPL230512C00050000,O:META230512C00020000,O:F230512C00005000"),
			haveRequestURL: "https://api.polygon.io/v3/snapshot?ticker.any_of=O%3AAAPL230512C00050000%2CO%3AMETA230512C00020000%2CO%3AF230512C00005000",
			wantResponse: `{
				"results": [
					` + indent(true, optionsSnapshotsTestData[0], "\t\t") + `,
					` + indent(true, optionsSnapshotsTestData[1], "\t\t") + `,
					` + indent(true, optionsSnapshotsTestData[2], "\t\t") + `
					],
					"status": "OK",
					"request_id": "0d350849-a2a8-43c5-8445-9c6f55d371e6",
					"next_url": "https://api.polygon.io/v3/snapshot/cursor=YXA9MSZhcz0mbGltaXQ9MSZzb3J0PXRpY2tlcg"
				}`,
			testData: optionsSnapshotsTestData,
		},
		{
			name:           "Crypto Snapshot",
			haveParams:     models.ListAssetSnapshotsParams{}.WithTickerAnyOf("X:BTCUSD,X:ETHUSD,X:FLOWUSD"),
			haveRequestURL: "https://api.polygon.io/v3/snapshot?ticker.any_of=X%3ABTCUSD%2CX%3AETHUSD%2CX%3AFLOWUSD",
			wantResponse: `{
				"results": [
					` + indent(true, cryptoSnapshotsTestData[0], "\t\t") + `,
					` + indent(true, cryptoSnapshotsTestData[1], "\t\t") + `,
					` + indent(true, cryptoSnapshotsTestData[2], "\t\t") + `
					],
					"status": "OK",
					"request_id": "0d350849-a2a8-43c5-8445-9c6f55d371e6",
					"next_url": "https://api.polygon.io/v3/snapshot/cursor=YXA9MSZhcz0mbGltaXQ9MSZzb3J0PXRpY2tlcg"
				}`,
			testData: cryptoSnapshotsTestData,
		},
		{
			name:           "Forex Snapshot",
			haveParams:     models.ListAssetSnapshotsParams{}.WithTickerAnyOf("C:USDCAD,C:USDEUR,C:USDAUD"),
			haveRequestURL: "https://api.polygon.io/v3/snapshot?ticker.any_of=C%3AUSDCAD%2CC%3AUSDEUR%2CC%3AUSDAUD",
			wantResponse: `{
				"results": [
					` + indent(true, forexSnapshotTestData[0], "\t\t") + `,
					` + indent(true, forexSnapshotTestData[1], "\t\t") + `,
					` + indent(true, forexSnapshotTestData[2], "\t\t") + `
					],
					"status": "OK",
					"request_id": "0d350849-a2a8-43c5-8445-9c6f55d371e6",
					"next_url": "https://api.polygon.io/v3/snapshot/cursor=YXA9MSZhcz0mbGltaXQ9MSZzb3J0PXRpY2tlcg"
				}`,
			testData: forexSnapshotTestData,
		},
		{
			name:           "Indices Snapshot",
			haveParams:     models.ListAssetSnapshotsParams{}.WithTickerAnyOf("I:A1HCR,I:SPX"),
			haveRequestURL: "https://api.polygon.io/v3/snapshot?ticker.any_of=C%3AUSDCAD%2CC%3AUSDEUR%2CC%3AUSDAUD",
			wantResponse: `{
				"results": [
					` + indent(true, indicesSnapshotTestData[0], "\t\t") + `,
					` + indent(true, indicesSnapshotTestData[1], "\t\t") + `,
					` + indent(true, indicesSnapshotTestData[2], "\t\t") + `
					],
					"status": "OK",
					"request_id": "0d350849-a2a8-43c5-8445-9c6f55d371e6",
					"next_url": "https://api.polygon.io/v3/snapshot/cursor=YXA9MSZhcz0mbGltaXQ9MSZzb3J0PXRpY2tlcg"
				}`,
			testData: indicesSnapshotTestData,
		},
		{
			name:           "Partial success (200/OK with an error message in the body)",
			haveParams:     models.ListAssetSnapshotsParams{}.WithTickerAnyOf("AAPL,APx"),
			haveRequestURL: "https://api.polygon.io/v3/snapshot?ticker.any_of=AAPL%2CAPx",
			wantResponse: `{
				"results": [
					` + indent(true, partialSuccessWithStocksTestData[0], "\t\t") + `,
					` + indent(true, partialSuccessWithStocksTestData[1], "\t\t") + `
					],
					"status": "OK",
					"request_id": "0d350849-a2a8-43c5-8445-9c6f55d371e6",
					"next_url": "https://api.polygon.io/v3/snapshot/cursor=YXA9MSZhcz0mbGltaXQ9MSZzb3J0PXRpY2tlcg"
				}`,
			testData: partialSuccessWithStocksTestData,
			wantErr:  true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			registerResponder(tc.haveRequestURL, tc.wantResponse)
			registerResponder("https://api.polygon.io/v3/snapshot/cursor=YXA9MSZhcz0mbGltaXQ9MSZzb3J0PXRpY2tlcg", "{}")

			iter := c.ListAssetSnapshots(
				context.Background(),
				tc.haveParams,
			)

			require.NoError(t, iter.Err())
			require.NotNil(t, iter.Item())

			var iterCount int
			for iter.Next() {
				var gotSnapshot models.SnapshotResponseModel
				err := json.Unmarshal([]byte(tc.testData[iterCount]), &gotSnapshot)
				require.Nil(t, err)

				require.Nil(t, iter.Err())
				assert.Equal(t, gotSnapshot, iter.Item())
				iterCount++
			}

			assert.Equal(t, len(tc.testData), iterCount, fmt.Sprintf("expected %d results", len(tc.testData)))
			assert.False(t, iter.Next())
			assert.Nil(t, iter.Err())
		})
	}
}

var stockSnapshotsTestData = []string{
	`{
		"market_status": "late_trading",
		"name": "Apple Inc.",
		"session": {
			"change": -0.07,
			"change_percent": -0.0403,
			"close": 173.5,
			"early_trading_change": 0,
			"early_trading_change_percent": 0,
			"high": 173.85,
			"low": 172.11,
			"open": 172.48,
			"previous_close": 173.57,
			"price": 173.5,
			"volume": 50823329
		},
		"last_quote": {
			"ask": 173.34,
			"ask_size": 3,
			"bid": 173.32,
			"bid_size": 4,
			"last_updated": 1683577209434314800,
			"timeframe": "REAL-TIME"
		},
		"last_trade": {
			"conditions": [
				12,
				22
			],
			"exchange": 4,
			"id": "247862",
			"last_updated": 1683577205678289200,
			"price": 173.5,
			"size": 31535,
			"timeframe": "REAL-TIME"
		},
		"ticker": "AAPL",
		"type": "stocks"
	}`,
	`{
		"market_status": "late_trading",
		"name": "Meta Platforms, Inc. Class A Common Stock",
		"session": {
			"change": -0.04,
			"change_percent": -0.0172,
			"close": 233.27,
			"early_trading_change": 0,
			"early_trading_change_percent": 0,
			"high": 235.62,
			"low": 230.27,
			"open": 231.415,
			"previous_close": 232.78,
			"price": 232.74,
			"volume": 14940329
		},
		"last_quote": {
			"ask": 232.83,
			"ask_size": 1,
			"bid": 232.73,
			"bid_size": 1,
			"last_updated": 1683577187244746200,
			"timeframe": "REAL-TIME"
		},
		"last_trade": {
			"conditions": [
				12,
				37
			],
			"exchange": 4,
			"id": "57128",
			"last_updated": 1683577202547284000,
			"price": 232.74,
			"size": 50,
			"timeframe": "REAL-TIME"
		},
		"ticker": "META",
		"type": "stocks"
	}`,
	`{
		"market_status": "late_trading",
		"name": "Ford Motor Company",
		"session": {
			"change": 0.005,
			"change_percent": 0.0417,
			"close": 12.02,
			"early_trading_change": 0,
			"early_trading_change_percent": 0,
			"high": 12.055,
			"low": 11.85,
			"open": 12.02,
			"previous_close": 11.99,
			"price": 11.995,
			"volume": 49539926
		},
		"last_quote": {
			"ask": 12,
			"ask_size": 23,
			"bid": 11.99,
			"bid_size": 28,
			"last_updated": 1683577084319878700,
			"timeframe": "REAL-TIME"
		},
		"last_trade": {
			"conditions": [
				12,
				37
			],
			"exchange": 4,
			"id": "71697320268354",
			"last_updated": 1683577186411804000,
			"price": 11.995,
			"size": 1,
			"timeframe": "REAL-TIME"
		},
		"ticker": "F",
		"type": "stocks"
	}`,
}

var optionsSnapshotsTestData = []string{
	`{
		"name": "AAPL $50.00 call",
		"market_status": "open",
		"ticker": "O:AAPL230512C00050000",
		"type": "options",
		"last_quote": {
		  "ask": 123.1,
		  "ask_size": 90,
		  "bid": 122.95,
		  "bid_size": 90,
		  "last_updated": 1683731850932649728,
		  "midpoint": 123.025,
		  "timeframe": "REAL-TIME"
		},
		"last_trade": {},
		"session": {},
		"break_even_price": 173.025,
		"details": {
		  "contract_type": "call",
		  "exercise_style": "american",
		  "expiration_date": "2023-05-12",
		  "shares_per_contract": 100,
		  "strike_price": 50
		},
		"greeks": {},
		"underlying_asset": {
		  "change_to_break_even": -0.11,
		  "last_updated": 1683732072879546553,
		  "price": 173.135,
		  "ticker": "AAPL",
		  "timeframe": "REAL-TIME"
		},
		"error": "",
		"message": ""
	}`,
	`{
		"name": "META $20.00 call",
		"market_status": "open",
		"ticker": "O:META230512C00020000",
		"type": "options",
		"last_quote": {},
		"last_trade": {
		  "sip_timestamp": 1682970890371000000,
		  "conditions": [
			209
		  ],
		  "price": 223.75,
		  "size": 1,
		  "exchange": 302,
		  "timeframe": "REAL-TIME"
		},
		"session": {},
		"details": {
		  "contract_type": "call",
		  "exercise_style": "american",
		  "expiration_date": "2023-05-12",
		  "shares_per_contract": 100,
		  "strike_price": 20
		},
		"greeks": {},
		"underlying_asset": {
		  "last_updated": 1683731579449632715,
		  "price": 232.37,
		  "ticker": "META",
		  "timeframe": "REAL-TIME"
		},
		"error": "",
		"message": ""
	}`,
	`{
		"name": "F $5.00 call",
		"market_status": "open",
		"ticker": "O:F230512C00005000",
		"type": "options",
		"last_quote": {},
		"last_trade": {
		  "sip_timestamp": 1683316735432000000,
		  "conditions": [
			232
		  ],
		  "price": 6.97,
		  "size": 1,
		  "exchange": 312,
		  "timeframe": "REAL-TIME"
		},
		"session": {},
		"details": {
		  "contract_type": "call",
		  "exercise_style": "american",
		  "expiration_date": "2023-05-12",
		  "shares_per_contract": 100,
		  "strike_price": 5
		},
		"greeks": {},
		"underlying_asset": {
		  "last_updated": 1683732072773028096,
		  "price": 11.93,
		  "ticker": "F",
		  "timeframe": "REAL-TIME"
		}
	}`,
}

var cryptoSnapshotsTestData = []string{
	`{
		"market_status": "open",
		"name": "Bitcoin - United States Dollar",
		"ticker": "X:BTCUSD",
		"type": "crypto",
		"session": {
		  "change": -181,
		  "change_percent": -0.661,
		  "close": 27236.1,
		  "high": 27506,
		  "low": 27010,
		  "open": 27402.3,
		  "volume": 10012.03414028,
		  "previous_close": 27400.74,
		  "price": 27220
		},
		"last_trade": {
		  "participant_timestamp": 1684422449502000000,
		  "timeframe": "REAL-TIME",
		  "id": "285449387",
		  "price": 27220,
		  "exchange": 6,
		  "conditions": [
			1
		  ]
		}
	  }`,
	`{
		"market_status": "open",
		"name": "Ethereum - United States Dollar",
		"ticker": "X:ETHUSD",
		"type": "crypto",
		"session": {
		  "change": -5.53,
		  "change_percent": -0.304,
		  "close": 1817.14,
		  "high": 1833.4,
		  "low": 1802.26,
		  "open": 1823.8,
		  "volume": 47673.72258305,
		  "previous_close": 1821.84,
		  "price": 1816.31
		},
		"last_trade": {
		  "participant_timestamp": 1684422449301037000,
		  "timeframe": "REAL-TIME",
		  "id": "451453400",
		  "price": 1816.31,
		  "exchange": 1,
		  "conditions": [
			1
		  ]
		}
	  }`,
	`{
		"market_status": "open",
		"name": "Flow - United States Dollar",
		"ticker": "X:FLOWUSD",
		"type": "crypto",
		"session": {
		  "change": -0.006,
		  "change_percent": -0.759,
		  "close": 0.784,
		  "high": 0.793,
		  "low": 0.779,
		  "open": 0.791,
		  "volume": 89094.52551417,
		  "previous_close": 0.79,
		  "price": 0.784
		},
		"last_trade": {
		  "participant_timestamp": 1684422355917759000,
		  "timeframe": "REAL-TIME",
		  "id": "1224329",
		  "price": 0.784,
		  "exchange": 1,
		  "conditions": [
			1
		  ]
		}
	  }`,
}
var forexSnapshotTestData = []string{
	`{
		"market_status": "open",
		"name": "United States dollar - Canadian dollar",
		"ticker": "C:USDCAD",
		"type": "fx",
		"session": {
		  "change": 0.00518,
		  "change_percent": 0.385,
		  "close": 1.35178,
		  "high": 1.35244,
		  "low": 1.34479,
		  "open": 1.34667,
		  "volume": 133712,
		  "previous_close": 1.34667
		},
		"last_quote": {
		  "last_updated": 1684431002000000000,
		  "timeframe": "REAL-TIME",
		  "ask": 1.35191,
		  "bid": 1.35185,
		  "exchange": 48
		}
	}`,
	`{
		"market_status": "open",
		"name": "United States dollar - Euro",
		"ticker": "C:USDEUR",
		"type": "fx",
		"session": {
		  "change": 0.00612,
		  "change_percent": 0.663,
		  "close": 0.92877,
		  "high": 0.92913,
		  "low": 0.9215,
		  "open": 0.92266,
		  "volume": 58275,
		  "previous_close": 0.92265
		},
		"last_quote": {
		  "last_updated": 1684430998000000000,
		  "timeframe": "REAL-TIME",
		  "ask": 0.92883,
		  "bid": 0.92877,
		  "exchange": 48
		}
	}`,
	`{
		"market_status": "open",
		"name": "United States dollar - Australian dollar",
		"ticker": "C:USDAUD",
		"type": "fx",
		"session": {
		  "change": 0.0104,
		  "change_percent": 0.692,
		  "close": 1.512722,
		  "high": 1.51398,
		  "low": 1.4965803,
		  "open": 1.502449,
		  "volume": 148378,
		  "previous_close": 1.502449
		},
		"last_quote": {
		  "last_updated": 1684431002000000000,
		  "timeframe": "REAL-TIME",
		  "ask": 1.51297,
		  "bid": 1.51281,
		  "exchange": 48
		}
	}`,
}
var indicesSnapshotTestData = []string{
	`{
		"TBD": "TBD"
	}`,
}

var partialSuccessWithStocksTestData = []string{
	`{
		"market_status": "late_trading",
		"name": "Apple Inc.",
		"session": {
			"change": -0.07,
			"change_percent": -0.0403,
			"close": 173.5,
			"early_trading_change": 0,
			"early_trading_change_percent": 0,
			"high": 173.85,
			"low": 172.11,
			"open": 172.48,
			"previous_close": 173.57,
			"price": 173.5,
			"volume": 50823329
		},
		"last_quote": {
			"ask": 173.34,
			"ask_size": 3,
			"bid": 173.32,
			"bid_size": 4,
			"last_updated": 1683577209434314800,
			"timeframe": "REAL-TIME"
		},
		"last_trade": {
			"conditions": [
				12,
				22
			],
			"exchange": 4,
			"id": "247862",
			"last_updated": 1683577205678289200,
			"price": 173.5,
			"size": 31535,
			"timeframe": "REAL-TIME"
		},
		"ticker": "AAPL",
		"type": "stocks"
	}`,
	`{
		"error": "NOT_ENTITLED",
		"message": "Not entitled to this ticker.",
		"ticker": "APy"
	}`,
}
