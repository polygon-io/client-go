package models_test

import (
	"testing"
	"time"

	models "github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListDividendsParams_WithMethods(t *testing.T) {
	date := models.Date(time.Date(2023, 3, 23, 0, 0, 0, 0, time.Local))
	freq := models.Annually
	divType := models.DividendCD
	order := models.Asc
	limit := 25
	sort := models.TickerSymbol

	params := models.ListDividendsParams{}.
		WithTicker(models.EQ, "AAPL").
		WithExDividendDate(models.LT, date).
		WithDeclarationDate(models.GTE, date).
		WithFrequency(freq).
		WithCashAmount(models.GT, 1.0).
		WithDividendType(divType).
		WithOrder(order).
		WithLimit(limit).
		WithSort(sort)

	assert.Equal(t, "AAPL", *params.TickerEQ)
	assert.Equal(t, date, *params.ExDividendDateLT)
	assert.Equal(t, date, *params.DeclarationDateGTE)
	assert.Equal(t, freq, *params.Frequency)
	assert.Equal(t, 1.0, *params.CashAmountGT)
	assert.Equal(t, divType, *params.DividendType)
	assert.Equal(t, order, *params.Order)
	assert.Equal(t, limit, *params.Limit)
	assert.Equal(t, sort, *params.Sort)
}
