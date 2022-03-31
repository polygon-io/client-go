package models

import (
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

// GetAggsParams is the set of path and query parameters for requesting aggs.
type GetAggsParams struct {
	Ticker     string     `validate:"required" path:"ticker"`
	Multiplier int        `validate:"required" path:"multiplier"`
	Resolution Resolution `validate:"required" path:"resolution"`
	From       time.Time  `validate:"required" path:"from"` // p.From.UnixMilli()
	To         time.Time  `validate:"required" path:"to"`   // p.From.UnixMilli()

	Sort     *Order `query:"sort"`
	Limit    *int   `query:"limit"`
	Adjusted *bool  `query:"adjusted"`
	Explain  *bool  `query:"explain"`
}

// GetPreviousCloseParams is the set of path and query parameters for requesting previous close aggs.
type GetPreviousCloseParams struct {
	Ticker string `validate:"required" path:"ticker"`

	Adjusted *bool `query:"adjusted"`
}

// GetGroupedDailyParams is the set of path and query parameters that can be used when requesting aggs through the GetGroupedDaily method.
type GetGroupedDailyParams struct {
	Locale     MarketLocale `validate:"required" path:"locale"`
	MarketType MarketType   `validate:"required" path:"marketType"`
	Date       time.Time    `validate:"required" path:"date"` // p.Date.Format("2006-01-02")

	Adjusted *bool `query:"adjusted"`
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
	Ticker string    `validate:"required" path:"ticker"`
	Date   time.Time `validate:"required" path:"date"` // p.Date.Format("2006-01-02")

	Adjusted *bool `query:"adjusted"`
}
