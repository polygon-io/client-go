package models_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListSplitsParams(t *testing.T) {
	ticker := "AAPL"
	date := models.Date(time.Date(2022, 2, 2, 0, 0, 0, 0, time.Local))
	reverseSplit := true
	order := models.Asc
	limit := 20
	sort := models.TickerSymbol
	allComparators := [5]models.Comparator{models.EQ, models.LT, models.LTE, models.GT, models.GTE}

	t.Run("Test WithTicker", func(t *testing.T) {
		for _, c := range allComparators {
			params := models.ListSplitsParams{}.WithTicker(c, ticker)
			switch c {
			case models.EQ:
				assert.Equal(t, &ticker, params.TickerEQ)
			case models.LT:
				assert.Equal(t, &ticker, params.TickerLT)
			case models.LTE:
				assert.Equal(t, &ticker, params.TickerLTE)
			case models.GT:
				assert.Equal(t, &ticker, params.TickerGT)
			case models.GTE:
				assert.Equal(t, &ticker, params.TickerGTE)
			}
		}
	})

	t.Run("Test WithExecutionDate", func(t *testing.T) {
		for _, c := range allComparators {
			params := models.ListSplitsParams{}.WithExecutionDate(c, date)
			switch c {
			case models.EQ:
				assert.Equal(t, &date, params.ExecutionDateEQ)
			case models.LT:
				assert.Equal(t, &date, params.ExecutionDateLT)
			case models.LTE:
				assert.Equal(t, &date, params.ExecutionDateLTE)
			case models.GT:
				assert.Equal(t, &date, params.ExecutionDateGT)
			case models.GTE:
				assert.Equal(t, &date, params.ExecutionDateGTE)
			}
		}
	})

	t.Run("Test WithReverseSplit", func(t *testing.T) {
		params := models.ListSplitsParams{}.WithReverseSplit(reverseSplit)
		assert.Equal(t, &reverseSplit, params.ReverseSplit)
	})

	t.Run("Test WithOrder", func(t *testing.T) {
		params := models.ListSplitsParams{}.WithOrder(order)
		assert.Equal(t, &order, params.Order)
	})

	t.Run("Test WithLimit", func(t *testing.T) {
		params := models.ListSplitsParams{}.WithLimit(limit)
		assert.Equal(t, &limit, params.Limit)
	})

	t.Run("Test WithSort", func(t *testing.T) {
		params := models.ListSplitsParams{}.WithSort(sort)
		assert.Equal(t, &sort, params.Sort)
	})
}
