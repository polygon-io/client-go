package models_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
)

func TestListStockFinancialsParams(t *testing.T) {
	ticker := "A"
	date := models.Date(time.Date(2023, 3, 23, 0, 0, 0, 0, time.Local))
	timeframe := models.TFAnnual
	cik := "0001650729"
	name := "Apple"
	sic := "3570"
	sources := false
	order := models.Asc
	limit := 100
	sort := models.TickerSymbol
	
	expect := models.ListStockFinancialsParams{
		Ticker:                &ticker,
		CIK:                   &cik,
		CompanyNameFull:       &name,
		CompanyNameSearch:     &name,
		SIC:                   &sic,
		FilingDateEQ:          &date,
		FilingDateLT:          &date,
		FilingDateLTE:         &date,
		FilingDateGT:          &date,
		FilingDateGTE:         &date,
		PeriodOfReportDateEQ:  &date,
		PeriodOfReportDateLT:  &date,
		PeriodOfReportDateLTE: &date,
		PeriodOfReportDateGT:  &date,
		PeriodOfReportDateGTE: &date,
		Timeframe:             &timeframe,
		IncludeSources:        &sources,
		Order:                 &order,
		Limit:                 &limit,
		Sort:                  &sort,
	}
	actual := models.ListStockFinancialsParams{}.
		WithTicker(ticker).
		WithCIK(cik).
		WithCompanyName(models.Full, name).
		WithCompanyName(models.Search, name).
		WithSIC(sic).
		WithFilingDate(models.EQ, date).
		WithFilingDate(models.LT, date).
		WithFilingDate(models.LTE, date).
		WithFilingDate(models.GT, date).
		WithFilingDate(models.GTE, date).
		WithPeriodOfReportDate(models.EQ, date).
		WithPeriodOfReportDate(models.LT, date).
		WithPeriodOfReportDate(models.LTE, date).
		WithPeriodOfReportDate(models.GT, date).
		WithPeriodOfReportDate(models.GTE, date).
		WithTimeframe(timeframe).
		WithIncludeSources(sources).
		WithOrder(order).
		WithLimit(limit).
		WithSort(sort)

	checkParams(t, expect, *actual)
}
