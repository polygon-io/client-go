package trades

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

// Client defines a REST client for the Polygon trades API.
type Client struct {
	client.Client
}

// TradesIter defines a domain specific iterator for the trades API.
type TradesIter struct {
	client.Iter
}

// Trade returns the current result that the iterator points to.
func (it *TradesIter) Trade() *models.Trade {
	if it.Item() != nil {
		return it.Item().(*models.Trade)
	}
	return nil
}

// ListTrades retrieves trades for a specified ticker. This method returns an iterator that should be used to
// access the results via this pattern:
//   iter := c.ListTrades(context.TODO(), params, opts...)
//   for iter.Next() {
//       // Do something with the current value
//       log.Print(iter.Trade())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *Client) ListTrades(ctx context.Context, params models.ListTradesParams, options ...client.Option) *TradesIter {
	return &TradesIter{
		Iter: client.GetIter(ctx, params.String(), func(url string) (client.ListResponse, []interface{}, error) {
			res := &models.TradesResponse{}
			err := c.Call(ctx, http.MethodGet, url, nil, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}

// GetLastTrade retrieves the last trade for a specified ticker.
func (c *Client) GetLastTrade(ctx context.Context, params models.GetLastTradeParams, options ...client.Option) (*models.LastTradeResponse, error) {
	res := &models.LastTradeResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetLastTradePath, params, res, options...)
	return res, err
}

// GetLastCryptoTrade retrieves the last trade for a crypto pair.
func (c *Client) GetLastCryptoTrade(ctx context.Context, params models.LastCryptoTradeParams, options ...client.Option) (*models.LastCryptoTradeResponse, error) {
	res := &models.LastCryptoTradeResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetLastCryptoTradePath, params, res, options...)
	return res, err
}
