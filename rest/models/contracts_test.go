package models_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/polygon-io/client-go/rest/models"
)

func TestGetOptionsContractParams(t *testing.T) {
	params := models.GetOptionsContractParams{
		Ticker: "TEST_TICKER",
	}
	assert.Equal(t, "TEST_TICKER", params.Ticker)
	date := models.Date(time.Date(2023, 3, 23, 0, 0, 0, 0, time.Local))
	params = *params.WithAsOf(date)
	assert.NotNil(t, params.AsOf)
	assert.Equal(t, &date, params.AsOf)
}

func TestListOptionsContractsParams(t *testing.T) {
	params := models.ListOptionsContractsParams{}
	date := models.Date(time.Date(2023, 3, 23, 0, 0, 0, 0, time.Local))
	comparatorCases := []models.Comparator{
		models.EQ,
		models.LT,
		models.LTE,
		models.GT,
		models.GTE,
	}

	for _, comparator := range comparatorCases {
		params = *params.WithUnderlyingTicker(comparator, "TEST_TICKER")
		params = *params.WithExpirationDate(comparator, date)
		params = *params.WithStrikePrice(comparator, 100.0)
	}

	params = *params.WithContractType("TEST_CONTRACT_TYPE")
	assert.NotNil(t, params.ContractType)
	assert.Equal(t, "TEST_CONTRACT_TYPE", *params.ContractType)

	params = *params.WithAsOf(date)
	assert.NotNil(t, params.AsOf)
	assert.Equal(t, &date, params.AsOf)

	params = *params.WithExpired(true)
	assert.NotNil(t, params.Expired)
	assert.Equal(t, true, *params.Expired)

	params = *params.WithSort(models.Sort("TEST_SORT"))
	assert.NotNil(t, params.Sort)
	assert.Equal(t, models.Sort("TEST_SORT"), *params.Sort)

	params = *params.WithOrder(models.Asc)
	assert.NotNil(t, params.Order)
	assert.Equal(t, models.Asc, *params.Order)

	params = *params.WithLimit(100)
	assert.NotNil(t, params.Limit)
	assert.Equal(t, 100, *params.Limit)
}
