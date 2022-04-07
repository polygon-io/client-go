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

// ListTradesIter is an iterator for the ListTickers method.
type ListTradesIter struct {
	client.Iter
}

// Trade returns the current result that the iterator points to.
func (it *ListTradesIter) Trade() models.Trade {
	if it.Item() != nil {
		return it.Item().(models.Trade)
	}
	return models.Trade{}
}

// ListTrades retrieves trades for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v3_trades__stockticker.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListTrades(context.TODO(), params, opts...)
//   if err != nil {
//       return err
//   }
//
//   for iter.Next() {
//       // do something with the current value
//       log.Print(iter.Trade())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *Client) ListTrades(ctx context.Context, params *models.ListTradesParams, options ...models.RequestOption) (*ListTradesIter, error) {
	uri, err := c.EncodeParams(models.ListTradesPath, params)
	if err != nil {
		return nil, err
	}

	return &ListTradesIter{
		Iter: client.NewIter(ctx, uri, func(uri string) (models.ListResponse, []interface{}, error) {
			res := &models.ListTradesResponse{}
			err := c.CallURL(ctx, http.MethodGet, uri, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}, nil
}

// GetLastTrade retrieves the last trade for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v2_last_trade__stocksticker.
func (c *Client) GetLastTrade(ctx context.Context, params *models.GetLastTradeParams, options ...models.RequestOption) (*models.GetLastTradeResponse, error) {
	res := &models.GetLastTradeResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetLastTradePath, params, res, options...)
	return res, err
}

// GetLastCryptoTrade retrieves the last trade for a crypto pair.
// For more details see https://polygon.io/docs/crypto/get_v1_last_crypto__from___to.
func (c *Client) GetLastCryptoTrade(ctx context.Context, params *models.GetLastCryptoTradeParams, options ...models.RequestOption) (*models.GetLastCryptoTradeResponse, error) {
	res := &models.GetLastCryptoTradeResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetLastCryptoTradePath, params, res, options...)
	return res, err
}
