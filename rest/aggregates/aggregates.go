package aggregates

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
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
type GetResponse struct {
	client.BaseResponse
	Ticker       string      `json:"ticker,omitempty"`
	QueryCount   int         `json:"queryCount"`
	ResultsCount int         `json:"resultsCount"`
	Adjusted     bool        `json:"adjusted"`
	Aggs         []Aggregate `json:"results,omitempty"`
}

// todo: could possibly use github.com/google/go-querystring here
type GetQueryParams struct {
	Sort     string
	Limit    int32
	Adjusted bool
	Explain  bool
}

// todo: maybe return map[string]string here (especially if we add a GetPathParams type)
func (p *GetQueryParams) Values() url.Values {
	v := url.Values{}

	if p.Sort != "" {
		v.Add("sort", p.Sort)
	}

	if p.Limit != 0 {
		v.Add("limit", strconv.FormatInt(int64(p.Limit), 10))
	}

	if !p.Adjusted {
		v.Add("adjusted", "false")
	}

	if p.Explain {
		v.Add("explain", "true")
	}

	return v
}

// todo: not a fan of the Sprintf, maybe a GetPathParams type would be cleaner (and easier to godoc)
func (ac *Client) Get(ctx context.Context, ticker string, multiplier int, resolution string, from, to time.Time, params *GetQueryParams, opts ...client.Option) (*GetResponse, error) {
	res := &GetResponse{}
	u := fmt.Sprintf("/v2/aggs/ticker/%s/range/%d/%s/%d/%d", ticker, multiplier, resolution, from.UnixMilli(), to.UnixMilli())
	err := ac.Call(ctx, http.MethodGet, u, params, res, opts...)

	return res, err
}

// todo: GetPreviousDay

// todo: GetGroupedDaily
