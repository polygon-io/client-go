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

func TestGetMarketHolidays(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	expectedResponse := `[
	{
		"exchange": "NYSE",
		"name": "Thanksgiving",
		"date": "2020-11-26T00:00:00.000Z",
		"status": "closed"
	},
	{
		"exchange": "NASDAQ",
		"name": "Thanksgiving",
		"date": "2020-11-26T00:00:00.000Z",
		"status": "closed"
	},
	{
		"exchange": "OTC",
		"name": "Thanksgiving",
		"date": "2020-11-26T00:00:00.000Z",
		"status": "closed"
	},
	{
		"exchange": "NASDAQ",
		"name": "Thanksgiving",
		"date": "2020-11-27T00:00:00.000Z",
		"status": "early-close",
		"open": "2020-11-27T14:30:00.000Z",
		"close": "2020-11-27T18:00:00.000Z"
	},
	{
		"exchange": "NYSE",
		"name": "Thanksgiving",
		"date": "2020-11-27T00:00:00.000Z",
		"status": "early-close",
		"open": "2020-11-27T14:30:00.000Z",
		"close": "2020-11-27T18:00:00.000Z"
	}
]`

	registerResponder("https://api.polygon.io/v1/marketstatus/upcoming", expectedResponse)
	res, err := c.Reference.GetMarketHolidays(context.Background(), models.GetMarketHolidaysParams{})

	assert.Nil(t, err)
	b, err := json.MarshalIndent(res, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, string(b))
}
