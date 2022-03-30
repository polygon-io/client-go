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
	ListTradesPath         = "/v3/trades/{ticker}"
	GetLastTradePath       = "/v2/last/trade/{ticker}"
	GetLastCryptoTradePath = "/v1/last/crypto/{from}/{to}"
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

// LastTrade is the most recent trade for a specified ticker.
// For more details see https://polygon.io/docs/stocks/get_v2_last_trade__stocksticker.
type LastTrade struct {
	Ticker               string  `json:"T,omitempty"`
	TRFTimestamp         int64   `json:"f,omitempty"`
	SequenceNumber       int64   `json:"q"`
	Timestamp            int64   `json:"t"`
	ParticipantTimestamp int64   `json:"y,omitempty"`
	Conditions           []int32 `json:"c,omitempty"`
	Correction           uint32  `json:"e,omitempty"`
	ID                   string  `json:"i"`
	Price                float64 `json:"p"`
	TRF                  int32   `json:"r,omitempty"`
	Size                 uint32  `json:"s"`
	Exchange             int32   `json:"x"`
	Tape                 int32   `json:"z,omitempty"`
}

// LastTradeResponse contains the most recent trade for a specified ticker.
type LastTradeResponse struct {
	client.BaseResponse
	Results LastTrade `json:"results,omitempty"`
}

// GetLastTradeParams is the set of path and query parameters for retrieving the most recent trade for a specified ticker.
type GetLastTradeParams struct {
	Ticker string
}

// Path maps the path parameters to their respective keys.
func (p GetLastTradeParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}

// Query maps the query parameters to their respective keys.
func (p GetLastTradeParams) Query() url.Values {
	q := url.Values{}
	return q
}

// CryptoTrade is a trade for a crypto pair.
// For more details see https://polygon.io/docs/crypto/get_v1_last_crypto__from___to.
type CryptoTrade struct {
	Price      float64 `json:"price,omitempty"`
	Size       float64 `json:"size,omitempty"`
	Exchange   int     `json:"exchange,omitempty"`
	Conditions []int   `json:"conditions,omitempty"`
	Timestamp  int64   `json:"timestamp,omitempty"`
}

// LastCryptoTradeResponse contains the most recent trade for a crypto pair.
type LastCryptoTradeResponse struct {
	client.BaseResponse
	Last CryptoTrade `json:"last,omitempty"`
}

// LastCryptoTradeParams is the set of path and query parameters for retrieving the most recent trade for a crypto pair.
type LastCryptoTradeParams struct {
	From string
	To   string
}

// Path maps the path parameters to their respective keys.
func (p LastCryptoTradeParams) Path() map[string]string {
	return map[string]string{
		"from": p.From,
		"to":   p.To,
	}
}

// Query maps the query parameters to their respective keys.
func (p LastCryptoTradeParams) Query() url.Values {
	q := url.Values{}
	return q
}
