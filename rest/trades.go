package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	ListTradesPath         = "/v3/trades/{ticker}"
	GetLastTradePath       = "/v2/last/trade/{ticker}"
	GetLastCryptoTradePath = "/v1/last/crypto/{from}/{to}"
)

// TradesClient defines a REST client for the Polygon trades API.
type TradesClient struct {
	client.Client
}

// ListTrades retrieves trades for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v3_trades__stockticker.
// This method returns an iterator that should be used to access the results via this pattern:
//   iter, err := c.ListTrades(context.TODO(), params, opts...)
//   for iter.Next() {
//       log.Print(iter.Item()) // do something with the current value
//   }
//   if iter.Err() != nil {
//       return err
//   }
func (c *TradesClient) ListTrades(ctx context.Context, params *models.ListTradesParams, options ...models.RequestOption) *iter.Iter[models.Trade] {
	return iter.NewIter(ctx, ListTradesPath, params, func(uri string) (iter.ListResponse, []models.Trade, error) {
		res := &models.ListTradesResponse{}
		err := c.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetLastTrade retrieves the last trade for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v2_last_trade__stocksticker.
func (c *TradesClient) GetLastTrade(ctx context.Context, params *models.GetLastTradeParams, options ...models.RequestOption) (*models.GetLastTradeResponse, error) {
	res := &models.GetLastTradeResponse{}
	err := c.Call(ctx, http.MethodGet, GetLastTradePath, params, res, options...)
	return res, err
}

// GetLastCryptoTrade retrieves the last trade for a crypto pair.
// For more details see https://polygon.io/docs/crypto/get_v1_last_crypto__from___to.
func (c *TradesClient) GetLastCryptoTrade(ctx context.Context, params *models.GetLastCryptoTradeParams, options ...models.RequestOption) (*models.GetLastCryptoTradeResponse, error) {
	res := &models.GetLastCryptoTradeResponse{}
	err := c.Call(ctx, http.MethodGet, GetLastCryptoTradePath, params, res, options...)
	return res, err
}
