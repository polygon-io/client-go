package models

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

// Agg is an aggregation of all the activity on a specified ticker between the start and end timestamps.
type Agg struct {
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

	// todo: these aren't listed in the docs
	Timestamp    int64  `json:"t,omitempty"`
	Transactions int64  `json:"n,omitempty"`
	Market       string `json:"m,omitempty"`
	Exchange     int32  `json:"x,omitempty"`
	Locale       string `json:"g,omitempty"`
}
