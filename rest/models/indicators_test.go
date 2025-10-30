package models_test

import (
	"testing"
	"time"

	"github.com/massive-com/client-go/rest/models"
)

func TestGetSMAParams(t *testing.T) {
	timespan := models.Week
	timestamp := models.Millis(time.Date(2022, 7, 25, 0, 0, 0, 0, time.UTC))
	series := models.Close
	expand := true
	adjusted := true
	order := models.Asc
	limit := 100
	window := 5
	expect := models.GetSMAParams{
		Timespan:         &timespan,
		TimestampEQ:      &timestamp,
		TimestampLT:      &timestamp,
		TimestampLTE:     &timestamp,
		TimestampGT:      &timestamp,
		TimestampGTE:     &timestamp,
		SeriesType:       &series,
		ExpandUnderlying: &expand,
		Adjusted:         &adjusted,
		Order:            &order,
		Limit:            &limit,
		Window:           &window,
	}
	actual := models.GetSMAParams{}.
		WithTimespan(timespan).
		WithTimestamp(models.EQ, timestamp).
		WithTimestamp(models.LT, timestamp).
		WithTimestamp(models.LTE, timestamp).
		WithTimestamp(models.GT, timestamp).
		WithTimestamp(models.GTE, timestamp).
		WithSeriesType(series).
		WithExpandUnderlying(expand).
		WithAdjusted(adjusted).
		WithOrder(order).
		WithLimit(limit).
		WithWindow(window)

	checkParams(t, expect, *actual)
}

func TestGetEMAParams(t *testing.T) {
	timespan := models.Week
	timestamp := models.Millis(time.Date(2022, 7, 25, 0, 0, 0, 0, time.UTC))
	series := models.Close
	expand := true
	adjusted := true
	order := models.Asc
	limit := 100
	window := 5
	expect := models.GetEMAParams{
		Timespan:         &timespan,
		TimestampEQ:      &timestamp,
		TimestampLT:      &timestamp,
		TimestampLTE:     &timestamp,
		TimestampGT:      &timestamp,
		TimestampGTE:     &timestamp,
		SeriesType:       &series,
		ExpandUnderlying: &expand,
		Adjusted:         &adjusted,
		Order:            &order,
		Limit:            &limit,
		Window:           &window,
	}
	actual := models.GetEMAParams{}.
		WithTimespan(timespan).
		WithTimestamp(models.EQ, timestamp).
		WithTimestamp(models.LT, timestamp).
		WithTimestamp(models.LTE, timestamp).
		WithTimestamp(models.GT, timestamp).
		WithTimestamp(models.GTE, timestamp).
		WithSeriesType(series).
		WithExpandUnderlying(expand).
		WithAdjusted(adjusted).
		WithOrder(order).
		WithLimit(limit).
		WithWindow(window)

	checkParams(t, expect, *actual)
}

func TestGetRSIParams(t *testing.T) {
	timespan := models.Week
	timestamp := models.Millis(time.Date(2022, 7, 25, 0, 0, 0, 0, time.UTC))
	series := models.Close
	expand := true
	adjusted := true
	order := models.Asc
	limit := 100
	window := 5
	expect := models.GetRSIParams{
		Timespan:         &timespan,
		TimestampEQ:      &timestamp,
		TimestampLT:      &timestamp,
		TimestampLTE:     &timestamp,
		TimestampGT:      &timestamp,
		TimestampGTE:     &timestamp,
		SeriesType:       &series,
		ExpandUnderlying: &expand,
		Adjusted:         &adjusted,
		Order:            &order,
		Limit:            &limit,
		Window:           &window,
	}
	actual := models.GetRSIParams{}.
		WithTimespan(timespan).
		WithTimestamp(models.EQ, timestamp).
		WithTimestamp(models.LT, timestamp).
		WithTimestamp(models.LTE, timestamp).
		WithTimestamp(models.GT, timestamp).
		WithTimestamp(models.GTE, timestamp).
		WithSeriesType(series).
		WithExpandUnderlying(expand).
		WithAdjusted(adjusted).
		WithOrder(order).
		WithLimit(limit).
		WithWindow(window)

	checkParams(t, expect, *actual)
}

func TestGetMACDParams(t *testing.T) {
	timespan := models.Week
	timestamp := models.Millis(time.Date(2022, 7, 25, 0, 0, 0, 0, time.UTC))
	series := models.Close
	expand := true
	adjusted := true
	order := models.Asc
	limit := 100
	shortWindow := 12
	longWindow := 26
	signalWindow := 9
	expect := models.GetMACDParams{
		Timespan:         &timespan,
		TimestampEQ:      &timestamp,
		TimestampLT:      &timestamp,
		TimestampLTE:     &timestamp,
		TimestampGT:      &timestamp,
		TimestampGTE:     &timestamp,
		SeriesType:       &series,
		ExpandUnderlying: &expand,
		Adjusted:         &adjusted,
		Order:            &order,
		Limit:            &limit,
		ShortWindow:      &shortWindow,
		LongWindow:       &longWindow,
		SignalWindow:     &signalWindow,
	}
	actual := models.GetMACDParams{}.
		WithTimespan(timespan).
		WithTimestamp(models.EQ, timestamp).
		WithTimestamp(models.LT, timestamp).
		WithTimestamp(models.LTE, timestamp).
		WithTimestamp(models.GT, timestamp).
		WithTimestamp(models.GTE, timestamp).
		WithSeriesType(series).
		WithExpandUnderlying(expand).
		WithAdjusted(adjusted).
		WithOrder(order).
		WithLimit(limit).
		WithShortWindow(shortWindow).
		WithLongWindow(longWindow).
		WithSignalWindow(signalWindow)

	checkParams(t, expect, *actual)
}
