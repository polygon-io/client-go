package models

import (
	"strconv"
	"time"
)

const (
	GetAggsPath              = "/v2/aggs/ticker/{ticker}/range/{multiplier}/{resolution}/{from}/{to}"
	GetGroupedDailyAggsPath  = "/v2/aggs/grouped/locale/{locale}/market/{marketType}/{date}"
	GetDailyOpenCloseAggPath = "/v1/open-close/{ticker}/{date}"
	GetPreviousCloseAggPath  = "/v2/aggs/ticker/{ticker}/prev"
)

type TimeMillis time.Time

func (t TimeMillis) ToTime() time.Time {
	return time.Time(t)
}

func (t *TimeMillis) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = TimeMillis(time.Unix(0, millis*int64(time.Millisecond)))
	return nil
}

// GetAggsParams is the set of parameters for the GetAggs method.
type GetAggsParams struct {
	Ticker     string     `validate:"required" path:"ticker"`
	Multiplier int        `validate:"required" path:"multiplier"`
	Resolution Resolution `validate:"required" path:"resolution"`
	From       time.Time  `validate:"required" path:"from" milli:"from"`
	To         time.Time  `validate:"required" path:"to" milli:"to"`

	Sort     *Order `query:"sort"`
	Limit    *int   `query:"limit"`
	Adjusted *bool  `query:"adjusted"`
	Explain  *bool  `query:"explain"`
}

// GetAggsResponse is the response returned by the GetAggs method.
type GetAggsResponse struct {
	BaseResponse
	Ticker       string `json:"ticker,omitempty"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Agg  `json:"results,omitempty"`
}

// GetGroupedDailyAggsParams is the set of parameters for the GetGroupedDailyAggs method.
type GetGroupedDailyAggsParams struct {
	Locale     MarketLocale `validate:"required" path:"locale"`
	MarketType MarketType   `validate:"required" path:"marketType"`
	Date       time.Time    `validate:"required" path:"date"`

	Adjusted *bool `query:"adjusted"`
}

// GetGroupedDailyAggsResponse is the response returned by the GetGroupedDailyAggs method.
type GetGroupedDailyAggsResponse struct {
	BaseResponse
	Ticker       string `json:"ticker,omitempty"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Agg  `json:"results,omitempty"`
}

// GetDailyOpenCloseAggParams is the set of parameters for the GetDailyOpenCloseAgg method.
type GetDailyOpenCloseAggParams struct {
	Ticker string    `validate:"required" path:"ticker"`
	Date   time.Time `validate:"required" path:"date"`

	Adjusted *bool `query:"adjusted"`
}

// GetDailyOpenCloseAggResponse is the response for the GetDailyOpenCloseAgg method.
type GetDailyOpenCloseAggResponse struct {
	BaseResponse
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

// GetPreviousCloseAggParams is the set of parameters for the GetPreviousCloseAgg method.
type GetPreviousCloseAggParams struct {
	Ticker string `validate:"required" path:"ticker"`

	Adjusted *bool `query:"adjusted"`
}

// GetPreviousCloseAggResponse is the response returned by the GetPreviousCloseAgg method.
type GetPreviousCloseAggResponse struct {
	BaseResponse
	Ticker       string `json:"ticker,omitempty"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Agg  `json:"results,omitempty"`
}

// Agg is an aggregation of all the activity on a specified ticker between the start and end timestamps.
type Agg struct {
	Ticker            string     `json:"T,omitempty"`
	Volume            float64    `json:"v"`
	VWAP              float64    `json:"vw,omitempty"`
	AggregateVWAP     float64    `json:"a,omitempty"`
	Open              float64    `json:"o"`
	Close             float64    `json:"c"`
	High              float64    `json:"h"`
	Low               float64    `json:"l"`
	Timestamp         TimeMillis `json:"t"`
	Transactions      int64      `json:"n,omitempty"`
	Market            string     `json:"m,omitempty"`
	Exchange          int32      `json:"x,omitempty"`
	Locale            string     `json:"g,omitempty"`
	OfficialOpenPrice float64    `json:"op,omitempty"`
	AverageSize       float64    `json:"z,omitempty"`
	AccumulatedVolume float64    `json:"av,omitempty"`
	StartTimestamp    int64      `json:"s,omitempty"`
	EndTimestamp      int64      `json:"e,omitempty"`
}
