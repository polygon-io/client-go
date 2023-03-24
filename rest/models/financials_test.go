package models_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListStockFinancialsParams(t *testing.T) {
	ticker := "AAPL"
	cik := "0000320193"
	companyName := "Apple Inc."
	sic := "3674"
	date := models.Date(time.Date(2022, 2, 2, 0, 0, 0, 0, time.Local))
	timeframe := models.Timeframe("Q")
	includeSources := true
	order := models.Order("asc")
	limit := 50
	sort := models.Sort("period_of_report_date")

	params := models.ListStockFinancialsParams{}.
		WithTicker(ticker).
		WithCIK(cik).
		WithCompanyName(models.Full, companyName).
		WithSIC(sic).
		WithFilingDate(models.EQ, date).
		WithPeriodOfReportDate(models.GTE, date).
		WithTimeframe(timeframe).
		WithIncludeSources(includeSources).
		WithOrder(order).
		WithLimit(limit).
		WithSort(sort)

	assert.NotNil(t, params.Ticker)
	assert.NotNil(t, params.CIK)
	assert.NotNil(t, params.CompanyNameFull)
	assert.Nil(t, params.CompanyNameSearch)
	assert.NotNil(t, params.SIC)
	assert.NotNil(t, params.FilingDateEQ)
	assert.Nil(t, params.FilingDateLT)
	assert.Nil(t, params.FilingDateLTE)
	assert.Nil(t, params.FilingDateGT)
	assert.Nil(t, params.FilingDateGTE)
	assert.Nil(t, params.PeriodOfReportDateEQ)
	assert.Nil(t, params.PeriodOfReportDateLT)
	assert.Nil(t, params.PeriodOfReportDateLTE)
	assert.Nil(t, params.PeriodOfReportDateGT)
	assert.NotNil(t, params.PeriodOfReportDateGTE)
	assert.NotNil(t, params.Timeframe)
	assert.NotNil(t, params.IncludeSources)
	assert.NotNil(t, params.Order)
	assert.NotNil(t, params.Limit)
	assert.NotNil(t, params.Sort)

	assert.Equal(t, ticker, *params.Ticker)
	assert.Equal(t, cik, *params.CIK)
	assert.Equal(t, companyName, *params.CompanyNameFull)
	assert.Equal(t, sic, *params.SIC)
	assert.Equal(t, date, *params.FilingDateEQ)
	assert.Equal(t, date, *params.PeriodOfReportDateGTE)
	assert.Equal(t, timeframe, *params.Timeframe)
	assert.Equal(t, includeSources, *params.IncludeSources)
	assert.Equal(t, order, *params.Order)
	assert.Equal(t, limit, *params.Limit)
	assert.Equal(t, sort, *params.Sort)
}
