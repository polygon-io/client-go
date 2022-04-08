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

func TestListDividends(t *testing.T) {
	c := polygon.NewClient("API_KEY")

	httpmock.ActivateNonDefault(c.Reference.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	dividend1 := `{
	"cash_amount": 0.59375,
	"declaration_date": "2020-09-09",
	"dividend_type": "CD",
	"ex_dividend_date": "2025-06-12",
	"frequency": 4,
	"pay_date": "2025-06-30",
	"record_date": "2025-06-15",
	"ticker": "CSSEN"
}`

	expectedResponse := `{
	"status": "OK",
	"request_id": "eca6d9a0d8dc1cd1b29d2d3112fe938e",
	"next_url": "https://api.polygon.io/v3/reference/dividends?cursor=YXA9MjUmYXM9JmxpbWl0PTEwJm9yZGVyPWRlc2Mmc29ydD1leF9kaXZpZGVuZF9kYXRlJnRpY2tlcj1DU1NFTg",
	"results": [
` + indent(true, dividend1, "\t\t") + `
	]
}`

	registerResponder("https://api.polygon.io/v3/reference/dividends?dividend_type=CD&ticker=CSSEN", expectedResponse)
	registerResponder("https://api.polygon.io/v3/reference/dividends?cursor=YXA9MjUmYXM9JmxpbWl0PTEwJm9yZGVyPWRlc2Mmc29ydD1leF9kaXZpZGVuZF9kYXRlJnRpY2tlcj1DU1NFTg", "{}")
	iter, err := c.Reference.ListDividends(context.Background(), models.ListDividendsParams{
		TickerEQ:     models.Ptr("CSSEN"),
		DividendType: models.Ptr(models.DividendTypeCD),
	})

	// iter creation
	assert.Nil(t, err)
	assert.Nil(t, iter.Err())
	assert.NotNil(t, iter.Dividend())

	// first item
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	b, err := json.MarshalIndent(iter.Dividend(), "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, dividend1, string(b))

	// end of list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}
