package models_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListTradesParams(t *testing.T) {
	params := models.ListTradesParams{Ticker: "AAPL"}

	// Test WithTimestamp
	paramsWithTimestamp := params.WithTimestamp(models.EQ, models.Nanos(time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)))
	assert.Equal(t,models.Nanos(time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)), *paramsWithTimestamp.TimestampEQ)

	// Test WithDay
	paramsWithDay := params.WithDay(2022, 2, 2)
	assert.Equal(t, models.Nanos(time.Date(2022, 2, 2, 0, 0, 0, 0, time.UTC)), *paramsWithDay.TimestampEQ)

	// Test WithOrder
	order := models.Asc
	paramsWithOrder := params.WithOrder(order)
	assert.Equal(t, &order, paramsWithOrder.Order)

	// Test WithLimit
	limit := 100
	paramsWithLimit := params.WithLimit(limit)
	assert.Equal(t, &limit, paramsWithLimit.Limit)

	// Test WithSort
	sort := models.Timestamp
	paramsWithSort := params.WithSort(sort)
	assert.Equal(t, &sort, paramsWithSort.Sort)
}

func TestGetLastTradeParams(t *testing.T) {
	params := models.GetLastTradeParams{Ticker: "AAPL"}
	assert.Equal(t, "AAPL", params.Ticker)
}

func TestGetLastCryptoTradeParams(t *testing.T) {
	params := models.GetLastCryptoTradeParams{From: "BTC", To: "USD"}
	assert.Equal(t, "BTC", params.From)
	assert.Equal(t, "USD", params.To)
}
