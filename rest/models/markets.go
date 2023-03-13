package models

// GetMarketHolidaysResponse is the response returned by the GetMarketHolidays method.
type GetMarketHolidaysResponse []MarketHoliday

// GetMarketStatusResponse is the response returned by the GetMarketStatus method.
type GetMarketStatusResponse struct {
	AfterHours    bool              `json:"afterHours"`
	Currencies    map[string]string `json:"currencies,omitempty"`
	EarlyHours    bool              `json:"earlyHours"`
	Exchanges     map[string]string `json:"exchanges,omitempty"`
	IndicesGroups map[string]string `json:"indicesGroups,omitempty"`
	Market        string            `json:"market,omitempty"`
	ServerTime    Time              `json:"serverTime,omitempty"`
}

// MarketHoliday represents a market holiday for a specific exchange.
type MarketHoliday struct {
	Exchange string `json:"exchange,omitempty"`
	Name     string `json:"name,omitempty"`
	Date     Date   `json:"date,omitempty"`
	Status   string `json:"status,omitempty"`
	Open     Time   `json:"open,omitempty"`
	Close    Time   `json:"close,omitempty"`
}
