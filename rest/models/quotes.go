package models

import (
	"time"

	"github.com/polygon-io/client-go/rest/client"
)

const (
	ListQuotesPath        = "/v3/quotes/{ticker}"
	GetLastQuotePath      = "/v2/last/nbbo/{ticker}"
	GetLastForexQuotePath = "/v1/last_quote/currencies/{from}/{to}"
)

// Quote is an NBBO for a ticker symbol in a given time range.
// For more details see https://polygon.io/docs/stocks/get_v3_quotes__stockticker.
type Quote struct {
	AskExchange          int     `json:"ask_exchange,omitempty"`
	AskPrice             float64 `json:"ask_price,omitempty"`
	AskSize              float64 `json:"ask_size,omitempty"`
	BidExchange          int     `json:"bid_exchange,omitempty"`
	BidPrice             float64 `json:"bid_price,omitempty"`
	BidSize              float64 `json:"bid_size,omitempty"`
	Conditions           []int32 `json:"conditions,omitempty"`
	Indicators           []int32 `json:"indicators,omitempty"`
	ParticipantTimestamp int64   `json:"participant_timestamp,omitempty"`
	SequenceNumber       int64   `json:"sequence_number,omitempty"`
	SipTimestamp         int64   `json:"sip_timestamp,omitempty"`
	Tape                 int32   `json:"tape,omitempty"`
	TrfTimestamp         int64   `json:"trf_timestamp,omitempty"`
}

// QuotesResponse is returned by the list quotes API and contains a list of quotes for the specified ticker.
type QuotesResponse struct {
	client.BaseResponse
	Results []Quote `json:"results,omitempty"`
}

// ListQuotesParams is the set of path and query parameters that are used when requesting quotes via the ListQuotes method.
type ListQuotesParams struct {
	Ticker string `validate:"required"`

	TimestampEQ  *time.Time `query:"timestamp"`
	TimestampLT  *time.Time `query:"timestamp.lt"`
	TimestampLTE *time.Time `query:"timestamp.lte"`
	TimestampGT  *time.Time `query:"timestamp.gt"`
	TimestampGTE *time.Time `query:"timestamp.gte"`
	Order        *Order     `query:"order"`
	Limit        *int       `query:"limit"`
	Sort         *Sort      `query:"sort"`
}

// Path returns a map of URL path parameters.
func (p ListQuotesParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}

// LastQuote is the most recent NBBO for a ticker symbol.
// For more details see https://polygon.io/docs/stocks/get_v2_last_nbbo__stocksticker.
type LastQuote struct {
	Ticker               string  `json:"T,omitempty"`
	AskExchange          int     `json:"X,omitempty"`
	AskPrice             float64 `json:"P,omitempty"`
	AskSize              float64 `json:"S,omitempty"`
	BidExchange          int     `json:"x,omitempty"`
	BidPrice             float64 `json:"p,omitempty"`
	BidSize              float64 `json:"s,omitempty"`
	Conditions           []int32 `json:"c,omitempty"`
	Indicators           []int32 `json:"i,omitempty"`
	ParticipantTimestamp int64   `json:"y,omitempty"`
	SequenceNumber       int64   `json:"q,omitempty"`
	SipTimestamp         int64   `json:"t,omitempty"`
	Tape                 int32   `json:"z,omitempty"`
	TrfTimestamp         int64   `json:"f,omitempty"`
}

// LastQuoteResponse contains the most recent quote (NBBO) for a specified ticker.
type LastQuoteResponse struct {
	client.BaseResponse
	Results LastQuote `json:"results,omitempty"`
}

// GetLastQuoteParams is the set of path and query parameters for retrieving the most recent quote (NBBO) for a specified ticker.
type GetLastQuoteParams struct {
	Ticker string `validate:"required"`
}

// Path returns a map of URL path parameters.
func (p GetLastQuoteParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}

// ForexQuote is a BBO for a forex currency pair.
// For more details see https://polygon.io/docs/forex/get_v1_last_quote_currencies__from___to.
type ForexQuote struct {
	Ask       float64 `json:"ask,omitempty"`
	Bid       float64 `json:"bid,omitempty"`
	Exchange  int     `json:"exchange,omitempty"`
	Timestamp int64   `json:"timestamp,omitempty"`
}

// LastForexQuoteResponse contains the most recent quote (BBO) for a forex currency pair.
type LastForexQuoteResponse struct {
	client.BaseResponse
	Last ForexQuote `json:"last,omitempty"`
}

// LastForexQuoteParams is the set of path and query parameters for retrieving the most recent quote (BBO) for a forex currency pair.
type LastForexQuoteParams struct {
	From string `validate:"required"`
	To   string `validate:"required"`
}

// Path returns a map of URL path parameters.
func (p LastForexQuoteParams) Path() map[string]string {
	return map[string]string{
		"from": p.From,
		"to":   p.To,
	}
}
