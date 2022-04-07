package reference

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

// ListExchangesIter is an iterator for the ListExchanges method.
type ListExchangesIter struct {
	client.Iter
}

// Exchange returns the current result that the iterator points to.
func (it *ListExchangesIter) Exchange() models.Exchange {
	if it.Item() != nil {
		return it.Item().(models.Exchange)
	}
	return models.Exchange{}
}

// ListExchanges retrieves reference exchanges.
func (c *Client) ListExchanges(ctx context.Context, params models.ListExchangesParams, options ...models.RequestOption) (*ListExchangesIter, error) {
	url, err := c.EncodeParams(models.ListExchangesPath, params)
	if err != nil {
		return nil, err
	}

	return &ListExchangesIter{
		Iter: client.NewIter(ctx, url, func(url string) (models.ListResponse, []interface{}, error) {
			res := &models.ListExchangesResponse{}
			err := c.Call(ctx, http.MethodGet, url, nil, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}, nil
}
