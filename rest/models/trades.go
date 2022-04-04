package models

import (
	"time"
)

const (
	ListTradesPath         = "/v3/trades/{ticker}"
	GetLastTradePath       = "/v2/last/trade/{ticker}"
	GetLastCryptoTradePath = "/v1/last/crypto/{from}/{to}"
)

// ListTradesParams is the set of parameters for the ListTrades method.
type ListTradesParams struct {
	Ticker string `validate:"required" path:"ticker"`

	TimestampEQ  *time.Time `query:"timestamp"`
	TimestampLT  *time.Time `query:"timestamp.lt"`
	TimestampLTE *time.Time `query:"timestamp.lte"`
	TimestampGT  *time.Time `query:"timestamp.gt"`
	TimestampGTE *time.Time `query:"timestamp.gte"`
	Limit        *int       `query:"limit"`
	Sort         *Sort      `query:"sort"`
	Order        *Order     `query:"order"`
}

// ListTradesResponse is the response returned by the ListTrades method.
type ListTradesResponse struct {
	BaseResponse
	Results []Trade `json:"results,omitempty"`
}

// GetLastTradeParams is the set of parameters for GetLastTrade method.
type GetLastTradeParams struct {
	Ticker string `validate:"required" path:"ticker"`
}

// GetLastTradeResponse is the response returned by the GetLastTradeResponse method.
type GetLastTradeResponse struct {
	BaseResponse
	Results LastTrade `json:"results,omitempty"`
}

// GetLastCryptoTradeParams is the set of parameters for the GetLastCryptoTrade method.
type GetLastCryptoTradeParams struct {
	From string `validate:"required" path:"from"`
	To   string `validate:"required" path:"to"`
}

// GetLastCryptoTradeResponse is the response returned by the GetLastCryptoTrade method.
type GetLastCryptoTradeResponse struct {
	BaseResponse
	Symbol string      `json:"symbol,omitempty"`
	Last   CryptoTrade `json:"last,omitempty"`
}

// Trade contains trade data for a specified ticker symbol.
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

// LastTrade is the most recent trade for a specified ticker.
type LastTrade struct {
	Ticker               string  `json:"T,omitempty"`
	TRFTimestamp         int64   `json:"f,omitempty"`
	SequenceNumber       int64   `json:"q,omitempty"`
	Timestamp            int64   `json:"t,omitempty"`
	ParticipantTimestamp int64   `json:"y,omitempty"`
	Conditions           []int32 `json:"c,omitempty"`
	Correction           uint32  `json:"e,omitempty"`
	ID                   string  `json:"i,omitempty"`
	Price                float64 `json:"p,omitempty"`
	TRF                  int32   `json:"r,omitempty"`
	Size                 uint32  `json:"s,omitempty"`
	Exchange             int32   `json:"x,omitempty"`
	Tape                 int32   `json:"z,omitempty"`
}

// CryptoTrade is a trade for a crypto pair.
type CryptoTrade struct {
	Price      float64 `json:"price,omitempty"`
	Size       float64 `json:"size,omitempty"`
	Exchange   int     `json:"exchange,omitempty"`
	Conditions []int   `json:"conditions,omitempty"`
	Timestamp  int64   `json:"timestamp,omitempty"`
}
