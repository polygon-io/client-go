package reference_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListConditions(t *testing.T) {
	c := polygon.NewClient("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	condition := `{
	"asset_class": "stocks",
	"data_types": [
		"trade"
	],
	"id": 1,
	"legacy": false,
	"name": "Acquisition",
	"sip_mapping": {
		"UTP": "A"
	},
	"type": "sale_condition",
	"update_rules": {
		"consolidated": {
			"updates_high_low": true,
			"updates_open_close": true,
			"updates_volume": true
		},
		"market_center": {
			"updates_high_low": true,
			"updates_open_close": true,
			"updates_volume": true
		}
	}
}`

	expectedResponse := `{
	"status": "OK",
	"request_id": "4599a4e2ba5e19e2e732f711e97b0d84",
	"count": 1,
	"next_url": "https://api.polygon.io/v3/reference/conditions?cursor=YXA9MiZhcz0mYXNzZXRfY2xhc3M9c3RvY2tzJmRhdGFfdHlwZT10cmFkZSZsaW1pdD0yJnNvcnQ9YXNzZXRfY2xhc3M",
	"results": [
` + indent(true, condition, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/conditions?asset_class=stocks&data_type=trade&limit=1", expectedResponse)
	registerResponder("https://api.polygon.io/v3/reference/conditions?cursor=YXA9MiZhcz0mYXNzZXRfY2xhc3M9c3RvY2tzJmRhdGFfdHlwZT10cmFkZSZsaW1pdD0yJnNvcnQ9YXNzZXRfY2xhc3M", "{}")
	iter, err := c.Reference.ListConditions(context.Background(), models.ListConditionsParams{
		AssetClass: models.Ptr(models.AssetClassStocks),
		DataType:   models.Ptr(models.DataTypeTrade),
		Limit:      models.Ptr(1),
	})

	// iter creation
	assert.Nil(t, err)
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Condition())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err := json.MarshalIndent(iter.Condition(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, condition, string(b))

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}
