package reference_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListExchanges(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	exchange1 := `{
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
}`

	expectedResponse := `{
	"status": "OK",
	"request_id": "c784b78622b5a68c932af78a68b5907c",
	"count": 1,
	"results": [
` + indent(true, exchange1, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/exchanges?asset_class=stocks&locale=us", expectedResponse)
	iter, err := c.Reference.ListExchanges(context.Background(), models.ListExchangesParams{
		AssetClass: models.Ptr("stocks"),
		Locale:     models.Ptr("us"),
	})

	// iter creation
	assert.Nil(t, err)
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Exchange())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err := json.MarshalIndent(iter.Exchange(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, exchange1, string(b))

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}
