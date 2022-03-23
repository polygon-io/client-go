package models

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/polygon-io/client-go/rest/client"
)

const (
	ListQuotesPath = "/v3/quotes/{ticker}"
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
	Results []*Quote `json:"results,omitempty"`
}

// ListQuotesParams is the set of path and query parameters that are used when requesting quotes via the ListQuotes method.
type ListQuotesParams struct {
	Ticker      string
	QueryParams ListQuotesQueryParams
}

// ListQuotesQueryParams is the set of query parameters that can be used when requesting quotes via the ListQuotes method.
type ListQuotesQueryParams struct {
	TimestampEQ  *string // todo: make these time.Time instead of strings
	TimestampLT  *string
	TimestampLTE *string
	TimestampGT  *string
	TimestampGTE *string

	Order *Order

	Limit *int

	Sort *Sort

	AfterPrimary   *string
	AfterSecondary *string

	Cursor *string
}

// Path maps the ListQuotes path parameters to their respective keys.
func (p ListQuotesParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}

// Query maps the ListQuotes query parameters to their respective keys.
func (p ListQuotesParams) Query() url.Values {
	q := url.Values{}

	if p.QueryParams.TimestampEQ != nil {
		q.Set("timestamp", *p.QueryParams.TimestampEQ)
	}
	if p.QueryParams.TimestampLT != nil {
		q.Set("timestamp.lt", *p.QueryParams.TimestampLT)
	}
	if p.QueryParams.TimestampLTE != nil {
		q.Set("timestamp.lte", *p.QueryParams.TimestampLTE)
	}
	if p.QueryParams.TimestampGT != nil {
		q.Set("timestamp.gt", *p.QueryParams.TimestampGT)
	}
	if p.QueryParams.TimestampGTE != nil {
		q.Set("timestamp.gte", *p.QueryParams.TimestampGTE)
	}

	if p.QueryParams.Order != nil {
		q.Set("order", string(*p.QueryParams.Order))
	}

	if p.QueryParams.Limit != nil {
		q.Set("limit", strconv.FormatInt(int64(*p.QueryParams.Limit), 10))
	}

	if p.QueryParams.Sort != nil {
		q.Set("sort", string(*p.QueryParams.Sort))
	}

	if p.QueryParams.AfterPrimary != nil {
		q.Set("ap", *p.QueryParams.AfterPrimary)
	}
	if p.QueryParams.AfterSecondary != nil {
		q.Set("as", *p.QueryParams.AfterSecondary)
	}

	if p.QueryParams.Cursor != nil {
		q.Set("cursor", *p.QueryParams.Cursor)
	}

	return q
}

// String returns a URL string that includes any path and query parameters that are set.
func (p ListQuotesParams) String() string {
	path := strings.ReplaceAll(ListQuotesPath, "{ticker}", url.PathEscape(p.Ticker))

	q := p.Query().Encode()
	if q != "" {
		path += "?" + q
	}

	return path
}
