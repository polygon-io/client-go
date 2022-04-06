package models

const (
	GetMarketHolidaysPath = "/v1/marketstatus/upcoming"
	GetMarketStatusPath   = "/v1/marketstatus/now"
)

// todo: this endpoint is unlikely to ever have params so should we delete this type?
// GetMarketHolidaysParams is the set of parameters for the GetMarketHolidays method.
type GetMarketHolidaysParams struct{}

// GetMarketHolidaysResponse is the response returned by the GetMarketHolidays method.
type GetMarketHolidaysResponse []MarketHoliday

// todo: this endpoint is unlikely to ever have params so should we delete this type?
// GetMarketStatusParams is the set of parameters for the GetMarketStatus method.
type GetMarketStatusParams struct{}

// GetMarketStatusResponse is the response returned by the GetMarketStats method.
type GetMarketStatusResponse MarketStatus

// MarketHoliday represents a market holiday for a specific exchange.
type MarketHoliday struct {
	Exchange string `json:"exchange,omitempty"`
	Name     string `json:"name,omitempty"`
	Date     string `json:"date,omitempty"` // todo: "2006-01-02" format
	Status   string `json:"status,omitempty"`
	Open     string `json:"open,omitempty"`  // todo: "2006-01-02T00:00:00.000Z" format
	Close    string `json:"close,omitempty"` // todo: "2006-01-02T00:00:00.000Z" format
}

// MarketStatus represents the current trading status of the exchanges and overall financial markets.
type MarketStatus struct {
	AfterHours bool              `json:"afterHours,omitempty"`
	Currencies map[string]string `json:"currencies,omitempty"`
	EarlyHours bool              `json:"earlyHours,omitempty"`
	Exchanges  map[string]string `json:"exchanges,omitempty"`
	Market     string            `json:"market,omitempty"`
	ServerTime string            `json:"serverTime,omitempty"` // todo: "2006-01-02T00:00:00.000Z" format
}
