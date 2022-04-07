package reference

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

// ListSplitsIter is an iterator for the ListSplits method.
type ListSplitsIter struct {
	client.Iter
}

// Split returns the current result that the iterator points to.
func (it *ListSplitsIter) Split() models.Split {
	if it.Item() != nil {
		return it.Item().(models.Split)
	}
	return models.Split{}
}

// ListSplits retrieves reference splits.
func (c *Client) ListSplits(ctx context.Context, params models.ListSplitsParams) (*ListSplitsIter, error) {
	url, err := c.EncodeParams(models.ListSplitsPath, params)
	if err != nil {
		return nil, err
	}

	return &ListSplitsIter{
		Iter: client.NewIter(ctx, url, func(url string) (models.ListResponse, []interface{}, error) {
			res := &models.ListSplitsResponse{}
			err := c.Call(ctx, http.MethodGet, url, nil, res)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}, nil
}
