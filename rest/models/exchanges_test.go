package models_test

import (
	"testing"

	"github.com/massive-com/client-go/rest/models"
)

func TestGetExchangesParams(t *testing.T) {
	assetClass := models.AssetStocks
	locale := models.US
	expect := models.GetExchangesParams{
		AssetClass: &assetClass,
		Locale:     &locale,
	}
	actual := models.GetExchangesParams{}.
		WithAssetClass(assetClass).
		WithLocale(locale)
	checkParams(t, expect, *actual)
}
