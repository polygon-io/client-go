package models

// ListTradesParams is the set of parameters for the ListTrades method.
type ListTradesParams struct {
	Ticker string `validate:"required" path:"ticker"`

	TimestampEQ  *Nanos `query:"timestamp"`
	TimestampLT  *Nanos `query:"timestamp.lt"`
	TimestampLTE *Nanos `query:"timestamp.lte"`
	TimestampGT  *Nanos `query:"timestamp.gt"`
	TimestampGTE *Nanos `query:"timestamp.gte"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
}

func (p ListTradesParams) WithTimestamp(c Comparator, q Nanos) *ListTradesParams {
	if c == EQ {
		p.TimestampEQ = &q
	} else if c == LT {
		p.TimestampLT = &q
	} else if c == LTE {
		p.TimestampLTE = &q
	} else if c == GT {
		p.TimestampGT = &q
	} else if c == GTE {
		p.TimestampGTE = &q
	}
	return &p
}

func (p ListTradesParams) WithSort(q Sort) *ListTradesParams {
	p.Sort = &q
	return &p
}

func (p ListTradesParams) WithOrder(q Order) *ListTradesParams {
	p.Order = &q
	return &p
}

func (p ListTradesParams) WithLimit(q int) *ListTradesParams {
	p.Limit = &q
	return &p
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
	ParticipantTimestamp Nanos   `json:"participant_timestamp,omitempty"`
	Price                float64 `json:"price,omitempty"`
	SequenceNumber       int64   `json:"sequence_number,omitempty"`
	SipTimestamp         Nanos   `json:"sip_timestamp,omitempty"`
	Size                 float64 `json:"size,omitempty"`
	Tape                 int32   `json:"tape,omitempty"`
	TrfID                int     `json:"trf_id,omitempty"`
	TrfTimestamp         Nanos   `json:"trf_timestamp,omitempty"`
}

// LastTrade is the most recent trade for a specified ticker.
type LastTrade struct {
	Ticker               string  `json:"T,omitempty"`
	TRFTimestamp         Nanos   `json:"f,omitempty"`
	SequenceNumber       int64   `json:"q,omitempty"`
	Timestamp            Nanos   `json:"t,omitempty"`
	ParticipantTimestamp Nanos   `json:"y,omitempty"`
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
	Timestamp  *Nanos  `json:"timestamp,omitempty"`
}
