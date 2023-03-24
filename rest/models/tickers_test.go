package models_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListTickersParams(t *testing.T) {
	params := models.ListTickersParams{}
	params = *params.WithTicker(models.EQ, "AAPL").
		WithType("Stock").
		WithMarket(models.AssetStocks).
		WithExchange("NASDAQ").
		WithCUSIP(12345).
		WithCIK(67890).
		WithDate(models.Date(time.Date(2022, 2, 2, 0, 0, 0, 0, time.Local))).
		WithActive(true).
		WithSearch("Apple").
		WithSort(models.TickerSymbol).
		WithOrder(models.Asc).
		WithLimit(100)

	assert.Equal(t, "AAPL", *params.TickerEQ)
	assert.Equal(t, "Stock", *params.Type)
	assert.Equal(t, models.AssetStocks, *params.Market)
	assert.Equal(t, "NASDAQ", *params.Exchange)
	assert.Equal(t, 12345, *params.CUSIP)
	assert.Equal(t, 67890, *params.CIK)
	assert.Equal(t, models.Date(time.Date(2022, 2, 2, 0, 0, 0, 0, time.Local)), *params.Date)
	assert.Equal(t, true, *params.Active)
	assert.Equal(t, "Apple", *params.Search)
	assert.Equal(t, models.TickerSymbol, *params.Sort)
	assert.Equal(t, models.Asc, *params.Order)
	assert.Equal(t, 100, *params.Limit)
}

func TestGetTickerDetailsParams(t *testing.T) {
	params := models.GetTickerDetailsParams{Ticker: "AAPL"}
	params = *params.WithDate(models.Date(time.Date(2022, 2, 2, 0, 0, 0, 0, time.Local)))

	assert.Equal(t, "AAPL", params.Ticker)
	assert.Equal(t, models.Date(time.Date(2022, 2, 2, 0, 0, 0, 0, time.Local)), *params.Date)
}

func TestListTickerNewsParams(t *testing.T) {
	millis := models.Millis(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))
	params := models.ListTickerNewsParams{}
	params = *params.WithTicker(models.EQ, "AAPL").
		WithPublishedUTC(models.GT, millis).
		WithSort(models.TickerSymbol).
		WithOrder(models.Asc).
		WithLimit(100)

	assert.Equal(t, "AAPL", *params.TickerEQ)
	assert.Equal(t, millis, *params.PublishedUtcGT)
	assert.Equal(t, models.TickerSymbol, *params.Sort)
	assert.Equal(t, models.Asc, *params.Order)
	assert.Equal(t, 100, *params.Limit)
}

func TestGetTickerTypesParams(t *testing.T) {
	params := models.GetTickerTypesParams{}
	params = *params.WithAssetClass(models.AssetStocks).WithLocale(models.US)

	assert.Equal(t, models.AssetStocks, *params.AssetClass)
	assert.Equal(t, models.US, *params.Locale)
}

func TestGetTickerEventsParams(t *testing.T) {
	params := models.GetTickerEventsParams{ID: "AAPL"}
	params = *params.WithTypes("ticker_change")

	assert.Equal(t, "AAPL", params.ID)
	assert.Equal(t, "ticker_change", *params.Types)
}
