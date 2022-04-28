package models

// GetAggsParams is the set of parameters for the GetAggs method.
type GetAggsParams struct {
	Ticker     string     `validate:"required" path:"ticker"`
	Multiplier int        `validate:"required" path:"multiplier"`
	Resolution Resolution `validate:"required" path:"resolution"`
	From       Millis     `validate:"required" path:"from"`
	To         Millis     `validate:"required" path:"to"`

	Order    *Order `query:"sort"`
	Limit    *int   `query:"limit"`
	Adjusted *bool  `query:"adjusted"`
	Explain  *bool  `query:"explain"`
}

func (p GetAggsParams) WithOrder(q Order) *GetAggsParams {
	p.Order = &q
	return &p
}

func (p GetAggsParams) WithLimit(q int) *GetAggsParams {
	p.Limit = &q
	return &p
}

func (p GetAggsParams) WithAdjusted(q bool) *GetAggsParams {
	p.Adjusted = &q
	return &p
}

func (p GetAggsParams) WithExplain(q bool) *GetAggsParams {
	p.Explain = &q
	return &p
}

// GetAggsResponse is the response returned by the GetAggs method.
type GetAggsResponse struct {
	BaseResponse
	Ticker       string `json:"ticker,omitempty"`
	QueryCount   int    `json:"queryCount,omitempty"`
	ResultsCount int    `json:"resultsCount,omitempty"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Agg  `json:"results,omitempty"`
}

// GetGroupedDailyAggsParams is the set of parameters for the GetGroupedDailyAggs method.
type GetGroupedDailyAggsParams struct {
	Locale     MarketLocale `validate:"required" path:"locale"`
	MarketType MarketType   `validate:"required" path:"marketType"`
	Date       Date         `validate:"required" path:"date"`

	Adjusted *bool `query:"adjusted"`
}

func (p GetGroupedDailyAggsParams) WithAdjusted(q bool) *GetGroupedDailyAggsParams {
	p.Adjusted = &q
	return &p
}

// GetGroupedDailyAggsResponse is the response returned by the GetGroupedDailyAggs method.
type GetGroupedDailyAggsResponse struct {
	BaseResponse
	Ticker       string `json:"ticker,omitempty"`
	QueryCount   int    `json:"queryCount,omitempty"`
	ResultsCount int    `json:"resultsCount,omitempty"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Agg  `json:"results,omitempty"`
}

// GetDailyOpenCloseAggParams is the set of parameters for the GetDailyOpenCloseAgg method.
type GetDailyOpenCloseAggParams struct {
	Ticker string `validate:"required" path:"ticker"`
	Date   Date   `validate:"required" path:"date"`

	Adjusted *bool `query:"adjusted"`
}

func (p GetDailyOpenCloseAggParams) WithAdjusted(q bool) *GetDailyOpenCloseAggParams {
	p.Adjusted = &q
	return &p
}

// GetDailyOpenCloseAggResponse is the response for the GetDailyOpenCloseAgg method.
type GetDailyOpenCloseAggResponse struct {
	BaseResponse
	Symbol     string  `json:"symbol,omitempty"`
	From       string  `json:"from,omitempty"`
	Open       float64 `json:"open,omitempty"`
	High       float64 `json:"high,omitempty"`
	Low        float64 `json:"low,omitempty"`
	Close      float64 `json:"close,omitempty"`
	Volume     float64 `json:"volume,omitempty"`
	AfterHours float64 `json:"afterHours,omitempty"`
	PreMarket  float64 `json:"preMarket,omitempty"`
}

// GetPreviousCloseAggParams is the set of parameters for the GetPreviousCloseAgg method.
type GetPreviousCloseAggParams struct {
	Ticker string `validate:"required" path:"ticker"`

	Adjusted *bool `query:"adjusted"`
}

func (p GetPreviousCloseAggParams) WithAdjusted(q bool) *GetPreviousCloseAggParams {
	p.Adjusted = &q
	return &p
}

// GetPreviousCloseAggResponse is the response returned by the GetPreviousCloseAgg method.
type GetPreviousCloseAggResponse struct {
	BaseResponse
	Ticker       string `json:"ticker,omitempty"`
	QueryCount   int    `json:"queryCount,omitempty"`
	ResultsCount int    `json:"resultsCount,omitempty"`
	Adjusted     bool   `json:"adjusted"`
	Results      []Agg  `json:"results,omitempty"`
}

// Agg is an aggregation of all the activity on a specified ticker between the start and end timestamps.
type Agg struct {
	Ticker            string  `json:"T,omitempty"`
	Volume            float64 `json:"v,omitempty"`
	VWAP              float64 `json:"vw,omitempty"`
	AggregateVWAP     float64 `json:"a,omitempty"`
	Open              float64 `json:"o,omitempty"`
	Close             float64 `json:"c,omitempty"`
	High              float64 `json:"h,omitempty"`
	Low               float64 `json:"l,omitempty"`
	Timestamp         Millis  `json:"t,omitempty"`
	Transactions      int64   `json:"n,omitempty"`
	Market            string  `json:"m,omitempty"`
	Exchange          int32   `json:"x,omitempty"`
	Locale            string  `json:"g,omitempty"`
	OfficialOpenPrice float64 `json:"op,omitempty"`
	AverageSize       float64 `json:"z,omitempty"`
	AccumulatedVolume float64 `json:"av,omitempty"`
	StartTimestamp    Millis  `json:"s,omitempty"`
	EndTimestamp      Millis  `json:"e,omitempty"`
}
