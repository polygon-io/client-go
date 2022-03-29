package models

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/polygon-io/client-go/rest/client"
)

const (
	GetAggsPath           = "/v2/aggs/ticker/{ticker}/range/{multiplier}/{resolution}/{from}/{to}"
	GetPreviousClosePath  = "/v2/aggs/ticker/{ticker}/prev"
	GetGroupedDailyPath   = "/v2/aggs/grouped/locale/{locale}/market/{marketType}/{date}"
	GetDailyOpenClosePath = "/v1/open-close/{ticker}/{date}"
)

// Aggregate is an aggregation of all the activity on a specified ticker between the start and end timestamps.
// For more details see https://polygon.io/docs/stocks/get_v2_aggs_ticker__stocksticker__range__multiplier___timespan___from___to.
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

// AggsResponse is returned by the aggs API. It contains a list of aggregates for the specified ticker.
type AggsResponse struct {
	client.BaseResponse
	Ticker       string      `json:"ticker,omitempty"`
	QueryCount   int         `json:"queryCount"`
	ResultsCount int         `json:"resultsCount"`
	Adjusted     bool        `json:"adjusted"`
	Aggs         []Aggregate `json:"results,omitempty"`
}

// GetAggsParams is the set of path and query parameters that can be used when requesting aggs through the Get method.
type GetAggsParams struct {
	Ticker      string
	Multiplier  int
	Resolution  Resolution
	From        time.Time
	To          time.Time
	QueryParams GetAggsQueryParams
}

// GetAggsQueryParams is the set of query parameters that can be used when requesting aggs through the Get method.
type GetAggsQueryParams struct {
	Sort     *Order
	Limit    *int
	Adjusted *bool
	Explain  *bool
}

// Path maps the input Get parameters to their respective keys.
func (p GetAggsParams) Path() map[string]string {
	return map[string]string{
		"ticker":     p.Ticker,
		"multiplier": fmt.Sprint(p.Multiplier),
		"resolution": fmt.Sprint(p.Resolution),
		"from":       fmt.Sprint(p.From.UnixMilli()),
		"to":         fmt.Sprint(p.To.UnixMilli()),
	}
}

// Query maps the input Get parameters to their respective keys.
func (p GetAggsParams) Query() url.Values {
	q := url.Values{}

	if p.QueryParams.Sort != nil {
		q.Set("sort", string(*p.QueryParams.Sort))
	}

	if p.QueryParams.Limit != nil {
		q.Set("limit", strconv.FormatInt(int64(*p.QueryParams.Limit), 10))
	}

	if p.QueryParams.Adjusted != nil {
		q.Set("adjusted", strconv.FormatBool(*p.QueryParams.Adjusted))
	}

	if p.QueryParams.Explain != nil {
		q.Set("explain", strconv.FormatBool(*p.QueryParams.Explain))
	}

	return q
}

// GetPreviousCloseParams is the set of path and query parameters that can be used when requesting aggs through the GetPreviousClose method.
type GetPreviousCloseParams struct {
	Ticker      string
	QueryParams GetPreviousCloseQueryParams
}

// GetPreviousCloseQueryParams is the set of query parameters that can be used when requesting aggs through the GetPreviousClose method.
type GetPreviousCloseQueryParams struct {
	Adjusted *bool
}

// Path maps the GetPreviousCloseParams path parameters to their respective keys.
func (p GetPreviousCloseParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}

// Query maps the GetPreviousCloseParams query parameters to their respective keys.
func (p GetPreviousCloseParams) Query() url.Values {
	q := url.Values{}

	if p.QueryParams.Adjusted != nil {
		q.Set("adjusted", strconv.FormatBool(*p.QueryParams.Adjusted))
	}

	return q
}

// GetGroupedDailyParams is the set of path and query parameters that can be used when requesting aggs through the GetGroupedDaily method.
type GetGroupedDailyParams struct {
	Locale      MarketLocale
	MarketType  MarketType
	Date        time.Time
	QueryParams GetGroupedDailyQueryParams
}

// GetGroupedDailyQueryParams is the set of query parameters that can be used when requesting aggs through the GetGroupedDaily method.
type GetGroupedDailyQueryParams struct {
	Adjusted *bool
}

// Path maps the GetGroupedDaily path parameters to their respective keys.
func (p GetGroupedDailyParams) Path() map[string]string {
	return map[string]string{
		"locale":     string(p.Locale),
		"marketType": string(p.MarketType),
		"date":       p.Date.Format("2006-01-02"),
	}
}

// Query maps the GetGroupedDaily query parameters to their respective keys.
func (p GetGroupedDailyParams) Query() url.Values {
	q := url.Values{}

	if p.QueryParams.Adjusted != nil {
		q.Set("adjusted", strconv.FormatBool(*p.QueryParams.Adjusted))
	}

	return q
}

// DailyOpenCloseResponse is the response for the DailyOpenClose method.
// Get the open, close and afterhours prices of a stock symbol on a certain date.
// For more details see https://polygon.io/docs/stocks/get_v1_open-close__stocksticker___date.
type DailyOpenCloseResponse struct {
	client.BaseResponse
	Symbol     string  `json:"symbol"`
	From       string  `json:"from"`
	Open       float64 `json:"open"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Close      float64 `json:"close"`
	Volume     float64 `json:"volume"`
	AfterHours float64 `json:"afterHours"`
	PreMarket  float64 `json:"preMarket"`
}

// GetDailyOpenCloseParams is the set of path and query parameters that can be used when requesting aggs through the GetDailyOpenClose method.
type GetDailyOpenCloseParams struct {
	Ticker      string
	Date        time.Time
	QueryParams GetDailyOpenCloseQueryParams
}

// GetDailyOpenCloseQueryParams is the set of query parameters that can be used when requesting aggs through the GetDailyOpenCloseQuery method.
type GetDailyOpenCloseQueryParams struct {
	Adjusted *bool
}

// Path maps the GetDailyOpenClose path parameters to their respective keys.
func (p GetDailyOpenCloseParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
		"date":   p.Date.Format("2006-01-02"),
	}
}

// Query maps the GetDailyOpenClose query parameters to their respective keys.
func (p GetDailyOpenCloseParams) Query() url.Values {
	q := url.Values{}

	if p.QueryParams.Adjusted != nil {
		q.Set("adjusted", strconv.FormatBool(*p.QueryParams.Adjusted))
	}

	return q
}
