package models_test

import (
	"testing"
	"time"

	"github.com/massive-com/client-go/v2/rest/models"
)

func TestListTickersParams(t *testing.T) {
	ticker := "A"
	assetType := "CS"
	assetMarket := models.AssetStocks
	cik := 1650729
	name := "Apple"
	date := models.Date(time.Date(2023, 3, 23, 0, 0, 0, 0, time.Local))
	active := true
	sort := models.TickerSymbol
	order := models.Asc
	limit := 100

	expect := models.ListTickersParams{
		TickerEQ:  &ticker,
		TickerLT:  &ticker,
		TickerLTE: &ticker,
		TickerGT:  &ticker,
		TickerGTE: &ticker,
		Type:      &assetType,
		Market:    &assetMarket,
		CIK:       &cik,
		Search:    &name,
		Date:      &date,
		Active:    &active,
		Sort:      &sort,
		Order:     &order,
		Limit:     &limit,
	}
	actual := models.ListTickersParams{}.
		WithTicker(models.EQ, ticker).
		WithTicker(models.LT, ticker).
		WithTicker(models.LTE, ticker).
		WithTicker(models.GT, ticker).
		WithTicker(models.GTE, ticker).
		WithType(assetType).
		WithMarket(assetMarket).
		WithCIK(cik).
		WithSearch(name).
		WithDate(date).
		WithActive(active).
		WithSort(sort).
		WithOrder(order).
		WithLimit(limit)

	checkParams(t, expect, *actual)
}

func TestGetTickerDetailsParams(t *testing.T) {
	date := models.Date(time.Date(2023, 3, 23, 0, 0, 0, 0, time.Local))

	expect := models.GetTickerDetailsParams{
		Date: &date,
	}
	actual := models.GetTickerDetailsParams{}.
		WithDate(date)

	checkParams(t, expect, *actual)
}

func TestListTickerNewsParams(t *testing.T) {
	ticker := "A"
	date := models.Millis(time.Date(2022, 7, 25, 0, 0, 0, 0, time.UTC))
	sort := models.TickerSymbol
	order := models.Asc
	limit := 100

	expect := models.ListTickerNewsParams{
		TickerEQ:        &ticker,
		TickerLT:        &ticker,
		TickerLTE:       &ticker,
		TickerGT:        &ticker,
		TickerGTE:       &ticker,
		PublishedUtcEQ:  &date,
		PublishedUtcLT:  &date,
		PublishedUtcLTE: &date,
		PublishedUtcGT:  &date,
		PublishedUtcGTE: &date,
		Sort:            &sort,
		Order:           &order,
		Limit:           &limit,
	}
	actual := models.ListTickerNewsParams{}.
		WithTicker(models.EQ, ticker).
		WithTicker(models.LT, ticker).
		WithTicker(models.LTE, ticker).
		WithTicker(models.GT, ticker).
		WithTicker(models.GTE, ticker).
		WithPublishedUTC(models.EQ, date).
		WithPublishedUTC(models.LT, date).
		WithPublishedUTC(models.LTE, date).
		WithPublishedUTC(models.GT, date).
		WithPublishedUTC(models.GTE, date).
		WithSort(sort).
		WithOrder(order).
		WithLimit(limit)

	checkParams(t, expect, *actual)
}

func TestGetTickerTypesParams(t *testing.T) {
	assetClass := models.AssetStocks
	locale := models.US

	expect := models.GetTickerTypesParams{
		AssetClass: &assetClass,
		Locale:     &locale,
	}
	actual := models.GetTickerTypesParams{}.
		WithAssetClass(assetClass).
		WithLocale(locale)
	checkParams(t, expect, *actual)
}

func TestGetTickerEventsParams(t *testing.T) {
	types := "ticker_change"
	expect := models.GetTickerEventsParams{
		Types: &types,
	}
	actual := models.GetTickerEventsParams{}.
		WithTypes(types)

	checkParams(t, expect, *actual)
}
