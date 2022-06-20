package models

// Action is the set of recognized actions used in control messages.
type Action string

const (
	Auth        Action = "auth"
	Subscribe   Action = "subscribe"
	Unsubscribe Action = "unsubscribe"
)

// EventType is the type of message received. It should be present in
// every message sent by the server.
type EventType struct {
	EventType string `json:"ev,omitempty"`
}

// ControlMessage is a message to signal status and control events to
// and from the server.
type ControlMessage struct {
	EventType
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Action  Action `json:"action,omitempty"`
	Params  string `json:"params,omitempty"`
}

// EquityAgg is an aggregate for either stock tickers or option contracts.
type EquityAgg struct {
	EventType
	Symbol            string  `json:"sym,omitempty"`
	Volume            float64 `json:"v,omitempty"`
	AccumulatedVolume float64 `json:"av,omitempty"`
	OfficialOpenPrice float64 `json:"op,omitempty"`
	VWAP              float64 `json:"vw,omitempty"`
	Open              float64 `json:"o,omitempty"`
	Close             float64 `json:"c,omitempty"`
	High              float64 `json:"h,omitempty"`
	Low               float64 `json:"l,omitempty"`
	AggregateVWAP     float64 `json:"a,omitempty"`
	AverageSize       float64 `json:"z,omitempty"`
	StartTimestamp    int64   `json:"s,omitempty"`
	EndTimestamp      int64   `json:"e,omitempty"`
	OTC               bool    `json:"otc,omitempty"`
}

// CurrencyAgg is an aggregate for either forex currency pairs or crypto pairs.
type CurrencyAgg struct {
	EventType
	Pair           string  `json:"pair,omitempty"`
	Open           float64 `json:"o,omitempty"`
	Close          float64 `json:"c,omitempty"`
	High           float64 `json:"h,omitempty"`
	Low            float64 `json:"l,omitempty"`
	Volume         float64 `json:"v,omitempty"`
	VWAP           float64 `json:"vw,omitempty"`
	StartTimestamp int64   `json:"s,omitempty"`
	EndTimestamp   int64   `json:"e,omitempty"`
	AVGTradeSize   int32   `json:"z,omitempty"`
}

// EquityTrade is trade data for either stock tickers or option contracts.
type EquityTrade struct {
	EventType
	Symbol         string  `json:"sym,omitempty"`
	Exchange       int32   `json:"x,omitempty"`
	ID             string  `json:"i,omitempty"`
	Tape           int32   `json:"z,omitempty"`
	Price          float64 `json:"p,omitempty"`
	Size           int64   `json:"s,omitempty"`
	Conditions     []int32 `json:"c,omitempty"`
	Timestamp      int64   `json:"t,omitempty"`
	SequenceNumber int64   `json:"q,omitempty"`
}

// CryptoTrade is a trade for a crypto pair.
type CryptoTrade struct {
	EventType
	Symbol            string  `json:"sym,omitempty"`
	Exchange          int32   `json:"x,omitempty"`
	ID                string  `json:"i,omitempty"`
	Price             float64 `json:"p,omitempty"`
	Size              float64 `json:"s,omitempty"`
	Conditions        []int32 `json:"c,omitempty"`
	Timestamp         int64   `json:"t,omitempty"`
	ReceivedTimestamp int64   `json:"r,omitempty"`
}

// EquityQuote is a quote for either stock tickers or option contracts.
type EquityQuote struct {
	EventType
	Symbol         string  `json:"sym,omitempty"`
	BidExchangeID  int32   `json:"bx,omitempty"`
	BidPrice       float64 `json:"bp,omitempty"`
	BidSize        int32   `json:"bs,omitempty"`
	AskExchangeID  int32   `json:"ax,omitempty"`
	AskPrice       float64 `json:"ap,omitempty"`
	AskSize        int32   `json:"as,omitempty"`
	Condition      int32   `json:"c,omitempty"`
	Timestamp      int64   `json:"t,omitempty"`
	Tape           int32   `json:"z,omitempty"`
	SequenceNumber int64   `json:"q,omitempty"`
}

// ForexQuote is a quote for a forex currency pair.
type ForexQuote struct {
	EventType
	Pair       string  `json:"p,omitempty"`
	ExchangeID int32   `json:"x,omitempty"`
	AskPrice   float64 `json:"a,omitempty"`
	BidPrice   float64 `json:"b,omitempty"`
	Timestamp  int64   `json:"t,omitempty"`
}

// CryptoQuote is a quote for a crypto pair.
type CryptoQuote struct {
	EventType
	Pair              string  `json:"pair,omitempty"`
	BidPrice          float64 `json:"bp,omitempty"`
	BidSize           int32   `json:"bs,omitempty"`
	AskPrice          float64 `json:"ap,omitempty"`
	AskSize           int32   `json:"as,omitempty"`
	Timestamp         int64   `json:"t,omitempty"`
	ExchangeID        int32   `json:"x,omitempty"`
	ReceivedTimestamp int64   `json:"r,omitempty"`
}

// Imbalance is an imbalance event for a given stock ticker symbol.
type Imbalance struct {
	EventType
	Symbol            string  `json:"T,omitempty"`
	Timestamp         int64   `json:"t,omitempty"`
	AuctionTime       int32   `json:"at,omitempty"`
	AuctionType       string  `json:"a,omitempty"`
	SymbolSequence    int32   `json:"i,omitempty"`
	ExchangeID        int32   `json:"x,omitempty"`
	ImbalanceQuantity int32   `json:"o,omitempty"`
	PairedQuantity    int32   `json:"p,omitempty"`
	BookClearingPrice float64 `json:"b,omitempty"`
}

// LimitUpLimitDown is a LULD event for a given stock ticker symbol.
type LimitUpLimitDown struct {
	EventType
	Symbol         string  `json:"T,omitempty"`
	HighPrice      float64 `json:"h,omitempty"`
	LowPrice       float64 `json:"l,omitempty"`
	Indicators     []int32 `json:"i,omitempty"`
	Tape           int32   `json:"z,omitempty"`
	Timestamp      int64   `json:"t,omitempty"`
	SequenceNumber int64   `json:"q,omitempty"`
}

// Level2Book is level 2 book data for a given crypto pair.
type Level2Book struct {
	EventType
	Pair              string    `json:"pair,omitempty"`
	BidPrices         []float64 `json:"b,omitempty"`
	AskPrices         []float64 `json:"a,omitempty"`
	Timestamp         int64     `json:"t,omitempty"`
	ExchangeID        int32     `json:"x,omitempty"`
	ReceivedTimestamp int64     `json:"r,omitempty"`
}
