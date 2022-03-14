package aggs

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-golang/rest/client"
)

const (
	getPath               = "/v2/aggs/ticker/{ticker}/range/{multiplier}/{resolution}/{from}/{to}"
	getPreviousClosePath  = "/v2/aggs/ticker/{ticker}/prev"
	getGroupedDailyPath   = "/v2/aggs/grouped/locale/{locale}/market/{marketType}/{date}"
	getDailyOpenClosePath = "/v1/open-close/{ticker}/{date}"
)

// Client defines a REST client for the Polygon aggregates API.
type Client struct {
	client.BaseClient
}

// Get retrieves aggregate bars for a specified ticker over a given date range in custom time window sizes.
// For example, if timespan = ‘minute’ and multiplier = ‘5’ then 5-minute bars will be returned.
func (ac *Client) Get(ctx context.Context, params GetParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	err := ac.Call(ctx, http.MethodGet, getPath, params, res, opts...)
	return res, err
}

// GetPreviousClose retrieves the previous day's open, high, low, and close (OHLC) for the specified ticker.
func (ac *Client) GetPreviousClose(ctx context.Context, params GetPreviousCloseParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	err := ac.Call(ctx, http.MethodGet, getPreviousClosePath, params, res, opts...)
	return res, err
}

// GetGroupedDaily retrieves the daily open, high, low, and close (OHLC) for the specified market type.
func (ac *Client) GetGroupedDaily(ctx context.Context, params GetGroupedDailyParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	err := ac.Call(ctx, http.MethodGet, getGroupedDailyPath, params, res, opts...)
	return res, err
}

// GetDailyOpenClose retrieves the open, close and afterhours prices of a specific symbol on a certain date.
func (ac *Client) GetDailyOpenClose(ctx context.Context, params GetDailyOpenCloseParams, opts ...client.Option) (*DailyOpenCloseResponse, error) {
	res := &DailyOpenCloseResponse{}
	err := ac.Call(ctx, http.MethodGet, getDailyOpenClosePath, params, res, opts...)
	return res, err
}
