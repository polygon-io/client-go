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

func TestListSplits(t *testing.T) {
	c := polygon.NewClient("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	split1 := `{
	"execution_date": "2020-08-31",
	"split_from": 1,
	"split_to": 4,
	"ticker": "AAPL"
}`

	expectedResponse := `{
	"status": "OK",
	"request_id": "2b539ae65c1478dee109b7397bd591b2",
	"results": [
` + indent(true, split1, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/splits?execution_date=2020-08-31&limit=2&order=asc&reverse_split=false&sort=ticker&ticker=AAPL", expectedResponse)
	iter, err := c.Reference.ListSplits(context.Background(), models.ListSplitsParams{
		TickerEQ:        models.Ptr("AAPL"),
		ExecutionDateEQ: models.Ptr("2020-08-31"),
		ReverseSplit:    models.Ptr(false),
		Sort:            models.Ptr(models.TickerSymbol),
		Order:           models.Ptr(models.Asc),
		Limit:           models.Ptr(2),
	})

	// iter creation
	assert.Nil(t, err)
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Split())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err := json.MarshalIndent(iter.Split(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, split1, string(b))

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}
