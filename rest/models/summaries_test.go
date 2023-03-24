package models_test

import (
	"testing"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestGetSummaryParams(t *testing.T) {
	tickers := []string{"AAPL", "GOOG", "MSFT"}

	t.Run("Test WithTickerAnyOf", func(t *testing.T) {
		params := models.GetSummaryParams{}.WithTickerAnyOf(tickers...)
		assert.NotNil(t, params.TickerAnyOf)
		assert.Equal(t, "AAPL,GOOG,MSFT", *params.TickerAnyOf)
	})
}
