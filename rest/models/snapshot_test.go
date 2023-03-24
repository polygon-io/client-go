package models

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestGetAllTickersSnapshotParams_WithTickers(t *testing.T) {
	params := GetAllTickersSnapshotParams{}
	tickers := "AAPL,GOOG,AMZN"
	params = *params.WithTickers(tickers)
	assert.Equal(t, tickers, *params.Tickers)
}

func TestGetAllTickersSnapshotParams_WithIncludeOTC(t *testing.T) {
	params := GetAllTickersSnapshotParams{}
	otc := true
	params = *params.WithIncludeOTC(otc)
	assert.Equal(t, otc, *params.IncludeOTC)
}

func TestGetGainersLosersSnapshotParams_WithIncludeOTC(t *testing.T) {
	params := GetGainersLosersSnapshotParams{}
	otc := true
	params = *params.WithIncludeOTC(otc)
	assert.Equal(t, otc, *params.IncludeOTC)
}

func TestListOptionsChainParams_WithStrikePrice(t *testing.T) {
	params := ListOptionsChainParams{}
	strikePrice := 100.0
	params = *params.WithStrikePrice(strikePrice)
	assert.Equal(t, strikePrice, *params.StrikePrice)
}

func TestListOptionsChainParams_WithContractType(t *testing.T) {
	params := ListOptionsChainParams{}
	contractType := ContractCall
	params = *params.WithContractType(contractType)
	assert.Equal(t, contractType, *params.ContractType)
}

func TestListOptionsChainParams_WithLimit(t *testing.T) {
	params := ListOptionsChainParams{}
	limit := 50
	params = *params.WithLimit(limit)
	assert.Equal(t, limit, *params.Limit)
}

func TestListOptionsChainParams_WithExpirationDate(t *testing.T) {
	params := ListOptionsChainParams{}
	expirationDate := Date(time.Date(2022, 2, 2, 0, 0, 0, 0, time.Local))
	comparator := EQ
	params = *params.WithExpirationDate(comparator, expirationDate)
	assert.Equal(t, expirationDate, *params.ExpirationDateEQ)
}

func TestListOptionsChainParams_WithOrder(t *testing.T) {
	params := ListOptionsChainParams{}
	order := Asc
	params = *params.WithOrder(order)
	assert.Equal(t, order, *params.Order)
}

func TestListOptionsChainParams_WithSort(t *testing.T) {
	params := ListOptionsChainParams{}
	sort := TickerSymbol
	params = *params.WithSort(sort)
	assert.Equal(t, sort, *params.Sort)
}

func TestGetIndicesSnapshotParams_WithTickerAnyOf(t *testing.T) {
	params := GetIndicesSnapshotParams{}
	tickers := []string{"AAPL", "GOOG", "AMZN"}
	params = *params.WithTickerAnyOf(tickers...)
	assert.Equal(t, strings.Join(tickers, ","), *params.TickerAnyOf)
}
