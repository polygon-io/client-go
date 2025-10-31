package models_test

import (
	"testing"
	"time"

	"github.com/massive-com/client-go/v2/rest/models"
)

func TestListTradesParams(t *testing.T) {
	timestamp := models.Nanos(time.Date(2023, 3, 23, 0, 0, 0, 0, time.UTC))
	order := models.Asc
	limit := 100
	sort := models.Timestamp

	expect := models.ListTradesParams{
		TimestampEQ:  &timestamp,
		TimestampLT:  &timestamp,
		TimestampLTE: &timestamp,
		TimestampGT:  &timestamp,
		TimestampGTE: &timestamp,
		Order:        &order,
		Limit:        &limit,
		Sort:         &sort,
	}
	actual := models.ListTradesParams{}.
		WithTimestamp(models.EQ, timestamp).
		WithTimestamp(models.LT, timestamp).
		WithTimestamp(models.LTE, timestamp).
		WithTimestamp(models.GT, timestamp).
		WithTimestamp(models.GTE, timestamp).
		WithOrder(order).
		WithLimit(limit).
		WithSort(sort)

	checkParams(t, expect, *actual)
}
