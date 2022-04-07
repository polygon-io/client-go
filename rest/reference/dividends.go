package reference

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

// ListDividendsIter is an iterator for the ListDividends method.
type ListDividendsIter struct {
	client.Iter
}

// Dividend returns the current result that the iterator points to.
func (it *ListDividendsIter) Dividend() models.Dividend {
	if it.Item() != nil {
		return it.Item().(models.Dividend)
	}
	return models.Dividend{}
}

// ListDividends retrieves reference dividends.
func (c *Client) ListDividends(ctx context.Context, params models.ListDividendsParams, options ...models.RequestOption) (*ListDividendsIter, error) {
	url, err := c.EncodeParams(models.ListDividendsPath, params)
	if err != nil {
		return nil, err
	}

	return &ListDividendsIter{
		Iter: client.NewIter(ctx, url, func(url string) (models.ListResponse, []interface{}, error) {
			res := &models.ListDividendsResponse{}
			err := c.Call(ctx, http.MethodGet, url, nil, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}, nil
}
