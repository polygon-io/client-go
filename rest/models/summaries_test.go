package models_test

import (
	"testing"

	"github.com/massive-com/client-go/v2/rest/models"
)

func TestGetSummaryParams(t *testing.T) {
	tickers := "AAPL,GOOL,TSLA"
	expect := models.GetSummaryParams{
		TickerAnyOf: &tickers,
	}
	actual := models.GetSummaryParams{}.WithTickerAnyOf(tickers)
	checkParams(t, expect, *actual)
}
