package models

const ListSplitsPath = "/v3/reference/splits"

// ListSplitsParams is the set of parameters for the ListSplits method.
type ListSplitsParams struct {
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	ExecutionDateEQ  *string `query:"execution_date"` // todo: this is "YYYY-MM-DD" format, need to figure out the best way to encode this without interfering with the default
	ExecutionDateLT  *string `query:"execution_date.lt"`
	ExecutionDateLTE *string `query:"execution_date.lte"`
	ExecutionDateGT  *string `query:"execution_date.gt"`
	ExecutionDateGTE *string `query:"execution_date.gte"`

	ReverseSplit *bool `query:"reverse_split"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
}

func (p ListSplitsParams) WithTicker(c Comparator, q string) *ListSplitsParams {
	if c == EQ {
		p.TickerEQ = &q
	} else if c == LT {
		p.TickerLT = &q
	} else if c == LTE {
		p.TickerLTE = &q
	} else if c == GT {
		p.TickerGT = &q
	} else if c == GTE {
		p.TickerGTE = &q
	}
	return &p
}

func (p ListSplitsParams) WithExecutionDate(c Comparator, q string) *ListSplitsParams {
	if c == EQ {
		p.ExecutionDateEQ = &q
	} else if c == LT {
		p.ExecutionDateLT = &q
	} else if c == LTE {
		p.ExecutionDateLTE = &q
	} else if c == GT {
		p.ExecutionDateGT = &q
	} else if c == GTE {
		p.ExecutionDateGTE = &q
	}
	return &p
}

func (p ListSplitsParams) WithReverseSplit(q bool) *ListSplitsParams {
	p.ReverseSplit = &q
	return &p
}

func (p ListSplitsParams) WithSort(q Sort) *ListSplitsParams {
	p.Sort = &q
	return &p
}

func (p ListSplitsParams) WithOrder(q Order) *ListSplitsParams {
	p.Order = &q
	return &p
}

func (p ListSplitsParams) WithLimit(q int) *ListSplitsParams {
	p.Limit = &q
	return &p
}

// ListSplitsResponse is the response returned by the ListSplits method.
type ListSplitsResponse struct {
	BaseResponse
	Results []Split `json:"results,omitempty"`
}

// Split contains detailed information on a specified stock split.
type Split struct {
	ExecutionDate string `json:"execution_date,omitempty"`
	SplitFrom     int64  `json:"split_from,omitempty"`
	SplitTo       int64  `json:"split_to,omitempty"`
	Ticker        string `json:"ticker,omitempty"`
}
