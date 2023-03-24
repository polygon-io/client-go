package models_test

import (
	"testing"

	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListAggsParamsMethods(t *testing.T) {
	params := models.ListAggsParams{}
	assert.Nil(t, params.Adjusted)
	assert.Nil(t, params.Order)
	assert.Nil(t, params.Limit)

	params = *params.WithAdjusted(true)
	assert.NotNil(t, params.Adjusted)
	assert.Equal(t, true, *params.Adjusted)

	params = *params.WithOrder(models.Asc)
	assert.NotNil(t, params.Order)
	assert.Equal(t, models.Asc, *params.Order)

	params = *params.WithLimit(50)
	assert.NotNil(t, params.Limit)
	assert.Equal(t, 50, *params.Limit)
}

func TestGetAggsParamsMethods(t *testing.T) {
	params := models.GetAggsParams{}
	assert.Nil(t, params.Adjusted)
	assert.Nil(t, params.Order)
	assert.Nil(t, params.Limit)

	params = *params.WithAdjusted(true)
	assert.NotNil(t, params.Adjusted)
	assert.Equal(t, true, *params.Adjusted)

	params = *params.WithOrder(models.Asc)
	assert.NotNil(t, params.Order)
	assert.Equal(t, models.Asc, *params.Order)

	params = *params.WithLimit(50)
	assert.NotNil(t, params.Limit)
	assert.Equal(t, 50, *params.Limit)
}

func TestGetGroupedDailyAggsParamsMethods(t *testing.T) {
	params := models.GetGroupedDailyAggsParams{}
	assert.Nil(t, params.Adjusted)
	assert.Nil(t, params.IncludeOTC)

	params = *params.WithAdjusted(true)
	assert.NotNil(t, params.Adjusted)
	assert.Equal(t, true, *params.Adjusted)

	params = *params.WithIncludeOTC(true)
	assert.NotNil(t, params.IncludeOTC)
	assert.Equal(t, true, *params.IncludeOTC)
}

func TestGetDailyOpenCloseAggParamsMethods(t *testing.T) {
	params := models.GetDailyOpenCloseAggParams{}
	assert.Nil(t, params.Adjusted)

	params = *params.WithAdjusted(true)
	assert.NotNil(t, params.Adjusted)
	assert.Equal(t, true, *params.Adjusted)
}

func TestGetPreviousCloseAggParamsMethods(t *testing.T) {
	params := models.GetPreviousCloseAggParams{}
	assert.Nil(t, params.Adjusted)

	params = *params.WithAdjusted(true)
	assert.NotNil(t, params.Adjusted)
	assert.Equal(t, true, *params.Adjusted)
}
