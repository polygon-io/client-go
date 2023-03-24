package models_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListQuotesParams(t *testing.T) {
	ticker := "AAPL"
	limit := 100
	order := models.Asc
	sort := models.Timestamp

	params := models.ListQuotesParams{
		Ticker: ticker,
	}

	// Test WithTimestamp
	params = *params.WithTimestamp(models.EQ, models.Nanos(time.Now()))
	assert.NotNil(t, params.TimestampEQ)

	// Test WithDay
	params = *params.WithDay(2022, 2, 2)
	assert.NotNil(t, params.TimestampEQ)

	// Test WithOrder
	params = *params.WithOrder(order)
	assert.NotNil(t, params.Order)

	// Test WithLimit
	params = *params.WithLimit(limit)
	assert.NotNil(t, params.Limit)

	// Test WithSort
	params = *params.WithSort(sort)
	assert.NotNil(t, params.Sort)
}

func TestGetLastQuoteParams(t *testing.T) {
	ticker := "AAPL"
	params := models.GetLastQuoteParams{
		Ticker: ticker,
	}
	assert.Equal(t, params.Ticker, ticker)
}

func TestGetLastForexQuoteParams(t *testing.T) {
	from := "USD"
	to := "EUR"

	params := models.GetLastForexQuoteParams{
		From: from,
		To:   to,
	}

	assert.Equal(t, params.From, from)
	assert.Equal(t, params.To, to)
}

func TestGetRealTimeCurrencyConversionParams(t *testing.T) {
	from := "USD"
	to := "EUR"

	params := models.GetRealTimeCurrencyConversionParams{
		From: from,
		To:   to,
	}

	assert.Equal(t, params.From, from)
	assert.Equal(t, params.To, to)
}

func TestQuote(t *testing.T) {
	quote := models.Quote{}
	assert.NotNil(t, quote)
}

func TestLastQuote(t *testing.T) {
	lastQuote := models.LastQuote{}
	assert.NotNil(t, lastQuote)
}

func TestForexQuote(t *testing.T) {
	forexQuote := models.ForexQuote{}
	assert.NotNil(t, forexQuote)
}
