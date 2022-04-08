package reference

import (
	"context"
	"fmt"
	"net/http"

	"github.com/polygon-io/client-go/rest/models"
)

// ListSplitsIter is an iterator for the ListSplits method.
type ListSplitsIter struct {
	models.Iter
}

// Split returns the current result that the iterator points to.
func (it *ListSplitsIter) Split() models.Split {
	if it.Item() != nil {
		return it.Item().(models.Split)
	}
	return models.Split{}
}

// ListSplits retrieves reference splits.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_splits.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListSplits(context.TODO(), params, opts...)
//   if err != nil {
//       return err
//   }
//
//   for iter.Next() {
//       // do something with the current value
//       log.Print(iter.Split())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *Client) ListSplits(ctx context.Context, params *models.ListSplitsParams, options ...models.RequestOption) (*ListSplitsIter, error) {
	uri, err := c.EncodeParams(models.ListSplitsPath, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create iterator: %w", err)
	}

	return &ListSplitsIter{
		Iter: models.NewIter(ctx, uri, func(uri string) (models.ListResponse, []interface{}, error) {
			res := &models.ListSplitsResponse{}
			err := c.CallURL(ctx, http.MethodGet, uri, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}, nil
}
