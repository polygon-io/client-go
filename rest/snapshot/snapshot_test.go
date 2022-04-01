package snapshot_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

var snapshotAAPL = models.TickerSnapshot{
	Day: models.DaySnapshot{
		Close:                 120.4229,
		High:                  120.53,
		Low:                   118.81,
		Open:                  119.62,
		Volume:                28727868,
		VolumeWeightedAverage: 119.725,
	},
	LastQuote: models.LastQuoteSnapshot{
		AskPrice:  120.47,
		BidPrice:  120.46,
		AskSize:   4,
		BidSize:   8,
		Timestamp: 1605195918507251700,
	},
	LastTrade: models.LastTradeSnapshot{
		Conditions: []int{114, 41},
		TradeID:    "158698",
		Price:      120.47,
		Size:       236,
		Timestamp:  1605195918306274000,
		ExchangeID: 10,
	},
	Minute: models.MinuteSnapshot{
		AccumulatedVolume:     28724441,
		Close:                 120.4201,
		High:                  120.468,
		Low:                   120.37,
		Open:                  120.435,
		Volume:                270796,
		VolumeWeightedAverage: 120.4129,
	},
	PrevDay: models.DaySnapshot{
		Close:                 119.4229,
		High:                  119.53,
		Low:                   118.81,
		Open:                  119.62,
		Volume:                28727868,
		VolumeWeightedAverage: 119.725,
	},
	Ticker:           "AAPL",
	TodaysChange:     0.98,
	TodaysChangePerc: 0.82,
	Updated:          1605195918306274000,
}

var expectedGetSnapshotTickerResponse = models.GetSnapshotTickerResponse{
	BaseResponse: client.BaseResponse{
		Status:       "OK",
		RequestID:    "cffb2db04ed53d1fdf2547f15c1ca14e",
		Count:        1,
		Message:      "",
		ErrorMessage: "",
	},
	Snapshot: snapshotAAPL,
}

var expectedListSnapshotAllTickers = models.ListSnapshotAllTickersResponse{
	BaseResponse: client.BaseResponse{
		Status:       "OK",
		RequestID:    "cffb2db04ed53d1fdf2547f15c1ca14e",
		Count:        1,
		Message:      "",
		ErrorMessage: "",
	},
	Snapshots: []models.TickerSnapshot{snapshotAAPL, snapshotAAPL},
}

var expectedListSnapshotGainersLosersResponse = models.ListSnapshotGainersLosersResponse{
	BaseResponse: client.BaseResponse{
		Status:       "OK",
		RequestID:    "cffb2db04ed53d1fdf2547f15c1ca14e",
		Count:        1,
		Message:      "",
		ErrorMessage: "",
	},
	Snapshots: []models.TickerSnapshot{snapshotAAPL, snapshotAAPL},
}

func TestGetSnapshotTicker(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/snapshot/locale/us/markets/stocks/tickers/AAPL",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedGetSnapshotTickerResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Snapshot.GetSnapshotTicker(context.Background(), models.GetSnapshotTickerParams{
		Ticker:     "AAPL",
		Locale:     "us",
		MarketType: "stocks",
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedGetSnapshotTickerResponse, res)
}

func TestListSnapshotAllTickers(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/snapshot/locale/us/markets/stocks/tickers",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedListSnapshotAllTickers)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Snapshot.ListSnapshotAllTickers(context.Background(), models.ListSnapshotAllTickersParams{
		Locale:     "us",
		MarketType: "stocks",
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedListSnapshotAllTickers, res)
}

func TestListSnapshotGainersLosers(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/snapshot/locale/us/markets/stocks/gainers",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedListSnapshotGainersLosersResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Snapshot.ListSnapshotGainersLosers(context.Background(), models.ListSnapshotGainersLosersParams{
		Locale:     "us",
		MarketType: "stocks",
		Direction:  "gainers",
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedListSnapshotGainersLosersResponse, res)
}

var snapshotAAPLOptionContract = models.OptionContractSnapshot{
	BreakEvenPrice: 171.075,
	Day: models.DayOptionContractSnapshot{
		Change:        -1.05,
		ChangePercent: -4.67,
		Close:         21.4,
		High:          22.49,
		LastUpdated:   1636520400000000000,
		Low:           21.35,
		Open:          22.49,
		PreviousClose: 22.45,
		Volume:        37,
		VWAP:          21.6741,
	},
	Details: models.OptionDetails{
		ContractType:      "call",
		ExerciseStyle:     "american",
		ExpirationDate:    "2023-06-16",
		SharesPerContract: 100,
		StrikePrice:       150,
		Ticker:            "O:AAPL230616C00150000",
	},
	Greeks: models.Greeks{
		Delta: 0.5520187372272933,
		Gamma: 0.00706756515659829,
		Theta: -0.018532772783847958,
		Vega:  0.7274811132998142,
	},
	ImpliedVolatility: 0.3048997097864957,
	LastQuote: models.LastQuoteOptionContractSnapshot{
		Ask:         21.25,
		AskSize:     110,
		Bid:         20.9,
		BidSize:     172,
		LastUpdated: 1636573458756383500,
		Midpoint:    21.075,
		Timeframe:   "REAL-TIME",
	},
	OpenInterest: 8921,
	UnderlyingAsset: models.UnderlyingAsset{
		ChangeToBreakEven: 23.123999999999995,
		LastUpdated:       1636573459862384600,
		Price:             147.951,
		Ticker:            "AAPL",
		Timeframe:         "REAL-TIME",
	},
}

var expectedGetSnapshotOptionContractResponse = models.GetSnapshotOptionContractResponse{
	BaseResponse: client.BaseResponse{
		Status:       "OK",
		RequestID:    "1bea8bbfb1ae1fe0d4c613c2b759d5be",
		Count:        1,
		Message:      "",
		ErrorMessage: "",
	},
	Results: snapshotAAPLOptionContract,
}

func TestGetSnapshotOptionContract(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/snapshot/options/AAPL/O:AAPL230616C00150000",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedGetSnapshotOptionContractResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Snapshot.GetSnapshotOptionContract(context.Background(), models.GetSnapshotOptionContractParams{
		UnderlyingAsset: "AAPL",
		OptionContract:  "O:AAPL230616C00150000",
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedGetSnapshotOptionContractResponse, res)
}

var snapshotXBTCUSDOrderBook = models.SnapshotTickerFullBook{
	AskCount: 593.1412981600005,
	Asks: []models.Ask{
		models.Ask{
			Price: 11454,
			ExchangeToShares: map[string]int{
				"2": 1,
			}},
		models.Ask{
			Price: 11455,
			ExchangeToShares: map[string]int{
				"2": 1,
			},
		}},
	BidCount: 694.951789670001,
	Bids: []models.Bid{
		models.Bid{
			Price: 11453,
			ExchangeToShares: map[string]int{
				"2": 1,
			}},
		models.Bid{
			Price: 11453,
			ExchangeToShares: map[string]int{
				"6": 1,
			},
		}},
	Spread:  1.00,
	Ticker:  "X:BTCUSD",
	Updated: 1605295074162,
}

var expectedListSnapshotTickerFullBookResponse = models.ListSnapshotTickerFullBookResponse{
	BaseResponse: client.BaseResponse{
		Status:       "OK",
		RequestID:    "1bea8bbfb1ae1fe0d4c613c2b759d5be",
		Count:        1,
		Message:      "",
		ErrorMessage: "",
	},
	Data: snapshotXBTCUSDOrderBook,
}

func TestListSnapshotTickerFullBook(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Aggs.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v2/snapshot/locale/global/markets/crypto/tickers/X:BTCUSD/book",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedListSnapshotTickerFullBookResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	res, err := c.Snapshot.ListSnapshotTickerFullBook(context.Background(), models.ListSnapshotTickerFullBookParams{
		Ticker: "X:BTCUSD",
	})

	assert.Nil(t, err)
	assert.Equal(t, &expectedListSnapshotTickerFullBookResponse, res)
}
