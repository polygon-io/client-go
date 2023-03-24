package models_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestGetMarketHolidaysResponse(t *testing.T) {
	holidays := models.GetMarketHolidaysResponse{
		{
			Exchange: "NYSE",
			Name:     "New Year's Day",
			Date:     models.Date(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)),
			Status:   "Closed",
			Open:     models.Time{},
			Close:    models.Time{},
		},
	}

	assert.Equal(t, "NYSE", holidays[0].Exchange)
	assert.Equal(t, "New Year's Day", holidays[0].Name)
	assert.Equal(t, models.Date(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)), holidays[0].Date)
	assert.Equal(t, "Closed", holidays[0].Status)
	assert.Equal(t, models.Time{}, holidays[0].Open)
	assert.Equal(t, models.Time{}, holidays[0].Close)
}

func TestGetMarketStatusResponse(t *testing.T) {
	status := models.GetMarketStatusResponse{
		AfterHours: true,
		Currencies: map[string]string{
			"USD": "US Dollar",
		},
		EarlyHours:    false,
		Exchanges:     map[string]string{"NYSE": "New York Stock Exchange"},
		IndicesGroups: map[string]string{"SPX": "S&P 500"},
		Market:        "US",
		ServerTime:    models.Time(time.Now()),
	}

	assert.True(t, status.AfterHours)
	assert.Equal(t, "US Dollar", status.Currencies["USD"])
	assert.False(t, status.EarlyHours)
	assert.Equal(t, "New York Stock Exchange", status.Exchanges["NYSE"])
	assert.Equal(t, "S&P 500", status.IndicesGroups["SPX"])
	assert.Equal(t, "US", status.Market)
	assert.NotNil(t, status.ServerTime)
}
