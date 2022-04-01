package reference

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

// Client defines a REST client for the Polygon reference API.
type Client struct {
	client.Client
}

// ListTickersIter is an iterator for the ListTickers method.
type ListTickersIter struct {
	client.Iter
}

// Ticker returns the current result that the iterator points to.
func (it *ListTickersIter) Ticker() models.Ticker {
	if it.Item() != nil {
		return it.Item().(models.Ticker)
	}
	return models.Ticker{}
}

// ListTickers retrieves reference tickers.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListTickers(context.TODO(), params, opts...)
//   if err != nil {
//       return err
//   }
//
//   for iter.Next() {
//       // do something with the current value
//       log.Print(iter.Ticker())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *Client) ListTickers(ctx context.Context, params models.ListTickersParams, options ...models.RequestOption) (*ListTickersIter, error) {
	url, err := c.EncodeParams(models.ListTickersPath, params)
	if err != nil {
		return nil, err
	}

	return &ListTickersIter{
		Iter: client.NewIter(ctx, url, func(url string) (models.ListResponse, []interface{}, error) {
			res := &models.ListTickersResponse{}
			err := c.Call(ctx, http.MethodGet, url, nil, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}, nil
}

// GetTickerDetails retrieves details for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker.
func (c *Client) GetTickerDetails(ctx context.Context, params models.GetTickerDetailsParams, options ...models.RequestOption) (*models.GetTickerDetailsResponse, error) {
	res := &models.GetTickerDetailsResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetTickerDetailsPath, params, res, options...)
	return res, err
}

// GetTickerTypes retrieves all the possible ticker types that can be queried.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers_types.
func (c *Client) GetTickerTypes(ctx context.Context, params models.GetTickerTypesParams, options ...models.RequestOption) (*models.GetTickerTypesResponse, error) {
	res := &models.GetTickerTypesResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetTickerTypesPath, params, res, options...)
	return res, err
}
