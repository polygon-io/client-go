package massive_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	massive "github.com/massive-com/client-go/rest"
	"github.com/massive-com/client-go/rest/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListStockFinancials(t *testing.T) {
	c := massive.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	financial := `
	{
		"financials":{
			"comprehensive_income":{
			"other_comprehensive_income_loss":{
				"label":"Other Comprehensive Income/Loss",
				"value":-3.9e+06,
				"unit":"USD",
				"order":400
			},
			"comprehensive_income_loss":{
				"label":"Comprehensive Income/Loss",
				"value":4.1e+07,
				"unit":"USD",
				"order":100
			},
			"comprehensive_income_loss_attributable_to_noncontrolling_interest":{
				"label":"Comprehensive Income/Loss Attributable To Noncontrolling Interest",
				"value":900000,
				"unit":"USD",
				"order":200
			},
			"comprehensive_income_loss_attributable_to_parent":{
				"label":"Comprehensive Income/Loss Attributable To Parent",
				"value":4.01e+07,
				"unit":"USD",
				"order":300
			}
			},
			"income_statement":{
			"income_loss_from_continuing_operations_before_tax":{
				"label":"Income/Loss From Continuing Operations Before Tax",
				"value":5.6e+07,
				"unit":"USD",
				"order":1500
			},
			"nonoperating_income_loss":{
				"label":"Nonoperating Income/Loss",
				"value":-1.02e+07,
				"unit":"USD",
				"order":900
			},
			"revenues":{
				"label":"Revenues",
				"value":5.191e+08,
				"unit":"USD",
				"order":100
			},
			"income_loss_from_equity_method_investments":{
				"label":"Income/Loss From Equity Method Investments",
				"value":100000,
				"unit":"USD",
				"order":2100
			},
			"net_income_loss_attributable_to_noncontrolling_interest":{
				"label":"Net Income/Loss Attributable To Noncontrolling Interest",
				"value":800000,
				"unit":"USD",
				"order":3300
			},
			"income_tax_expense_benefit":{
				"label":"Income Tax Expense/Benefit",
				"value":1.12e+07,
				"unit":"USD",
				"order":2200
			},
			"basic_earnings_per_share":{
				"label":"Basic Earnings Per Share",
				"value":1.33,
				"unit":"USD / shares",
				"order":4200
			},
			"income_loss_from_continuing_operations_after_tax":{
				"label":"Income/Loss From Continuing Operations After Tax",
				"value":4.48e+07,
				"unit":"USD",
				"order":1400
			},
			"participating_securities_distributed_and_undistributed_earnings_loss_basic":{
				"label":"Participating Securities, Distributed And Undistributed Earnings/Loss, Basic",
				"value":0,
				"unit":"USD",
				"order":3800
			},
			"benefits_costs_expenses":{
				"label":"Benefits Costs and Expenses",
				"value":4.631e+08,
				"unit":"USD",
				"order":200
			},
			"gross_profit":{
				"label":"Gross Profit",
				"value":1.217e+08,
				"unit":"USD",
				"order":800
			},
			"costs_and_expenses":{
				"label":"Costs And Expenses",
				"value":4.631e+08,
				"unit":"USD",
				"order":600
			},
			"operating_expenses":{
				"label":"Operating Expenses",
				"value":5.55e+07,
				"unit":"USD",
				"order":1000
			},
			"preferred_stock_dividends_and_other_adjustments":{
				"label":"Preferred Stock Dividends And Other Adjustments",
				"value":0,
				"unit":"USD",
				"order":3900
			},
			"interest_expense_operating":{
				"label":"Interest Expense, Operating",
				"value":9.8e+06,
				"unit":"USD",
				"order":2700
			},
			"cost_of_revenue":{
				"label":"Cost Of Revenue",
				"value":3.974e+08,
				"unit":"USD",
				"order":300
			},
			"diluted_earnings_per_share":{
				"label":"Diluted Earnings Per Share",
				"value":1.33,
				"unit":"USD / shares",
				"order":4300
			},
			"net_income_loss_available_to_common_stockholders_basic":{
				"label":"Net Income/Loss Available To Common Stockholders, Basic",
				"value":4.41e+07,
				"unit":"USD",
				"order":3700
			},
			"net_income_loss":{
				"label":"Net Income/Loss",
				"value":4.49e+07,
				"unit":"USD",
				"order":3200
			},
			"operating_income_loss":{
				"label":"Operating Income/Loss",
				"value":6.62e+07,
				"unit":"USD",
				"order":1100
			},
			"income_loss_before_equity_method_investments":{
				"label":"Income/Loss Before Equity Method Investments",
				"value":5.6e+07,
				"unit":"USD",
				"order":1300
			},
			"net_income_loss_attributable_to_parent":{
				"label":"Net Income/Loss Attributable To Parent",
				"value":4.41e+07,
				"unit":"USD",
				"order":3500
			}
			},
			"balance_sheet":{
			"equity":{
				"label":"Equity",
				"value":1.603e+09,
				"unit":"USD",
				"order":1400
			},
			"current_assets":{
				"label":"Current Assets",
				"value":1.0903e+09,
				"unit":"USD",
				"order":200
			},
			"liabilities":{
				"label":"Liabilities",
				"value":1.8228e+09,
				"unit":"USD",
				"order":600
			},
			"fixed_assets":{
				"label":"Fixed Assets",
				"value":1.0462e+09,
				"unit":"USD",
				"order":400
			},
			"equity_attributable_to_parent":{
				"label":"Equity Attributable To Parent",
				"value":1.562e+09,
				"unit":"USD",
				"order":1600
			},
			"noncurrent_assets":{
				"label":"Noncurrent Assets",
				"value":2.3355e+09,
				"unit":"USD",
				"order":300
			},
			"equity_attributable_to_noncontrolling_interest":{
				"label":"Equity Attributable To Noncontrolling Interest",
				"value":4.1e+07,
				"unit":"USD",
				"order":1500
			},
			"assets":{
				"label":"Assets",
				"value":3.4258e+09,
				"unit":"USD",
				"order":100
			},
			"liabilities_and_equity":{
				"label":"Liabilities And Equity",
				"value":3.4258e+09,
				"unit":"USD",
				"order":1900
			},
			"noncurrent_liabilities":{
				"label":"Noncurrent Liabilities",
				"value":1.3677e+09,
				"unit":"USD",
				"order":800
			},
			"other_than_fixed_noncurrent_assets":{
				"label":"Other Than Fixed Noncurrent Assets",
				"value":1.2893e+09,
				"unit":"USD",
				"order":500
			},
			"current_liabilities":{
				"label":"Current Liabilities",
				"value":4.551e+08,
				"unit":"USD",
				"order":700
			}
			},
			"cash_flow_statement":{
			"net_cash_flow_from_operating_activities":{
				"label":"Net Cash Flow From Operating Activities",
				"value":300000,
				"unit":"USD",
				"order":100
			},
			"exchange_gains_losses":{
				"label":"Exchange Gains/Losses",
				"value":-3.6e+06,
				"unit":"USD",
				"order":1000
			},
			"net_cash_flow_from_investing_activities":{
				"label":"Net Cash Flow From Investing Activities",
				"value":-1.71e+07,
				"unit":"USD",
				"order":400
			},
			"net_cash_flow_continuing":{
				"label":"Net Cash Flow, Continuing",
				"value":-7.9e+06,
				"unit":"USD",
				"order":1200
			},
			"net_cash_flow":{
				"label":"Net Cash Flow",
				"value":-1.15e+07,
				"unit":"USD",
				"order":1100
			},
			"net_cash_flow_from_financing_activities":{
				"label":"Net Cash Flow From Financing Activities",
				"value":8.9e+06,
				"unit":"USD",
				"order":700
			},
			"net_cash_flow_from_operating_activities_continuing":{
				"label":"Net Cash Flow From Operating Activities, Continuing",
				"value":300000,
				"unit":"USD",
				"order":200
			},
			"net_cash_flow_from_financing_activities_continuing":{
				"label":"Net Cash Flow From Financing Activities, Continuing",
				"value":8.9e+06,
				"unit":"USD",
				"order":800
			},
			"net_cash_flow_from_investing_activities_continuing":{
				"label":"Net Cash Flow From Investing Activities, Continuing",
				"value":-1.71e+07,
				"unit":"USD",
				"order":500
			}
			}
		},
		"start_date":"2022-01-01",
		"end_date":"2022-04-03",
		"filing_date":"2022-04-29",
		"cik":"0000891014",
		"company_name":"MINERALS TECHNOLOGIES INC.",
		"fiscal_period":"Q1",
		"fiscal_year":"2022",
		"source_filing_url":"https://api.massive.com/v1/reference/sec/filings/0000891014-22-000022",
		"source_filing_file_url":"https://api.massive.com/v1/reference/sec/filings/0000891014-22-000022/files/form10q_htm.xml"
	}
	`

	expectedResponse := `{
	"status": "OK",
	"request_id":"874d62dbbce4b437bde7885d44a6be36",
	"count":1,
	"next_url":"https://api.massive.com/vX/reference/financials?cursor=YXA9MjAyMjA0MDMmYXM9MDAwMDg5MTAxNC0yMi0wMDAwMjImaGFzX3hicmw9dHJ1ZSZsaW1pdD0xJnNvcnQ9cGVyaW9kX29mX3JlcG9ydF9kYXRlJnR5cGU9MTAtUQ",
	"results": [
` + indent(true, financial, "\t\t") + `
	]
}`

	registerResponder("https://api.massive.com/vX/reference/financials?ticker=MTX", expectedResponse)
	registerResponder("https://api.massive.com/vX/reference/financials?cursor=YXA9MjAyMjA0MDMmYXM9MDAwMDg5MTAxNC0yMi0wMDAwMjImaGFzX3hicmw9dHJ1ZSZsaW1pdD0xJnNvcnQ9cGVyaW9kX29mX3JlcG9ydF9kYXRlJnR5cGU9MTAtUQ", "{}")
	iter := c.VX.ListStockFinancials(context.Background(), models.ListStockFinancialsParams{}.WithTicker("MTX"))

	// iter creation
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Item())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	var expect models.StockFinancial
	err := json.Unmarshal([]byte(financial), &expect)
	assert.Nil(t, err)
	assert.Equal(t, expect, iter.Item())

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}

func TestGetTickerEvents(t *testing.T) {
	c := massive.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	events := `
	{
	  "events": [
		{
		  "date": "2022-06-09",
		  "ticker_change": {
			"ticker": "META"
		  },
		  "type": "ticker_change"
		},
		{
		  "date": "2012-05-18",
		  "ticker_change": {
			"ticker": "FB"
		  },
		  "type": "ticker_change"
		}
	  ],
	  "name": "Meta Platforms, Inc. Class A Common Stock"
	}
	`

	expectedResponse := `{
	"status": "OK",
	"request_id":"874d62dbbce4b437bde7885d44a6be36",
	"count":1,
	"next_url":"https://api.massive.com/vX/reference/financials?cursor=YXA9MjAyMjA0MDMmYXM9MDAwMDg5MTAxNC0yMi0wMDAwMjImaGFzX3hicmw9dHJ1ZSZsaW1pdD0xJnNvcnQ9cGVyaW9kX29mX3JlcG9ydF9kYXRlJnR5cGU9MTAtUQ",
	"results": [
` + indent(true, events, "\t\t") + `
	]
}`

	registerResponder("https://api.massive.com/vX/reference/tickers/META/events", expectedResponse)

	res, err := c.VX.GetTickerEvents(context.Background(), &models.GetTickerEventsParams{ID: "META"})
	require.NoError(t, err)

	var expect models.GetTickerEventsResponse
	err = json.Unmarshal([]byte(expectedResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}
