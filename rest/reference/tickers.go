package reference

import (
	"context"
	"fmt"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

type Client struct {
	client.Client
}

// TickersIter defines a domain specific iterator for the reference tickers API.
type TickersIter struct {
	client.Iter
}

// Ticker returns the current result that the iterator points to.
func (it *TickersIter) Ticker() *models.TickerDetails {
	if it.Item() != nil {
		return it.Item().(*models.TickerDetails)
	}
	return nil
}

// ListTickers retrieves reference tickers. This method returns an iterator that should be used to
// access the results via this pattern:
//   iter, err := c.ListTickers(context.TODO(), params, opts...)
//   if err != nil {
//       return err
//   }
//
//   for iter.Next() {
//       // Do something with the current value
//       log.Print(iter.Ticker())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *Client) ListTickers(ctx context.Context, params models.ListTickersParams, options ...client.Option) (*TickersIter, error) {
	iter, err := c.NewIter(ctx, models.ListTickersPath, params, func(url string) (client.ListResponse, []interface{}, error) {
		res := &models.TickersResponse{}
		err := c.Call(ctx, http.MethodGet, url, nil, res, options...)

		results := make([]interface{}, len(res.Results))
		for i, v := range res.Results {
			results[i] = v
		}

		return res, results, err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create iterator: %w", err)
	}

	return &TickersIter{
		Iter: *iter,
	}, nil
}

// GetTickerDetails retrieves details for a specified ticker.
func (c *Client) GetTickerDetails(ctx context.Context, params models.GetTickerDetailsParams, options ...client.Option) (*models.TickersResponse, error) {
	res := &models.TickersResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetTickerDetailsPath, params, res, options...)
	return res, err
}

// GetTickerTypes retrieves all the possible ticker types that can be queried.
func (c *Client) GetTickerTypes(ctx context.Context, params models.GetTickerTypesParams, options ...client.Option) (*models.TickerTypesResponse, error) {
	res := &models.TickerTypesResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetTickerTypesPath, params, res, options...)
	return res, err
}
