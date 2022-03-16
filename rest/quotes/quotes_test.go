package quotes_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/stretchr/testify/assert"
)

func TestListQuotes(t *testing.T) {
	c := polygon.New("API_KEY")

	httpmock.ActivateNonDefault(c.Quotes.HTTP.GetClient())
	defer httpmock.DeactivateAndReset()

	quote1 := models.Quote{AskPrice: 1.23}
	quote2 := models.Quote{AskPrice: 1.5}
	expectedResponse := models.QuotesResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req1",
			Count:     2,
			PaginationHooks: client.PaginationHooks{
				NextURL: "https://api.polygon.io/v3/quotes/AAPL?cursor=YXA9OT",
			},
		},
		Results: []*models.Quote{&quote1, &quote2},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/quotes/AAPL?limit=2&order=asc&sort=timestamp&timestamp.lte=1626912000000000000",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	quote3 := models.Quote{AskPrice: 1.40}
	expectedNextResponse := models.QuotesResponse{
		BaseResponse: client.BaseResponse{
			Status:    "OK",
			RequestID: "req2",
			Count:     1,
		},
		Results: []*models.Quote{&quote3},
	}

	httpmock.RegisterResponder("GET", "https://api.polygon.io/v3/quotes/AAPL?cursor=YXA9OT",
		func(req *http.Request) (*http.Response, error) {
			b, err := json.Marshal(expectedNextResponse)
			assert.Nil(t, err)
			resp := httpmock.NewStringResponse(200, string(b))
			resp.Header.Add("Content-Type", "application/json")
			return resp, nil
		},
	)

	iter := c.Quotes.ListQuotes(context.Background(), models.ListQuotesParams{
		Ticker: "AAPL",
		QueryParams: models.ListQuotesQueryParams{
			TimestampLTE: models.String("1626912000000000000"),
			Order:        models.OrderBy(models.Asc),
			Limit:        models.Int(2),
			Sort:         models.SortOn(models.Timestamp),
		},
	})

	// verify the first page
	assert.Nil(t, iter.Err())
	assert.Equal(t, &expectedResponse, iter.Page())
	assert.Nil(t, iter.Quote())
	// verify the first and second quotes
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, &quote1, iter.Quote())
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	assert.Equal(t, &quote2, iter.Quote())

	// verify the second page
	assert.True(t, iter.Next())
	assert.Nil(t, iter.Err())
	// verify the third quote
	assert.Equal(t, &expectedNextResponse, iter.QuotesList())
	assert.Equal(t, &quote3, iter.Quote())

	// verify the end of the list
	assert.False(t, iter.Next())
	assert.Nil(t, iter.Err())
}
