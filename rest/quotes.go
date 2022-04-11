package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

// QuotesClient defines a REST client for the Polygon quotes API.
type QuotesClient struct {
	client.Client
}

// ListQuotesIter is an iterator for the ListQuotes method.
type ListQuotesIter struct {
	iter.Iter
}

// Quote returns the current result that the iterator points to.
func (it *ListQuotesIter) Quote() models.Quote {
	if it.Item() != nil {
		return it.Item().(models.Quote)
	}
	return models.Quote{}
}

// ListQuotes retrieves quotes for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v3_quotes__stockticker.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListQuotes(context.TODO(), params, opts...)
//   for iter.Next() {
//       // do something with the current value
//       log.Print(iter.Quote())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *QuotesClient) ListQuotes(ctx context.Context, params *models.ListQuotesParams, options ...models.RequestOption) *ListQuotesIter {
	return &ListQuotesIter{
		Iter: iter.NewIter(ctx, models.ListQuotesPath, params, func(uri string) (iter.ListResponse, []interface{}, error) {
			res := &models.ListQuotesResponse{}
			err := c.CallURL(ctx, http.MethodGet, uri, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}

// GetLastQuote retrieves the last quote (NBBO) for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v2_last_nbbo__stocksticker.
func (c *QuotesClient) GetLastQuote(ctx context.Context, params *models.GetLastQuoteParams, options ...models.RequestOption) (*models.GetLastQuoteResponse, error) {
	res := &models.GetLastQuoteResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetLastQuotePath, params, res, options...)
	return res, err
}

// GetLastForexQuote retrieves the last quote (BBO) for a forex currency pair.
// For more details see https://polygon.io/docs/forex/get_v1_last_quote_currencies__from___to.
func (c *QuotesClient) GetLastForexQuote(ctx context.Context, params *models.GetLastForexQuoteParams, options ...models.RequestOption) (*models.GetLastForexQuoteResponse, error) {
	res := &models.GetLastForexQuoteResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetLastForexQuotePath, params, res, options...)
	return res, err
}
