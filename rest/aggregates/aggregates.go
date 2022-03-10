package aggregates

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/polygon-io/client-golang/rest/client"
)

// todo: add comments for godoc

type Client struct {
	client.HTTPBase
}

// easyjson:json
type Aggregate struct {
	Ticker            string  `json:"T,omitempty"`
	Volume            float64 `json:"v"`
	VWAP              float64 `json:"vw,omitempty"`
	AggregateVWAP     float64 `json:"a,omitempty"`
	Open              float64 `json:"o"`
	Close             float64 `json:"c"`
	High              float64 `json:"h"`
	Low               float64 `json:"l"`
	Timestamp         int64   `json:"t"`
	Transactions      int64   `json:"n,omitempty"`
	Market            string  `json:"m,omitempty"`
	Exchange          int32   `json:"x,omitempty"`
	Locale            string  `json:"g,omitempty"`
	OfficialOpenPrice float64 `json:"op,omitempty"`
	AverageSize       float64 `json:"z,omitempty"`
	AccumulatedVolume float64 `json:"av,omitempty"`
	StartTimestamp    int64   `json:"s,omitempty"`
	EndTimestamp      int64   `json:"e,omitempty"`
}

// easyjson:json
type AggsResponse struct {
	client.BaseResponse
	Ticker       string      `json:"ticker,omitempty"`
	QueryCount   int         `json:"queryCount"`
	ResultsCount int         `json:"resultsCount"`
	Adjusted     bool        `json:"adjusted"`
	Aggs         []Aggregate `json:"results,omitempty"`
}

type GetPathParams struct {
	Ticker     string
	Multiplier int
	Resolution string
	From       time.Time
	To         time.Time
}

func (p GetPathParams) Values() map[string]string {
	return map[string]string{
		"ticker":     p.Ticker,
		"multiplier": fmt.Sprint(p.Multiplier),
		"resolution": fmt.Sprint(p.Resolution),
		"from":       fmt.Sprint(p.From.UnixMilli()),
		"to":         fmt.Sprint(p.To.UnixMilli()),
	}
}

type GetQueryParams struct {
	Sort     string
	Limit    int32
	Adjusted bool
	Explain  bool
}

func (p GetQueryParams) Values() map[string]string {
	v := map[string]string{}

	if p.Sort != "" {
		v["sort"] = p.Sort
	}

	if p.Limit != 0 {
		v["limit"] = strconv.FormatInt(int64(p.Limit), 10)
	}

	if !p.Adjusted {
		v["adjusted"] = "false"
	}

	if p.Explain {
		v["explain"] = "true"
	}

	return v
}

type GetPreviousClosePathParams struct {
	Ticker string
}

func (p GetPreviousClosePathParams) Values() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}

type GetPreviousCloseQueryParams struct {
	Adjusted bool
}

func (p GetPreviousCloseQueryParams) Values() map[string]string {
	v := map[string]string{}

	if !p.Adjusted {
		v["adjusted"] = "false"
	}

	return v
}

func (ac *Client) Get(ctx context.Context, pathParams GetPathParams, queryParams *GetQueryParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	url := "/v2/aggs/ticker/{ticker}/range/{multiplier}/{resolution}/{from}/{to}"
	err := ac.Call(http.MethodGet, url, pathParams, queryParams, res, append([]client.Option{client.WithContext(ctx)}, opts...)...)
	return res, err
}

func (ac *Client) GetPreviousClose(ctx context.Context, pathParams GetPreviousClosePathParams, queryParams *GetPreviousCloseQueryParams, opts ...client.Option) (*AggsResponse, error) {
	res := &AggsResponse{}
	url := "/v2/aggs/ticker/{ticker}/prev"
	err := ac.Call(http.MethodGet, url, pathParams, queryParams, res, append([]client.Option{client.WithContext(ctx)}, opts...)...)
	return res, err
}

// todo: GetGroupedDaily
