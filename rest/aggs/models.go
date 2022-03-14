package aggs

import (
	"fmt"
	"strconv"
	"time"

	"github.com/polygon-io/client-golang/rest/client"
)

// Aggregate is an aggregation of all the activity on a specified ticker between the start and end timestamps.
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

type Sort string

const (
	Asc  Sort = "asc"
	Desc Sort = "desc"
)

type Resolution string

const (
	Minute  Resolution = "minute"
	Hour    Resolution = "hour"
	Day     Resolution = "day"
	Week    Resolution = "week"
	Month   Resolution = "month"
	Quarter Resolution = "quarter"
	Year    Resolution = "year"
)

type GetParams struct {
	Ticker      string
	Multiplier  int
	Resolution  Resolution
	From        time.Time
	To          time.Time
	QueryParams *GetQueryParams
}

type GetQueryParams struct {
	Sort     Sort
	Limit    int32
	Adjusted bool
	Explain  bool
}

func (p GetParams) Path() map[string]string {
	return map[string]string{
		"ticker":     p.Ticker,
		"multiplier": fmt.Sprint(p.Multiplier),
		"resolution": fmt.Sprint(p.Resolution),
		"from":       fmt.Sprint(p.From.UnixMilli()),
		"to":         fmt.Sprint(p.To.UnixMilli()),
	}
}

func (p GetParams) Query() map[string]string {
	q := map[string]string{}
	if p.QueryParams == nil {
		return q
	}

	if p.QueryParams.Sort != "" {
		q["sort"] = string(p.QueryParams.Sort)
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

type MarketType string

const (
	Stocks MarketType = "stocks"
	Forex  MarketType = "forex"
	Crypto MarketType = "crypto"
)

type GetGroupedDailyParams struct {
	Locale      string
	MarketType  MarketType
	Date        time.Time
	QueryParams *GetGroupedDailyQueryParams
}

type GetGroupedDailyQueryParams struct {
	Adjusted bool
}

func (p GetGroupedDailyParams) Path() map[string]string {
	return map[string]string{
		"locale":     p.Locale,
		"marketType": string(p.MarketType),
		"date":       p.Date.Format("2006-01-02"),
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

type DailyOpenCloseResponse struct {
	client.BaseResponse
	Symbol     string  `json:"symbol"`
	From       string  `json:"from"` // todo: use ptime
	Open       float64 `json:"open"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Close      float64 `json:"close"`
	Volume     float64 `json:"volume"`
	AfterHours float64 `json:"afterHours"`
	PreMarket  float64 `json:"preMarket"`
}

type GetDailyOpenCloseParams struct {
	Ticker      string
	Date        time.Time
	QueryParams *GetDailyOpenCloseQueryParams
}

type GetDailyOpenCloseQueryParams struct {
	Adjusted bool
}

func (p GetDailyOpenCloseParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
		"date":   p.Date.Format("2006-01-02"),
	}
}

func (p GetDailyOpenCloseParams) Query() map[string]string {
	q := map[string]string{}
	if p.QueryParams == nil {
		return q
	}

	if !p.QueryParams.Adjusted {
		q["adjusted"] = "false"
	}

	return q
}
