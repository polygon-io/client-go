package models_test

import (
	"testing"
	"time"

	"github.com/massive-com/client-go/v2/rest/models"
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
		StrikePriceLT:     &strikePrice,
		StrikePriceLTE:    &strikePrice,
		StrikePriceGT:     &strikePrice,
		StrikePriceGTE:    &strikePrice,
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
		WithStrikePrice(models.EQ, strikePrice).
		WithStrikePrice(models.LT, strikePrice).
		WithStrikePrice(models.LTE, strikePrice).
		WithStrikePrice(models.GT, strikePrice).
		WithStrikePrice(models.GTE, strikePrice).
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
