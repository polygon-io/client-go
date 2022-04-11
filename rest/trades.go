package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

// TradesClient defines a REST client for the Polygon trades API.
type TradesClient struct {
	client.Client
}

// ListTradesIter is an iterator for the ListTickers method.
type ListTradesIter struct {
	iter.Iter
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
//   for iter.Next() {
//       // do something with the current value
//       log.Print(iter.Trade())
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *TradesClient) ListTrades(ctx context.Context, params *models.ListTradesParams, options ...models.RequestOption) *ListTradesIter {
	return &ListTradesIter{
		Iter: iter.NewIter(ctx, models.ListTradesPath, params, func(uri string) (iter.ListResponse, []interface{}, error) {
			res := &models.ListTradesResponse{}
			err := c.CallURL(ctx, http.MethodGet, uri, res, options...)

			results := make([]interface{}, len(res.Results))
			for i, v := range res.Results {
				results[i] = v
			}

			return res, results, err
		}),
	}
}

// GetLastTrade retrieves the last trade for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v2_last_trade__stocksticker.
func (c *TradesClient) GetLastTrade(ctx context.Context, params *models.GetLastTradeParams, options ...models.RequestOption) (*models.GetLastTradeResponse, error) {
	res := &models.GetLastTradeResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetLastTradePath, params, res, options...)
	return res, err
}

// GetLastCryptoTrade retrieves the last trade for a crypto pair.
// For more details see https://polygon.io/docs/crypto/get_v1_last_crypto__from___to.
func (c *TradesClient) GetLastCryptoTrade(ctx context.Context, params *models.GetLastCryptoTradeParams, options ...models.RequestOption) (*models.GetLastCryptoTradeResponse, error) {
	res := &models.GetLastCryptoTradeResponse{}
	err := c.Call(ctx, http.MethodGet, models.GetLastCryptoTradePath, params, res, options...)
	return res, err
}
