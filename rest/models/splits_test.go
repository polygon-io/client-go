package models_test

import (
	"testing"
	"time"

	"github.com/massive-com/client-go/v2/rest/models"
)

func TestListSplitsParams(t *testing.T) {
	ticker := "A"
	date := models.Date(time.Date(2023, 3, 23, 0, 0, 0, 0, time.Local))
	reverseSplit := true
	sort := models.TickerSymbol
	order := models.Asc
	limit := 100
	expect := models.ListSplitsParams{
		TickerEQ:         &ticker,
		TickerLT:         &ticker,
		TickerLTE:        &ticker,
		TickerGT:         &ticker,
		TickerGTE:        &ticker,
		ExecutionDateEQ:  &date,
		ExecutionDateLT:  &date,
		ExecutionDateLTE: &date,
		ExecutionDateGT:  &date,
		ExecutionDateGTE: &date,
		ReverseSplit:     &reverseSplit,
		Sort:             &sort,
		Order:            &order,
		Limit:            &limit,
	}
	actual := models.ListSplitsParams{}.
		WithTicker(models.EQ, ticker).
		WithTicker(models.LT, ticker).
		WithTicker(models.LTE, ticker).
		WithTicker(models.GT, ticker).
		WithTicker(models.GTE, ticker).
		WithExecutionDate(models.EQ, date).
		WithExecutionDate(models.LT, date).
		WithExecutionDate(models.LTE, date).
		WithExecutionDate(models.GT, date).
		WithExecutionDate(models.GTE, date).
		WithReverseSplit(reverseSplit).
		WithSort(sort).
		WithOrder(order).
		WithLimit(limit)

	checkParams(t, expect, *actual)
}
