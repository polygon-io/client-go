package aggs

import (
	"context"
	"net/http"

	"github.com/polygon-io/client-golang/rest/client"
)

// todo: add comments for godoc

type Client struct {
	client.BaseClient
}

func (ac *Client) Get(ctx context.Context, params GetParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	path := "/v2/aggs/ticker/{ticker}/range/{multiplier}/{resolution}/{from}/{to}"
	err := ac.Call(http.MethodGet, path, params, res, append([]client.Option{client.WithContext(ctx)}, opts...)...)
	return res, err
}

func (ac *Client) GetPreviousClose(ctx context.Context, params GetPreviousCloseParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	path := "/v2/aggs/ticker/{ticker}/prev"
	err := ac.Call(http.MethodGet, path, params, res, append([]client.Option{client.WithContext(ctx)}, opts...)...)
	return res, err
}

func (ac *Client) GetGroupedDaily(ctx context.Context, params GetGroupedDailyParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	path := "/v2/aggs/grouped/locale/{locale}/market/{marketType}/{date}"
	err := ac.Call(http.MethodGet, path, params, res, append([]client.Option{client.WithContext(ctx)}, opts...)...)
	return res, err
}

func (ac *Client) GetDailyOpenClose(ctx context.Context, params GetDailyOpenCloseParams, opts ...client.Option) (*DailyOpenCloseResponse, error) {
	res := &DailyOpenCloseResponse{}
	path := "/v1/open-close/{ticker}/{date}"
	err := ac.Call(http.MethodGet, path, params, res, append([]client.Option{client.WithContext(ctx)}, opts...)...)
	return res, err
}
