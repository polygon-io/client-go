package models

// todo: verify models before release

type Action string

const (
	Auth        Action = "auth"
	Subscribe   Action = "subscribe"
	Unsubscribe Action = "unsubscribe"
)

type EventType struct {
	EventType string `json:"ev,omitempty"`
}

type ControlMessage struct {
	EventType
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Action  Action `json:"action,omitempty"`
	Params  string `json:"params,omitempty"`
}

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
}

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
}

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

type CurrencyTrade struct {
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
