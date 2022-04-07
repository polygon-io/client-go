package models

const ListSplitsPath = "/v3/reference/splits"

// ListSplitsParams is the set of parameters for the ListSplits method.
type ListSplitsParams struct {
	Ticker    *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	ExeDate    *string `query:"execution_date"` // todo: this is "YYYY-MM-DD" format, need to figure out the best way to encode this without interfering with the default
	ExeDateLT  *string `query:"execution_date.lt"`
	ExeDateLTE *string `query:"execution_date.lte"`
	ExeDateGT  *string `query:"execution_date.gt"`
	ExeDateGTE *string `query:"execution_date.gte"`

	ReverseSplit *bool `query:"reverse_split"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
}

// ListSplitsResponse is the response returned by the ListSplits method.
type ListSplitsResponse struct {
	BaseResponse
	Results []Split `json:"results,omitempty"`
}

// Split contains detailed information on a specified stock split.
type Split struct {
	ExeDate string `json:"execution_date,omitempty"`
	From    int64  `json:"split_from,omitempty"`
	To      int64  `json:"split_to,omitempty"`
	Ticker  string `json:"ticker,omitempty"`
}
