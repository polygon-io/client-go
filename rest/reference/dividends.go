package reference

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

// ListDividendsIter is an iterator for the ListDividends method.
type ListDividendsIter struct {
	iter.Iter
}

// Dividend returns the current result that the iterator points to.
func (it *ListDividendsIter) Dividend() models.Dividend {
	if it.Item() != nil {
		return it.Item().(models.Dividend)
	}
	return models.Dividend{}
}

// ListDividends retrieves reference dividends.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_dividends.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListDividends(context.TODO(), params, opts...)
//   for iter.Next() {
//       // do something with the current value
//       log.Print(iter.Dividend())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *Client) ListDividends(ctx context.Context, params *models.ListDividendsParams, options ...models.RequestOption) *ListDividendsIter {
	return &ListDividendsIter{
		Iter: iter.NewIter(ctx, models.ListDividendsPath, params, func(uri string) (iter.ListResponse, []interface{}, error) {
			res := &models.ListDividendsResponse{}
			err := c.CallURL(ctx, http.MethodGet, uri, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}
