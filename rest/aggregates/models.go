package aggregates

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/polygon-io/client-golang/rest/client"
)

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

type AggsResponse struct {
	client.BaseResponse
	Ticker       string      `json:"ticker,omitempty"`
	QueryCount   int         `json:"queryCount"`
	ResultsCount int         `json:"resultsCount"`
	Adjusted     bool        `json:"adjusted"`
	Aggs         []Aggregate `json:"results,omitempty"`
}

func (r *AggsResponse) UnmarshalJSON(data []byte) error {
	type aggsResponse AggsResponse
	var v aggsResponse
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*r = AggsResponse(v)
	return nil
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

type DailyOpenCloseResponse struct {
	client.BaseResponse
	Symbol     string    `json:"symbol"`
	From       time.Time `json:"from"`
	Open       float64   `json:"open"`
	High       float64   `json:"high"`
	Low        float64   `json:"low"`
	Close      float64   `json:"close"`
	Volume     float64   `json:"volume"`
	AfterHours float64   `json:"afterHours"`
	PreMarket  float64   `json:"preMarket"`
}

func (ar *DailyOpenCloseResponse) UnmarshalJSON(data []byte) error {
	type dailyOpenCloseResponse DailyOpenCloseResponse
	var v dailyOpenCloseResponse
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*ar = DailyOpenCloseResponse(v)
	return nil
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
		"date":   fmt.Sprint(p.Date.Format("2006-01-02")),
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
