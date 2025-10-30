package models_test

import (
	"reflect"
	"testing"

	"github.com/massive-com/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListAggsParams(t *testing.T) {
	adjusted := true
	order := models.Asc
	limit := 50

	expect := models.ListAggsParams{
		Adjusted: &adjusted,
		Order:    &order,
		Limit:    &limit,
	}
	actual := models.ListAggsParams{}.WithAdjusted(adjusted).WithOrder(order).WithLimit(limit)
	checkParams(t, expect, *actual)
}

func TestGetAggsParamsMethods(t *testing.T) {
	adjusted := true
	order := models.Asc
	limit := 50

	expect := models.GetAggsParams{
		Adjusted: &adjusted,
		Order:    &order,
		Limit:    &limit,
	}
	actual := models.GetAggsParams{}.WithAdjusted(adjusted).WithOrder(order).WithLimit(limit)
	checkParams(t, expect, *actual)
}

func TestGetGroupedDailyAggsParamsMethods(t *testing.T) {
	adjusted := true
	includeOTC := true

	expect := models.GetGroupedDailyAggsParams{
		Adjusted:   &adjusted,
		IncludeOTC: &includeOTC,
	}
	actual := models.GetGroupedDailyAggsParams{}.WithAdjusted(adjusted).WithIncludeOTC(includeOTC)
	checkParams(t, expect, *actual)
}

func TestGetDailyOpenCloseAggParamsMethods(t *testing.T) {
	adjusted := true

	expect := models.GetDailyOpenCloseAggParams{
		Adjusted: &adjusted,
	}
	actual := models.GetDailyOpenCloseAggParams{}.WithAdjusted(adjusted)
	checkParams(t, expect, *actual)
}

func TestGetPreviousCloseAggParamsMethods(t *testing.T) {
	adjusted := true

	expect := models.GetPreviousCloseAggParams{
		Adjusted: &adjusted,
	}
	actual := models.GetPreviousCloseAggParams{}.WithAdjusted(adjusted)
	checkParams(t, expect, *actual)
}

func checkParams(t *testing.T, expect, actual interface{}) {
	for _, field := range reflect.VisibleFields(reflect.TypeOf(actual)) {
		assert.NotNil(t, field)
	}
	assert.Equal(t, expect, actual)
}
