package aggregates

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-golang/rest/client"
)

// todo: add comments for godoc

const (
	GetPath               = "/v2/aggs/ticker/{ticker}/range/{multiplier}/{resolution}/{from}/{to}"
	GetPreviousClosePath  = "/v2/aggs/ticker/{ticker}/prev"
	GetGroupedDailyPath   = "/v2/aggs/grouped/locale/{locale}/market/{marketType}/{date}"
	GetDailyOpenClosePath = "/v1/open-close/{ticker}/{date}"
)

type Client struct {
	client.HTTPBase
}

func (ac *Client) Get(ctx context.Context, params GetParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	err := ac.Call(http.MethodGet, GetPath, params, res, append([]client.Option{client.WithContext(ctx)}, opts...)...)
	return res, err
}

func (ac *Client) GetPreviousClose(ctx context.Context, params GetPreviousCloseParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	err := ac.Call(http.MethodGet, GetPreviousClosePath, params, res, append([]client.Option{client.WithContext(ctx)}, opts...)...)
	return res, err
}

func (ac *Client) GetGroupedDaily(ctx context.Context, params GetGroupedDailyParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	err := ac.Call(http.MethodGet, GetGroupedDailyPath, params, res, append([]client.Option{client.WithContext(ctx)}, opts...)...)
	return res, err
}

func (ac *Client) GetDailyOpenClose(ctx context.Context, params GetDailyOpenCloseParams, opts ...client.Option) (*DailyOpenCloseResponse, error) {
	res := &DailyOpenCloseResponse{}
	err := ac.Call(http.MethodGet, GetDailyOpenClosePath, params, res, append([]client.Option{client.WithContext(ctx)}, opts...)...)
	return res, err
}
