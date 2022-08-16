package models

// GetSMAParams is the set of parameters for the GetSMA method.
type GetSMAParams struct {
	// The ticker symbol of the stock/equity.
	Ticker string `validate:"required" path:"ticker"`

	// The size of the timespan multiplier for the underlying aggregates.
	Multiplier *int `query:"multiplier"`

	// The size of the timespan of the underlying aggregates.
	Timespan *Timespan `query:"timespan"`

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

func (p GetSMAParams) WithTimestamp(c Comparator, q Millis) *GetSMAParams {
	switch c {
	case EQ:
		p.TimestampEQ = &q
	case LT:
		p.TimestampLT = &q
	case LTE:
		p.TimestampLTE = &q
	case GT:
		p.TimestampGT = &q
	case GTE:
		p.TimestampGTE = &q
	}
	return &p
}

func (p GetSMAParams) WithMultiplier(q int) *GetSMAParams {
	p.Multiplier = &q
	return &p
}

func (p GetSMAParams) WithTimespan(q Timespan) *GetSMAParams {
	p.Timespan = &q
	return &p
}

func (p GetSMAParams) WithSeriesType(q SeriesType) *GetSMAParams {
	p.SeriesType = &q
	return &p
}

func (p GetSMAParams) WithWindow(q int) *GetSMAParams {
	p.Window = &q
	return &p
}

// GetEMAParams is the set of parameters for the GetEMA method.
type GetEMAParams struct {
	// The ticker symbol of the stock/equity.
	Ticker string `validate:"required" path:"ticker"`

	// The size of the timespan multiplier for the underlying aggregates.
	Multiplier *int `query:"multiplier"`

	// The size of the timespan of the underlying aggregates.
	Timespan *Timespan `query:"timespan"`

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

func (p GetEMAParams) WithTimestamp(c Comparator, q Millis) *GetEMAParams {
	switch c {
	case EQ:
		p.TimestampEQ = &q
	case LT:
		p.TimestampLT = &q
	case LTE:
		p.TimestampLTE = &q
	case GT:
		p.TimestampGT = &q
	case GTE:
		p.TimestampGTE = &q
	}
	return &p
}

func (p GetEMAParams) WithMultiplier(q int) *GetEMAParams {
	p.Multiplier = &q
	return &p
}

func (p GetEMAParams) WithTimespan(q Timespan) *GetEMAParams {
	p.Timespan = &q
	return &p
}

func (p GetEMAParams) WithSeriesType(q SeriesType) *GetEMAParams {
	p.SeriesType = &q
	return &p
}

func (p GetEMAParams) WithWindow(q int) *GetEMAParams {
	p.Window = &q
	return &p
}

// GetRSIParams is the set of parameters for the GetRSI method.
type GetRSIParams struct {
	// The ticker symbol of the stock/equity.
	Ticker string `validate:"required" path:"ticker"`

	// The size of the timespan multiplier for the underlying aggregates.
	Multiplier *int `query:"multiplier"`

	// The size of the timespan of the underlying aggregates.
	Timespan *Timespan `query:"timespan"`

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

func (p GetRSIParams) WithTimestamp(c Comparator, q Millis) *GetRSIParams {
	switch c {
	case EQ:
		p.TimestampEQ = &q
	case LT:
		p.TimestampLT = &q
	case LTE:
		p.TimestampLTE = &q
	case GT:
		p.TimestampGT = &q
	case GTE:
		p.TimestampGTE = &q
	}
	return &p
}

func (p GetRSIParams) WithMultiplier(q int) *GetRSIParams {
	p.Multiplier = &q
	return &p
}

func (p GetRSIParams) WithTimespan(q Timespan) *GetRSIParams {
	p.Timespan = &q
	return &p
}

func (p GetRSIParams) WithSeriesType(q SeriesType) *GetRSIParams {
	p.SeriesType = &q
	return &p
}

func (p GetRSIParams) WithWindow(q int) *GetRSIParams {
	p.Window = &q
	return &p
}

// GetMACDParams is the set of parameters for the GetMACD method.
type GetMACDParams struct {
	// The ticker symbol of the stock/equity.
	Ticker string `validate:"required" path:"ticker"`

	// The size of the timespan multiplier for the underlying aggregates.
	Multiplier *int `query:"multiplier"`

	// The size of the timespan of the underlying aggregates.
	Timespan *Timespan `query:"timespan"`

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
	ShortWindow *int `query:"short_window"`

	// The size of the window over which the indicator will be calculated.
	LongWindow *int `query:"long_window"`

	// The size of the window over which the indicator will be calculated.
	SignalWindow *int `query:"signal_window"`
}

func (p GetMACDParams) WithTimestamp(c Comparator, q Millis) *GetMACDParams {
	switch c {
	case EQ:
		p.TimestampEQ = &q
	case LT:
		p.TimestampLT = &q
	case LTE:
		p.TimestampLTE = &q
	case GT:
		p.TimestampGT = &q
	case GTE:
		p.TimestampGTE = &q
	}
	return &p
}

func (p GetMACDParams) WithMultiplier(q int) *GetMACDParams {
	p.Multiplier = &q
	return &p
}

func (p GetMACDParams) WithTimespan(q Timespan) *GetMACDParams {
	p.Timespan = &q
	return &p
}

func (p GetMACDParams) WithSeriesType(q SeriesType) *GetMACDParams {
	p.SeriesType = &q
	return &p
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

func (p GetMACDParams) WithShortWindow(q int) *GetMACDParams {
	p.ShortWindow = &q
	return &p
}

func (p GetMACDParams) WithLongWindow(q int) *GetMACDParams {
	p.LongWindow = &q
	return &p
}

func (p GetMACDParams) WithSignalWindow(q int) *GetMACDParams {
	p.SignalWindow = &q
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
