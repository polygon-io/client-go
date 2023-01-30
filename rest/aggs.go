package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
)

const (
	ListAggsPath             = "/v2/aggs/ticker/{ticker}/range/{multiplier}/{timespan}/{from}/{to}"
	GetAggsPath              = "/v2/aggs/ticker/{ticker}/range/{multiplier}/{timespan}/{from}/{to}"
	GetGroupedDailyAggsPath  = "/v2/aggs/grouped/locale/{locale}/market/{marketType}/{date}"
	GetDailyOpenCloseAggPath = "/v1/open-close/{ticker}/{date}"
	GetPreviousCloseAggPath  = "/v2/aggs/ticker/{ticker}/prev"
)

// AggsClient defines a REST client for the Polygon aggs API.
type AggsClient struct {
	client.Client
}

// ListAggs retrieves aggregate bars for a specified ticker over a given date range in custom time window sizes.
// For example, if timespan = 'minute' and multiplier = '5' then 5-minute bars will be returned.
// For more details see https://polygon.io/docs/stocks/get_v2_aggs_ticker__stocksticker__range__multiplier___timespan___from___to.
//
// This method returns an iterator that should be used to access the results via this pattern:
//
//	iter := c.ListAggs(context.TODO(), params, opts...)
//	for iter.Next() {
//		log.Print(iter.Item()) // do something with the current value
//	}
//	if iter.Err() != nil {
//		return iter.Err()
//	}
func (ac *AggsClient) ListAggs(ctx context.Context, params *models.ListAggsParams, options ...models.RequestOption) *iter.Iter[models.Agg] {
	return iter.NewIter(ctx, ListAggsPath, params, func(uri string) (iter.ListResponse, []models.Agg, error) {
		res := &models.ListAggsResponse{}
		err := ac.CallURL(ctx, http.MethodGet, uri, res, options...)
		return res, res.Results, err
	})
}

// GetAggs retrieves aggregate bars for a specified ticker over a given date range in custom time window sizes.
// For example, if timespan = 'minute' and multiplier = '5' then 5-minute bars will be returned.
// For more details see https://polygon.io/docs/stocks/get_v2_aggs_ticker__stocksticker__range__multiplier___timespan___from___to.
//
// Deprecated: This method does not return an iterator and forces users to handle pagination manually. Use
// pkg.go.dev/github.com/polygon-io/client-go/rest#AggsClient.ListAggs instead if you want automatic pagination.
func (ac *AggsClient) GetAggs(ctx context.Context, params *models.GetAggsParams, opts ...models.RequestOption) (*models.GetAggsResponse, error) {
	res := &models.GetAggsResponse{}
	err := ac.Call(ctx, http.MethodGet, GetAggsPath, params, res, opts...)
	return res, err
}

// GetGroupedDailyAggs retrieves the daily open, high, low, and close (OHLC) for the specified market type.
// For more details see https://polygon.io/docs/stocks/get_v2_aggs_grouped_locale_us_market_stocks__date.
func (ac *AggsClient) GetGroupedDailyAggs(ctx context.Context, params *models.GetGroupedDailyAggsParams, opts ...models.RequestOption) (*models.GetGroupedDailyAggsResponse, error) {
	res := &models.GetGroupedDailyAggsResponse{}
	err := ac.Call(ctx, http.MethodGet, GetGroupedDailyAggsPath, params, res, opts...)
	return res, err
}

// GetDailyOpenCloseAgg retrieves the open, close and afterhours prices of a specific symbol on a certain date.
// For more details see https://polygon.io/docs/stocks/get_v1_open-close__stocksticker___date.
func (ac *AggsClient) GetDailyOpenCloseAgg(ctx context.Context, params *models.GetDailyOpenCloseAggParams, opts ...models.RequestOption) (*models.GetDailyOpenCloseAggResponse, error) {
	res := &models.GetDailyOpenCloseAggResponse{}
	err := ac.Call(ctx, http.MethodGet, GetDailyOpenCloseAggPath, params, res, opts...)
	return res, err
}

// GetPreviousCloseAgg retrieves the previous day's open, high, low, and close (OHLC) for the specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v2_aggs_ticker__stocksticker__prev.
func (ac *AggsClient) GetPreviousCloseAgg(ctx context.Context, params *models.GetPreviousCloseAggParams, opts ...models.RequestOption) (*models.GetPreviousCloseAggResponse, error) {
	res := &models.GetPreviousCloseAggResponse{}
	err := ac.Call(ctx, http.MethodGet, GetPreviousCloseAggPath, params, res, opts...)
	return res, err
}
