package models

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/polygon-io/client-go/rest/client"
)

const (
	ListTradesPath = "/v3/trades/{ticker}"
)

type Trade struct {
	Conditions           []int32 `json:"conditions,omitempty"`
	Correction           int     `json:"correction,omitempty"`
	Exchange             int     `json:"exchange,omitempty"`
	ID                   string  `json:"id,omitempty"`
	ParticipantTimestamp int64   `json:"participant_timestamp,omitempty"`
	Price                float64 `json:"price,omitempty"`
	SequenceNumber       int64   `json:"sequence_number,omitempty"`
	SipTimestamp         int64   `json:"sip_timestamp,omitempty"`
	Size                 float64 `json:"size,omitempty"`
	Tape                 int32   `json:"tape,omitempty"`
	TrfID                int     `json:"trf_id,omitempty"`
	TrfTimestamp         int64   `json:"trf_timestamp,omitempty"`
}

type TradesResponse struct {
	client.BaseResponse
	Results []*Trade `json:"results,omitempty"`
}

// ListTradesParams is the set of path and query parameters that are used when requesting trades via the ListTrades method.
type ListTradesParams struct {
	Ticker      string
	QueryParams ListTradesQueryParams
}

// ListTradesQueryParams is the set of query parameters that can be used when requesting trades via the ListTrades method.
type ListTradesQueryParams struct {
	TimestampEQ  *time.Time
	TimestampLT  *time.Time
	TimestampLTE *time.Time
	TimestampGT  *time.Time
	TimestampGTE *time.Time

	Order *Order

	Limit *int

	Sort *Sort

	AfterPrimary   *string
	AfterSecondary *string

	Cursor *string
}

// Path maps the ListTrades path parameters to their respective keys.
func (p ListTradesParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}

// Query maps the ListTrades query parameters to their respective keys.
func (p ListTradesParams) Query() url.Values {
	q := url.Values{}

	if p.QueryParams.TimestampEQ != nil {
		q.Set("timestamp", fmt.Sprint(p.QueryParams.TimestampEQ.UnixNano()))
	}

	if p.QueryParams.TimestampLT != nil {
		q.Set("timestamp.lt", fmt.Sprint(p.QueryParams.TimestampLT.UnixNano()))
	}

	if p.QueryParams.TimestampLTE != nil {
		q.Set("timestamp.lte", fmt.Sprint(p.QueryParams.TimestampLTE.UnixNano()))
	}

	if p.QueryParams.TimestampGT != nil {
		q.Set("timestamp.gt", fmt.Sprint(p.QueryParams.TimestampGT.UnixNano()))
	}

	if p.QueryParams.TimestampGTE != nil {
		q.Set("timestamp.gte", fmt.Sprint(p.QueryParams.TimestampGTE.UnixNano()))
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
func (p ListTradesParams) String() string {
	path := strings.ReplaceAll(ListTradesPath, "{ticker}", url.PathEscape(p.Ticker))

	q := p.Query().Encode()
	if q != "" {
		path += "?" + q
	}

	return path
}
