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

const (
	GetPath              = "/v2/aggs/ticker/{ticker}/range/{multiplier}/{resolution}/{from}/{to}"
	GetPreviousClosePath = "/v2/aggs/ticker/{ticker}/prev"
	GetGroupedDailyPath  = "/v2/aggs/grouped/locale/{locale}/market/{marketType}/{date}"
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

type GetParams struct {
	Ticker      string
	Multiplier  int
	Resolution  string
	From        time.Time
	To          time.Time
	QueryParams *GetQueryParams
}

type GetQueryParams struct {
	Sort     string
	Limit    int32
	Adjusted bool
	Explain  bool
}

func (p GetParams) Path() map[string]string {
	return map[string]string{
		"ticker":     p.Ticker,
		"multiplier": fmt.Sprint(p.Multiplier),
		"resolution": fmt.Sprint(p.Resolution),
		"from":       fmt.Sprint(p.From.UnixMilli()), // todo: decide if we want to do Format("2006-01-02") instead of UnixMilli()
		"to":         fmt.Sprint(p.To.UnixMilli()),
	}
}

func (p GetParams) Query() map[string]string {
	q := map[string]string{}
	if p.QueryParams == nil {
		return q
	}

	if p.QueryParams.Sort != "" {
		q["sort"] = p.QueryParams.Sort
	}

	if p.QueryParams.Limit != 0 {
		q["limit"] = strconv.FormatInt(int64(p.QueryParams.Limit), 10)
	}

	if !p.QueryParams.Adjusted {
		q["adjusted"] = "false"
	}

	if p.QueryParams.Explain {
		q["explain"] = "true"
	}

	return q
}

type GetPreviousCloseParams struct {
	Ticker      string
	QueryParams *GetPreviousCloseQueryParams
}

type GetPreviousCloseQueryParams struct {
	Adjusted bool
}

func (p GetPreviousCloseParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}

func (p GetPreviousCloseParams) Query() map[string]string {
	q := map[string]string{}
	if p.QueryParams == nil {
		return q
	}

	if !p.QueryParams.Adjusted {
		q["adjusted"] = "false"
	}

	return q
}

type GetGroupedDailyParams struct {
	Locale      string
	MarketType  string
	Date        time.Time
	QueryParams *GetGroupedDailyQueryParams
}

type GetGroupedDailyQueryParams struct {
	Adjusted bool
}

func (p GetGroupedDailyParams) Path() map[string]string {
	return map[string]string{
		"locale":     p.Locale,
		"marketType": p.MarketType,
		"date":       fmt.Sprint(p.Date.Format("2006-01-02")),
	}
}

func (p GetGroupedDailyParams) Query() map[string]string {
	q := map[string]string{}
	if p.QueryParams == nil {
		return q
	}

	if !p.QueryParams.Adjusted {
		q["adjusted"] = "false"
	}

	return q
}
