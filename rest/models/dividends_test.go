package models_test

import (
	"testing"
	"time"

	"github.com/massive-com/client-go/v2/rest/models"
)

func TestListDividendsParams(t *testing.T) {
	ticker := "A"
	date := models.Date(time.Date(2023, 3, 23, 0, 0, 0, 0, time.Local))
	cash := 1.25
	frequency := models.Annually
	dividendType := models.DividendCD
	order := models.Asc
	limit := 100
	sort := models.TickerSymbol
	expect := models.ListDividendsParams{
		TickerEQ:           &ticker,
		TickerLT:           &ticker,
		TickerLTE:          &ticker,
		TickerGT:           &ticker,
		TickerGTE:          &ticker,
		ExDividendDateEQ:   &date,
		ExDividendDateLT:   &date,
		ExDividendDateLTE:  &date,
		ExDividendDateGT:   &date,
		ExDividendDateGTE:  &date,
		DeclarationDateEQ:  &date,
		DeclarationDateLT:  &date,
		DeclarationDateLTE: &date,
		DeclarationDateGT:  &date,
		DeclarationDateGTE: &date,
		PayDateEQ:          &date,
		PayDateLT:          &date,
		PayDateLTE:         &date,
		PayDateGT:          &date,
		PayDateGTE:         &date,
		CashAmountEQ:       &cash,
		CashAmountLT:       &cash,
		CashAmountLTE:      &cash,
		CashAmountGT:       &cash,
		CashAmountGTE:      &cash,
		Frequency:          &frequency,
		DividendType:       &dividendType,
		Order:              &order,
		Limit:              &limit,
		Sort:               &sort,
	}
	actual := models.ListDividendsParams{}.
		WithTicker(models.EQ, ticker).
		WithTicker(models.LT, ticker).
		WithTicker(models.LTE, ticker).
		WithTicker(models.GT, ticker).
		WithTicker(models.GTE, ticker).
		WithExDividendDate(models.EQ, date).
		WithExDividendDate(models.LT, date).
		WithExDividendDate(models.LTE, date).
		WithExDividendDate(models.GT, date).
		WithExDividendDate(models.GTE, date).
		WithDeclarationDate(models.EQ, date).
		WithDeclarationDate(models.LT, date).
		WithDeclarationDate(models.LTE, date).
		WithDeclarationDate(models.GT, date).
		WithDeclarationDate(models.GTE, date).
		WithPayDate(models.EQ, date).
		WithPayDate(models.LT, date).
		WithPayDate(models.LTE, date).
		WithPayDate(models.GT, date).
		WithPayDate(models.GTE, date).
		WithCashAmount(models.EQ, cash).
		WithCashAmount(models.LT, cash).
		WithCashAmount(models.LTE, cash).
		WithCashAmount(models.GT, cash).
		WithCashAmount(models.GTE, cash).
		WithFrequency(models.Annually).
		WithDividendType(models.DividendCD).
		WithOrder(order).
		WithLimit(limit).
		WithSort(sort)

	checkParams(t, expect, *actual)
}
