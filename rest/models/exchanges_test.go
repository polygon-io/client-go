package models_test

import (
	"testing"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestGetExchangesParams_WithAssetClass(t *testing.T) {
	assetClass := models.AssetClass("test")
	params := models.GetExchangesParams{}

	updatedParams := params.WithAssetClass(assetClass)

	assert.NotNil(t, updatedParams.AssetClass)
	assert.Equal(t, assetClass, *updatedParams.AssetClass)
}

func TestGetExchangesParams_WithLocale(t *testing.T) {
	locale := models.MarketLocale("test")
	params := models.GetExchangesParams{}

	updatedParams := params.WithLocale(locale)

	assert.NotNil(t, updatedParams.Locale)
	assert.Equal(t, locale, *updatedParams.Locale)
}

func TestExchange(t *testing.T) {
	exchange := models.Exchange{
		Acronym:       "TEST",
		AssetClass:    "Equity",
		ID:            1,
		Locale:        "US",
		MIC:           "XMIC",
		Name:          "Test Exchange",
		OperatingMIC:  "XOPM",
		ParticipantID: "1",
		Type:          "Stock",
		URL:           "https://test.exchange",
	}

	assert.Equal(t, "TEST", exchange.Acronym)
	assert.Equal(t, "Equity", exchange.AssetClass)
	assert.Equal(t, int64(1), exchange.ID)
	assert.Equal(t, "US", exchange.Locale)
	assert.Equal(t, "XMIC", exchange.MIC)
	assert.Equal(t, "Test Exchange", exchange.Name)
	assert.Equal(t, "XOPM", exchange.OperatingMIC)
	assert.Equal(t, "1", exchange.ParticipantID)
	assert.Equal(t, "Stock", exchange.Type)
	assert.Equal(t, "https://test.exchange", exchange.URL)
}
