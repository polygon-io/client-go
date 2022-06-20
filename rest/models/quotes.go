package models

// ListQuotesParams is the set of parameters for the ListQuotes method.
type ListQuotesParams struct {
	// The ticker symbol to get quotes for.
	Ticker string `validate:"required" path:"ticker"`

	// Query by quote timestamp. To query for a specific day instead of a nanosecond timestamp, set this field
	// via this pattern: WithTimestamp(models.EQ, models.Nanos(time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC))).
	TimestampEQ  *Nanos `query:"timestamp"`
	TimestampLT  *Nanos `query:"timestamp.lt"`
	TimestampLTE *Nanos `query:"timestamp.lte"`
	TimestampGT  *Nanos `query:"timestamp.gt"`
	TimestampGTE *Nanos `query:"timestamp.gte"`

	// Order results based on the sort field.
	Order *Order `query:"order"`

	// Limit the number of results returned, default is 10 and max is 50000.
	Limit *int `query:"limit"`

	// Sort field used for ordering.
	Sort *Sort `query:"sort"`
}

func (p ListQuotesParams) WithTimestamp(c Comparator, q Nanos) *ListQuotesParams {
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

func (p ListQuotesParams) WithOrder(q Order) *ListQuotesParams {
	p.Order = &q
	return &p
}

func (p ListQuotesParams) WithLimit(q int) *ListQuotesParams {
	p.Limit = &q
	return &p
}

func (p ListQuotesParams) WithSort(q Sort) *ListQuotesParams {
	p.Sort = &q
	return &p
}

// ListQuotesResponse is the response returned by the ListQuotes method.
type ListQuotesResponse struct {
	BaseResponse
	Results []Quote `json:"results,omitempty"`
}

// GetLastQuoteParams is the set of parameters for the GetLastQuote method.
type GetLastQuoteParams struct {
	// The ticker symbol of the stock/equity.
	Ticker string `validate:"required" path:"ticker"`
}

// GetLastQuoteResponse is the response returned by the GetLastQuote method.
type GetLastQuoteResponse struct {
	BaseResponse
	Results LastQuote `json:"results,omitempty"`
}

// GetLastForexQuoteParams is the set of parameters for the GetLastForexQuote method.
type GetLastForexQuoteParams struct {
	// The "from" symbol of the pair.
	From string `validate:"required" path:"from"`

	// The "to" symbol of the pair.
	To string `validate:"required" path:"to"`
}

// GetLastForexQuoteResponse is the response returned by the GetLastForexQuote method.
type GetLastForexQuoteResponse struct {
	BaseResponse
	Symbol string     `json:"symbol,omitempty"`
	Last   ForexQuote `json:"last,omitempty"`
}

// GetRealTimeCurrencyConversionParams is the set of parameters for the GetRealTimeCurrencyConversion method.
type GetRealTimeCurrencyConversionParams struct {
	From string `validate:"required" path:"from"`
	To   string `validate:"required" path:"to"`
}

// GetRealTimeCurrencyConversionResponse is the response returned by the GetRealTimeCurrencyConversion method.
type GetRealTimeCurrencyConversionResponse struct {
	BaseResponse
	InitialAmount float64    `json:"initialAmount,omitempty"`
	Converted     float64    `json:"converted,omitempty"`
	From          string     `json:"from,omitempty"`
	To            string     `json:"to,omitempty"`
	LastQuote     ForexQuote `json:"last,omitempty"`
}

// Quote is an NBBO for a ticker symbol in a given time range.
type Quote struct {
	AskExchange          int     `json:"ask_exchange,omitempty"`
	AskPrice             float64 `json:"ask_price,omitempty"`
	AskSize              float64 `json:"ask_size,omitempty"`
	BidExchange          int     `json:"bid_exchange,omitempty"`
	BidPrice             float64 `json:"bid_price,omitempty"`
	BidSize              float64 `json:"bid_size,omitempty"`
	Conditions           []int32 `json:"conditions,omitempty"`
	Indicators           []int32 `json:"indicators,omitempty"`
	ParticipantTimestamp Nanos   `json:"participant_timestamp,omitempty"`
	SequenceNumber       Nanos   `json:"sequence_number,omitempty"`
	SipTimestamp         Nanos   `json:"sip_timestamp,omitempty"`
	Tape                 int32   `json:"tape,omitempty"`
	TrfTimestamp         Nanos   `json:"trf_timestamp,omitempty"`
}

// LastQuote is the most recent NBBO for a ticker symbol.
type LastQuote struct {
	Ticker               string  `json:"T,omitempty"`
	TrfTimestamp         Nanos   `json:"f,omitempty"`
	SequenceNumber       Nanos   `json:"q,omitempty"`
	SipTimestamp         Nanos   `json:"t,omitempty"`
	ParticipantTimestamp Nanos   `json:"y,omitempty"`
	AskPrice             float64 `json:"P,omitempty"`
	AskSize              float64 `json:"S,omitempty"`
	AskExchange          int     `json:"X,omitempty"`
	Conditions           []int32 `json:"c,omitempty"`
	Indicators           []int32 `json:"i,omitempty"`
	BidPrice             float64 `json:"p,omitempty"`
	BidSize              float64 `json:"s,omitempty"`
	BidExchange          int     `json:"x,omitempty"`
	Tape                 int32   `json:"z,omitempty"`
}

// ForexQuote is a BBO for a forex currency pair.
type ForexQuote struct {
	Ask       float64 `json:"ask,omitempty"`
	Bid       float64 `json:"bid,omitempty"`
	Exchange  int     `json:"exchange,omitempty"`
	Timestamp Nanos   `json:"timestamp,omitempty"`
}
