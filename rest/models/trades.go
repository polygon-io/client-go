package models

import (
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
	Results []Trade `json:"results,omitempty"`
}

// ListTradesParams is the set of path and query parameters that are used when requesting trades via the ListTrades method.
type ListTradesParams struct {
	Ticker string `validate:"required"`

	TimestampEQ  *time.Time `query:"timestamp"`
	TimestampLT  *time.Time `query:"timestamp.lt"`
	TimestampLTE *time.Time `query:"timestamp.lte"`
	TimestampGT  *time.Time `query:"timestamp.gt"`
	TimestampGTE *time.Time `query:"timestamp.gte"`
	Limit        *int       `query:"limit"`
	Sort         *Sort      `query:"sort"`
	Order        *Order     `query:"order"`
}

// Path returns a map of URL path parameters.
func (p ListTradesParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
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

// Path returns a map of URL path parameters.
func (p GetLastTradeParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
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

// Path returns a map of URL path parameters.
func (p LastCryptoTradeParams) Path() map[string]string {
	return map[string]string{
		"from": p.From,
		"to":   p.To,
	}
}
