package polygon_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestGetSummary(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()
	expectedSummaryResponse := `{
		"request_id": "abc123",
		"results": [
		 {
		  "branding": {
		   "icon_url": "https://api.polygon.io/icon.png",
		   "logo_url": "https://api.polygon.io/logo.svg"
		  },
		  "market_status": "closed",
		  "name": "Norwegian Cruise Lines",
		  "price": 22.3,
		  "session": {
		   "change": -1.05,
		   "change_percent": -4.67,
		   "close": 21.4,
		   "early_trading_change": -0.39,
		   "early_trading_change_percent": -0.07,
		   "high": 22.49,
		   "late_trading_change": 1.2,
		   "late_trading_change_percent": 3.92,
		   "low": 21.35,
		   "open": 22.49,
		   "previous_close": 22.45,
		   "volume": 37
		  },
		  "ticker": "NCLH",
		  "type": "stock"
		 },
		 {
		  "market_status": "closed",
		  "name": "NCLH $5 Call",
		  "options": {
		   "contract_type": "call",
		   "exercise_style": "american",
		   "expiration_date": "2022-10-14",
		   "shares_per_contract": 100,
		   "strike_price": 5,
		   "underlying_ticker": "NCLH"
		  },
		  "price": 6.6,
		  "session": {
		   "change": -0.05,
		   "change_percent": -1.07,
		   "close": 6.65,
		   "early_trading_change": -0.01,
		   "early_trading_change_percent": -0.03,
		   "high": 7.01,
		   "late_trading_change": -0.4,
		   "late_trading_change_percent": -0.02,
		   "low": 5.42,
		   "open": 6.7,
		   "previous_close": 6.71,
		   "volume": 67
		  },
		  "ticker": "O:NCLH221014C00005000",
		  "type": "option"
		 },
		 {
		  "market_status": "open",
		  "name": "Euro - United States Dollar",
		  "price": 0.97989,
		  "session": {
		   "change": -0.0001,
		   "change_percent": -0.67,
		   "close": 0.97989,
		   "high": 0.98999,
		   "low": 0.96689,
		   "open": 0.97889,
		   "previous_close": 0.98001
		  },
		  "ticker": "C:EURUSD",
		  "type": "forex"
		 },
		 {
		  "branding": {
		   "icon_url": "https://api.polygon.io/icon.png",
		   "logo_url": "https://api.polygon.io/logo.svg"
		  },
		  "market_status": "open",
		  "name": "Bitcoin - United States Dollar",
		  "price": 32154.68,
		  "session": {
		   "change": -201.23,
		   "change_percent": -0.77,
		   "close": 32154.68,
		   "high": 33124.28,
		   "low": 28182.88,
		   "open": 31129.32,
		   "previous_close": 33362.18
		  },
		  "ticker": "X:BTCUSD",
		  "type": "crypto"
		 },
		 {
		  "error": "NOT_FOUND",
		  "message": "Ticker not found.",
		  "ticker": "APx"
		 }
		],
		"status": "OK"
	   }`
	expectedGetSummaryUrl := "https://api.polygon.io/v1/summaries?ticker.any_of=NCLH%2CO%3ANCLH221014C00005000%2CC%3AEURUSD%2CX%3ABTCUSD%2CAPx"
	registerResponder(expectedGetSummaryUrl, expectedSummaryResponse)
	tickerAnyOf := []string{"NCLH", "O:NCLH221014C00005000", "C:EURUSD", "X:BTCUSD", "APx"}

	res, err := c.GetSummaries(context.Background(), models.GetSummaryParams{}.WithTickerAnyOf(tickerAnyOf...))
	assert.Nil(t, err)

	var expect models.GetSummaryResponse
	err = json.Unmarshal([]byte(expectedSummaryResponse), &expect)
	assert.Nil(t, err)
	assert.Equal(t, &expect, res)
}
