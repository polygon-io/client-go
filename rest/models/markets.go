package models

const (
	GetMarketHolidaysPath = "/v1/marketstatus/upcoming"
	GetMarketStatusPath   = "/v1/marketstatus/now"
)

// GetMarketHolidaysResponse is the response returned by the GetMarketHolidays method.
type GetMarketHolidaysResponse []MarketHoliday

// GetMarketStatusResponse is the response returned by the GetMarketStatus method.
type GetMarketStatusResponse struct {
	AfterHours bool              `json:"afterHours"`
	Currencies map[string]string `json:"currencies,omitempty"`
	EarlyHours bool              `json:"earlyHours"`
	Exchanges  map[string]string `json:"exchanges,omitempty"`
	Market     string            `json:"market,omitempty"`
	ServerTime string            `json:"serverTime,omitempty"` // todo: "2006-01-02T00:00:00.000Z" format
}

// MarketHoliday represents a market holiday for a specific exchange.
type MarketHoliday struct {
	Exchange string `json:"exchange,omitempty"`
	Name     string `json:"name,omitempty"`
	Date     string `json:"date,omitempty"` // todo: "2006-01-02" format
	Status   string `json:"status,omitempty"`
	Open     string `json:"open,omitempty"`  // todo: "2006-01-02T00:00:00.000Z" format
	Close    string `json:"close,omitempty"` // todo: "2006-01-02T00:00:00.000Z" format
}
