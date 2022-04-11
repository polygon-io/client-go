package polygon

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-go/rest/client"
	"github.com/polygon-io/client-go/rest/models"
)

// AggsClient defines a REST client for the Polygon aggs API.
type AggsClient struct {
	client.Client
}

// GetAggs retrieves aggregate bars for a specified ticker over a given date range in custom time window sizes.
// For example, if timespan = ‘minute’ and multiplier = ‘5’ then 5-minute bars will be returned.
// For more details see https://polygon.io/docs/stocks/get_v2_aggs_ticker__stocksticker__range__multiplier___timespan___from___to.
func (ac *AggsClient) GetAggs(ctx context.Context, params *models.GetAggsParams, opts ...models.RequestOption) (*models.GetAggsResponse, error) {
	res := &models.GetAggsResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetAggsPath, params, res, opts...)
	return res, err
}

// GetGroupedDailyAggs retrieves the daily open, high, low, and close (OHLC) for the specified market type.
// For more details see https://polygon.io/docs/stocks/get_v2_aggs_grouped_locale_us_market_stocks__date.
func (ac *AggsClient) GetGroupedDailyAggs(ctx context.Context, params *models.GetGroupedDailyAggsParams, opts ...models.RequestOption) (*models.GetGroupedDailyAggsResponse, error) {
	res := &models.GetGroupedDailyAggsResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetGroupedDailyAggsPath, params, res, opts...)
	return res, err
}

// GetDailyOpenClose retrieves the open, close and afterhours prices of a specific symbol on a certain date.
// For more details see https://polygon.io/docs/stocks/get_v1_open-close__stocksticker___date.
func (ac *AggsClient) GetDailyOpenCloseAgg(ctx context.Context, params *models.GetDailyOpenCloseAggParams, opts ...models.RequestOption) (*models.GetDailyOpenCloseAggResponse, error) {
	res := &models.GetDailyOpenCloseAggResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetDailyOpenCloseAggPath, params, res, opts...)
	return res, err
}

// GetPreviousClose retrieves the previous day's open, high, low, and close (OHLC) for the specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v2_aggs_ticker__stocksticker__prev.
func (ac *AggsClient) GetPreviousCloseAgg(ctx context.Context, params *models.GetPreviousCloseAggParams, opts ...models.RequestOption) (*models.GetPreviousCloseAggResponse, error) {
	res := &models.GetPreviousCloseAggResponse{}
	err := ac.Call(ctx, http.MethodGet, models.GetPreviousCloseAggPath, params, res, opts...)
	return res, err
}
