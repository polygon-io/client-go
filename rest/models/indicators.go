package models

// GetSMAParams is the set of parameters for the GetSMA method.
type GetSMAParams struct {
	// The ticker symbol of the stock/equity.
	Ticker string `validate:"required" path:"ticker"`

	// The size of the timespan multiplier for the underlying aggregates.
	Multiplier int `query:"multiplier"`

	// The size of the timespan of the underlying aggregates.
	Timespan Timespan `query:"timespan"`

	// Query indicators by timestamp.
	TimestampEQ  *Millis `query:"timestamp"`
	TimestampLT  *Millis `query:"timestamp.lt"`
	TimestampLTE *Millis `query:"timestamp.lte"`
	TimestampGT  *Millis `query:"timestamp.gt"`
	TimestampGTE *Millis `query:"timestamp.gte"`

	// The attribute of the underlying aggregate which will be used to calculate the indicator.
	SeriesType *SeriesType `query:"series_type"`

	// Whether to also return the underlying aggregates used to calculate the indicator.
	ExpandUnderlying *bool `query:"expand_underlying"`

	// Whether or not the underlying aggregates used to calculate the indicator are adjusted for splits. By default, the aggregates are adjusted. Set this to false to get
	// results that are NOT adjusted for splits.
	Adjusted *bool `query:"adjusted"`

	// Order the results by timestamp. asc will return results in ascending order (oldest at the top), desc will return
	// results in descending order (newest at the top).
	Order *Order `query:"order"`

	// The size of the window over which the indicator will be calculated.
	Window *int `query:"window"`
}

func (p GetSMAParams) WithAdjusted(q bool) *GetSMAParams {
	p.Adjusted = &q
	return &p
}

func (p GetSMAParams) WithOrder(q Order) *GetSMAParams {
	p.Order = &q
	return &p
}

func (p GetSMAParams) WithExpandUnderlying(q bool) *GetSMAParams {
	p.ExpandUnderlying = &q
	return &p
}

// GetEMAParams is the set of parameters for the GetEMA method.
type GetEMAParams struct {
	// The ticker symbol of the stock/equity.
	Ticker string `validate:"required" path:"ticker"`

	// The size of the timespan multiplier for the underlying aggregates.
	Multiplier int `query:"multiplier"`

	// The size of the timespan of the underlying aggregates.
	Timespan Timespan `query:"timespan"`

	// Query indicators by timestamp.
	TimestampEQ  *Millis `query:"timestamp"`
	TimestampLT  *Millis `query:"timestamp.lt"`
	TimestampLTE *Millis `query:"timestamp.lte"`
	TimestampGT  *Millis `query:"timestamp.gt"`
	TimestampGTE *Millis `query:"timestamp.gte"`

	// The attribute of the underlying aggregate which will be used to calculate the indicator.
	SeriesType *SeriesType `query:"series_type"`

	// Whether to also return the underlying aggregates used to calculate the indicator.
	ExpandUnderlying *bool `query:"expand_underlying"`

	// Whether or not the underlying aggregates used to calculate the indicator are adjusted for splits. By default, the aggregates are adjusted. Set this to false to get
	// results that are NOT adjusted for splits.
	Adjusted *bool `query:"adjusted"`

	// Order the results by timestamp. asc will return results in ascending order (oldest at the top), desc will return
	// results in descending order (newest at the top).
	Order *Order `query:"order"`

	// The size of the window over which the indicator will be calculated.
	Window *int `query:"window"`
}

func (p GetEMAParams) WithAdjusted(q bool) *GetEMAParams {
	p.Adjusted = &q
	return &p
}

func (p GetEMAParams) WithOrder(q Order) *GetEMAParams {
	p.Order = &q
	return &p
}

func (p GetEMAParams) WithExpandUnderlying(q bool) *GetEMAParams {
	p.ExpandUnderlying = &q
	return &p
}

// GetRSIParams is the set of parameters for the GetRSI method.
type GetRSIParams struct {
	// The ticker symbol of the stock/equity.
	Ticker string `validate:"required" path:"ticker"`

	// The size of the timespan multiplier for the underlying aggregates.
	Multiplier int `query:"multiplier"`

	// The size of the timespan of the underlying aggregates.
	Timespan Timespan `query:"timespan"`

	// Query indicators by timestamp.
	TimestampEQ  *Millis `query:"timestamp"`
	TimestampLT  *Millis `query:"timestamp.lt"`
	TimestampLTE *Millis `query:"timestamp.lte"`
	TimestampGT  *Millis `query:"timestamp.gt"`
	TimestampGTE *Millis `query:"timestamp.gte"`

	// The attribute of the underlying aggregate which will be used to calculate the indicator.
	SeriesType *SeriesType `query:"series_type"`

	// Whether to also return the underlying aggregates used to calculate the indicator.
	ExpandUnderlying *bool `query:"expand_underlying"`

	// Whether or not the underlying aggregates used to calculate the indicator are adjusted for splits. By default, the aggregates are adjusted. Set this to false to get
	// results that are NOT adjusted for splits.
	Adjusted *bool `query:"adjusted"`

	// Order the results by timestamp. asc will return results in ascending order (oldest at the top), desc will return
	// results in descending order (newest at the top).
	Order *Order `query:"order"`

	// The size of the window over which the indicator will be calculated.
	Window *int `query:"window"`
}

func (p GetRSIParams) WithAdjusted(q bool) *GetRSIParams {
	p.Adjusted = &q
	return &p
}

func (p GetRSIParams) WithOrder(q Order) *GetRSIParams {
	p.Order = &q
	return &p
}

func (p GetRSIParams) WithExpandUnderlying(q bool) *GetRSIParams {
	p.ExpandUnderlying = &q
	return &p
}

// GetMACDParams is the set of parameters for the GetMACD method.
type GetMACDParams struct {
	// The ticker symbol of the stock/equity.
	Ticker string `validate:"required" path:"ticker"`

	// The size of the timespan multiplier for the underlying aggregates.
	Multiplier int `query:"multiplier"`

	// The size of the timespan of the underlying aggregates.
	Timespan Timespan `query:"timespan"`

	// Query indicators by timestamp.
	TimestampEQ  *Millis `query:"timestamp"`
	TimestampLT  *Millis `query:"timestamp.lt"`
	TimestampLTE *Millis `query:"timestamp.lte"`
	TimestampGT  *Millis `query:"timestamp.gt"`
	TimestampGTE *Millis `query:"timestamp.gte"`

	// The attribute of the underlying aggregate which will be used to calculate the indicator.
	SeriesType *SeriesType `query:"series_type"`

	// Whether to also return the underlying aggregates used to calculate the indicator.
	ExpandUnderlying *bool `query:"expand_underlying"`

	// Whether or not the underlying aggregates used to calculate the indicator are adjusted for splits. By default, the aggregates are adjusted. Set this to false to get
	// results that are NOT adjusted for splits.
	Adjusted *bool `query:"adjusted"`

	// Order the results by timestamp. asc will return results in ascending order (oldest at the top), desc will return
	// results in descending order (newest at the top).
	Order *Order `query:"order"`

	// The size of the window over which the indicator will be calculated.
	ShortWindow int `query:"short_window"`

	// The size of the window over which the indicator will be calculated.
	LongWindow int `query:"long_window"`

	// The size of the window over which the indicator will be calculated.
	SignalWindow int `query:"signal_window"`
}

func (p GetMACDParams) WithAdjusted(q bool) *GetMACDParams {
	p.Adjusted = &q
	return &p
}

func (p GetMACDParams) WithOrder(q Order) *GetMACDParams {
	p.Order = &q
	return &p
}

func (p GetMACDParams) WithExpandUnderlying(q bool) *GetMACDParams {
	p.ExpandUnderlying = &q
	return &p
}

// Response Models

// GetAggsResponse is the response returned by the GetAggs method.
type GetSMAResponse struct {
	BaseResponse
	Results SingleIndicatorResults `json:"results,omitempty"`
}

type GetEMAResponse struct {
	BaseResponse
	Results SingleIndicatorResults `json:"results,omitempty"`
}

type GetRSIResponse struct {
	BaseResponse
	Results SingleIndicatorResults `json:"results,omitempty"`
}

type SingleIndicatorResults struct {
	Underlying UnderlyingResults     `json:"underlying,omitempty"`
	Values     SingleIndicatorValues `json:"values,omitempty"`
}

type UnderlyingResults struct {
	Aggregates []Agg  `json:"aggregates,omitempty"`
	URL        string `json:"url,omitempty"`
}

type SingleIndicatorValues []SingleIndicatorValue

type SingleIndicatorValue struct {
	Timestamp Millis  `json:"timestamp,omitempty"`
	Value     float64 `json:"value,omitempty"`
}

type GetMACDResponse struct {
	BaseResponse
	Results MACDIndicatorResults `json:"results,omitempty"`
}

type MACDIndicatorResults struct {
	Underlying UnderlyingResults   `json:"underlying,omitempty"`
	Values     MACDIndicatorValues `json:"values,omitempty"`
}

type MACDIndicatorValues []MACDIndicatorValue

type MACDIndicatorValue struct {
	Timestamp Millis  `json:"timestamp,omitempty"`
	Value     float64 `json:"value,omitempty"`
	Signal    float64 `json:"signal,omitempty"`
	Histogram float64 `json:"histogram,omitempty"`
}
