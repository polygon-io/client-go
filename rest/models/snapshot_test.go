package models_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
)

func TestGetAllTickersSnapshotParams(t *testing.T) {
	tickers := "AAPL,GOOL,TSLA"
	otc := false
	expect := models.GetAllTickersSnapshotParams{
		Tickers:    &tickers,
		IncludeOTC: &otc,
	}
	actual := models.GetAllTickersSnapshotParams{}.
		WithTickers(tickers).
		WithIncludeOTC(otc)

	checkParams(t, expect, *actual)
}

func TestGetIndicesSnapshotParams(t *testing.T) {
	tickers := "AAPL,GOOL,TSLA"
	expect := models.GetIndicesSnapshotParams{
		TickerAnyOf: &tickers,
	}
	actual := models.GetIndicesSnapshotParams{}.WithTickerAnyOf(tickers)
	checkParams(t, expect, *actual)
}

func TestListUniversalSnapshotsParams(t *testing.T) {
	ticker := "A"
	tickers := "AAPL,GOOL,TSLA"
	snapshot := "stocks"
	expect := models.ListUniversalSnapshotsParams{
		TickerAnyOf: &tickers,
		Ticker:      &ticker,
		TickerLT:    &ticker,
		TickerLTE:   &ticker,
		TickerGT:    &ticker,
		TickerGTE:   &ticker,
		Type:        &snapshot,
	}
	actual := models.ListUniversalSnapshotsParams{}.
		WithTickerAnyOf(tickers).
		WithTicker(ticker).
		WithTickersByComparison(models.LT, ticker).
		WithTickersByComparison(models.LTE, ticker).
		WithTickersByComparison(models.GT, ticker).
		WithTickersByComparison(models.GTE, ticker).
		WithType(snapshot)

	checkParams(t, expect, *actual)
}

func TestListOptionsChainParams(t *testing.T) {
	date := models.Date(time.Date(2023, 3, 23, 0, 0, 0, 0, time.Local))
	contractType := models.ContractCall
	strikePrice := 1.23
	limit := 100
	sort := models.TickerSymbol
	order := models.Asc
	expect := models.ListOptionsChainParams{
		StrikePrice:       &strikePrice,
		ContractType:      &contractType,
		ExpirationDateEQ:  &date,
		ExpirationDateLT:  &date,
		ExpirationDateLTE: &date,
		ExpirationDateGT:  &date,
		ExpirationDateGTE: &date,
		Limit:             &limit,
		Sort:              &sort,
		Order:             &order,
	}
	actual := models.ListOptionsChainParams{}.
		WithStrikePrice(strikePrice).
		WithContractType(contractType).
		WithExpirationDate(models.EQ, date).
		WithExpirationDate(models.LT, date).
		WithExpirationDate(models.LTE, date).
		WithExpirationDate(models.GT, date).
		WithExpirationDate(models.GTE, date).
		WithLimit(limit).
		WithSort(sort).
		WithOrder(order)

	checkParams(t, expect, *actual)
}

func TestGetTickerSnapshotParams(t *testing.T) {
	tests := []struct {
		name   string
		params models.GetTickerSnapshotParams
		want   string
	}{
		{
			name:   "All fields valid",
			params: models.GetTickerSnapshotParams{Locale: models.US, MarketType: models.Stocks, Ticker: "AAPL"},
			want:   "",
		},
		{
			name:   "Missing Locale",
			params: models.GetTickerSnapshotParams{MarketType: models.Stocks, Ticker: "AAPL"},
			want:   "Locale is required",
		},
		{
			name:   "Missing MarketType",
			params: models.GetTickerSnapshotParams{Locale: models.US, Ticker: "AAPL"},
			want:   "MarketType is required",
		},
		{
			name:   "Missing Ticker",
			params: models.GetTickerSnapshotParams{Locale: models.US, MarketType: models.Stocks},
			want:   "Ticker is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := ""
			if tt.params.Locale == "" {
				errMsg = "Locale is required"
			} else if tt.params.MarketType == "" {
				errMsg = "MarketType is required"
			} else if tt.params.Ticker == "" {
				errMsg = "Ticker is required"
			}

			if errMsg != tt.want {
				t.Errorf("Expected error: '%s', but got: '%s'", tt.want, errMsg)
			}
		})
	}
}

func TestGetGainersLosersSnapshotParamsOTC(t *testing.T) {
	otc := false
	expect := models.GetGainersLosersSnapshotParams{
		IncludeOTC: &otc,
	}
	actual := models.GetGainersLosersSnapshotParams{}.
		WithIncludeOTC(otc)

	checkParams(t, expect, *actual)
}

func TestGetGainersLosersSnapshotParams(t *testing.T) {
	tests := []struct {
		name   string
		params models.GetGainersLosersSnapshotParams
		want   string
	}{
		{
			name:   "All fields valid",
			params: models.GetGainersLosersSnapshotParams{Locale: models.US, MarketType: models.Stocks, Direction: models.Gainers},
			want:   "",
		},
		{
			name:   "Missing Locale",
			params: models.GetGainersLosersSnapshotParams{MarketType: models.Stocks, Direction: models.Gainers},
			want:   "Locale is required",
		},
		{
			name:   "Missing MarketType",
			params: models.GetGainersLosersSnapshotParams{Locale: models.US, Direction: models.Losers},
			want:   "MarketType is required",
		},
		{
			name:   "Missing Direction",
			params: models.GetGainersLosersSnapshotParams{Locale: models.US, MarketType: models.Stocks},
			want:   "Direction is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := ""
			if tt.params.Locale == "" {
				errMsg = "Locale is required"
			} else if tt.params.MarketType == "" {
				errMsg = "MarketType is required"
			} else if tt.params.Direction == "" {
				errMsg = "Direction is required"
			}

			if errMsg != tt.want {
				t.Errorf("Expected error: '%s', but got: '%s'", tt.want, errMsg)
			}
		})
	}
}

func TestGetOptionContractSnapshotParams(t *testing.T) {
	tests := []struct {
		name   string
		params models.GetOptionContractSnapshotParams
		want   string
	}{
		{
			name:   "All fields valid",
			params: models.GetOptionContractSnapshotParams{UnderlyingAsset: "EVRI", OptionContract: "O:EVRI240119C00002500"},
			want:   "",
		},
		{
			name:   "Missing UnderlyingAsset",
			params: models.GetOptionContractSnapshotParams{OptionContract: "O:EVRI240119C00002500"},
			want:   "UnderlyingAsset is required",
		},
		{
			name:   "Missing OptionContract",
			params: models.GetOptionContractSnapshotParams{UnderlyingAsset: "EVRI"},
			want:   "OptionContract is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := ""
			if tt.params.UnderlyingAsset == "" {
				errMsg = "UnderlyingAsset is required"
			} else if tt.params.OptionContract == "" {
				errMsg = "OptionContract is required"
			}

			if errMsg != tt.want {
				t.Errorf("Expected error: '%s', but got: '%s'", tt.want, errMsg)
			}
		})
	}
}

func TestGetCryptoFullBookSnapshotParams(t *testing.T) {
	tests := []struct {
		name   string
		params models.GetCryptoFullBookSnapshotParams
		want   string
	}{
		{
			name:   "Valid Ticker",
			params: models.GetCryptoFullBookSnapshotParams{Ticker: "X:BTCUSD"},
			want:   "",
		},
		{
			name:   "Missing Ticker",
			params: models.GetCryptoFullBookSnapshotParams{},
			want:   "Ticker is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errMsg := ""
			if tt.params.Ticker == "" {
				errMsg = "Ticker is required"
			}

			if errMsg != tt.want {
				t.Errorf("Expected error: '%s', but got: '%s'", tt.want, errMsg)
			}
		})
	}
}
