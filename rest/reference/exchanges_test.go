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

func TestListExchanges(t *testing.T) {
	c := polygon.NewClient("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `{
	"status": "OK",
	"request_id": "c784b78622b5a68c932af78a68b5907c",
	"count": 1,
	"results": [
		{
			"acronym": "AMEX",
			"asset_class": "stocks",
			"id": 1,
			"locale": "us",
			"mic": "XASE",
			"name": "NYSE American, LLC",
			"operating_mic": "XNYS",
			"participant_id": "A",
			"type": "exchange",
			"url": "https://www.nyse.com/markets/nyse-american"
		}
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/exchanges?asset_class=stocks&locale=us", expectedResponse)
	res, err := c.Reference.GetExchanges(context.Background(), models.GetExchangesParams{}.WithAssetClass(models.AssetStocks).WithLocale(models.US))

	assert.Nil(t, err)
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}
