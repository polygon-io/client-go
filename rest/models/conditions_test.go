package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/polygon-io/client-go/rest/models"
)

func TestListConditionsParams(t *testing.T) {
	var (
		assetClass AssetClass = "Equity"
		dataType   DataType   = "EndOfDay"
		id         int64      = 1
		sip        SIP        = "CTA"
		order      Order      = "asc"
		limit      int        = 20
		sort       Sort       = "name"
	)

	params := ListConditionsParams{}
	paramsWithAssetClass := params.WithAssetClass(assetClass)
	assert.NotNil(t, paramsWithAssetClass.AssetClass)
	assert.Equal(t, assetClass, *paramsWithAssetClass.AssetClass)

	paramsWithDataType := params.WithDataType(dataType)
	assert.NotNil(t, paramsWithDataType.DataType)
	assert.Equal(t, dataType, *paramsWithDataType.DataType)

	paramsWithID := params.WithID(id)
	assert.NotNil(t, paramsWithID.ID)
	assert.Equal(t, id, *paramsWithID.ID)

	paramsWithSIP := params.WithSIP(sip)
	assert.NotNil(t, paramsWithSIP.SIP)
	assert.Equal(t, sip, *paramsWithSIP.SIP)

	paramsWithOrder := params.WithOrder(order)
	assert.NotNil(t, paramsWithOrder.Order)
	assert.Equal(t, order, *paramsWithOrder.Order)

	paramsWithLimit := params.WithLimit(limit)
	assert.NotNil(t, paramsWithLimit.Limit)
	assert.Equal(t, limit, *paramsWithLimit.Limit)

	paramsWithSort := params.WithSort(sort)
	assert.NotNil(t, paramsWithSort.Sort)
	assert.Equal(t, sort, *paramsWithSort.Sort)
}
