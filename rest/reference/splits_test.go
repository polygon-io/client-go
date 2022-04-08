package reference_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

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

	registerResponder("https://api.polygon.io/v3/reference/splits?execution_date=2021-07-22&limit=2&order=asc&reverse_split=false&sort=ticker&ticker=AAPL", expectedResponse)
	iter := c.Reference.ListSplits(context.Background(), models.ListSplitsParams{}.
		WithTicker(models.EQ, "AAPL").WithExecutionDate(models.EQ, models.Date(time.Date(2021, 7, 22, 0, 0, 0, 0, time.UTC))).WithReverseSplit(false).
		WithSort(models.TickerSymbol).WithOrder(models.Asc).WithLimit(2))

	// iter creation
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
