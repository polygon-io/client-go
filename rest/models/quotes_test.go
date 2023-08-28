package models_test

import (
	"testing"
	"time"

	"github.com/polygon-io/client-go/rest/models"
)

func TestListQuotesParams(t *testing.T) {
	timestamp := models.Nanos(time.Date(2023, 3, 23, 0, 0, 0, 0, time.UTC))
	year, month, day := 2023, time.March, 23
	order := models.Asc
	limit := 100
	sort := models.Timestamp

	expect := models.ListQuotesParams{
		TimestampEQ:  &timestamp,
		TimestampLT:  &timestamp,
		TimestampLTE: &timestamp,
		TimestampGT:  &timestamp,
		TimestampGTE: &timestamp,
		Order:        &order,
		Limit:        &limit,
		Sort:         &sort,
	}
	actual := models.ListQuotesParams{}.
		WithTimestamp(models.EQ, timestamp).
		WithTimestamp(models.LT, timestamp).
		WithTimestamp(models.LTE, timestamp).
		WithTimestamp(models.GT, timestamp).
		WithTimestamp(models.GTE, timestamp).
		WithDay(year, month, day).
		WithOrder(order).
		WithLimit(limit).
		WithSort(sort)

	checkParams(t, expect, *actual)
}
