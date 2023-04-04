package models_test

import (
	"testing"

	"github.com/polygon-io/client-go/rest/models"
)

func TestListConditionsParams(t *testing.T) {
	assetClass := models.AssetClass("Equity")
	dataType := models.DataType("EndOfDay")
	id := int64(1)
	sip := models.SIP("CTA")
	order := models.Order("asc")
	limit := 20
	sort := models.Sort("name")

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
