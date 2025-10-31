package models_test

import (
	"testing"

	"github.com/massive-com/client-go/v2/rest/models"
)

func TestListConditionsParams(t *testing.T) {
	assetClass := models.AssetStocks
	dataType := models.DataTrade
	id := int64(1)
	sip := models.CTA
	order := models.Asc
	limit := 20
	sort := models.Name

	expect := models.ListConditionsParams{
		AssetClass: &assetClass,
		DataType:   &dataType,
		ID:         &id,
		SIP:        &sip,
		Order:      &order,
		Limit:      &limit,
		Sort:       &sort,
	}
	actual := models.ListConditionsParams{}.
		WithAssetClass(assetClass).
		WithDataType(dataType).
		WithID(id).
		WithSIP(sip).
		WithOrder(order).
		WithLimit(limit).
		WithSort(sort)
	checkParams(t, expect, *actual)
}
