package models

const (
	GetMarketHolidaysPath = "/v1/marketstatus/upcoming"
)

// todo: this endpoint is unlikely to ever have params so should we delete this type?
// GetMarketHolidaysParams is the set of parameters for the GetMarketHolidays method.
type GetMarketHolidaysParams struct{}

// GetMarketHolidaysResponse is the response returned by the GetMarketHolidays method.
type GetMarketHolidaysResponse []MarketHoliday

// MarketHoliday represents a market holiday for a specific exchange.
type MarketHoliday struct {
	Exchange string `json:"exchange,omitempty"`
	Name     string `json:"name,omitempty"`
	Date     string `json:"date,omitempty"` // todo: "2006-01-02" format
	Status   string `json:"status,omitempty"`
	Open     string `json:"open,omitempty"`  // todo: "2006-01-02T00:00:00.000Z" format
	Close    string `json:"close,omitempty"` // todo: "2006-01-02T00:00:00.000Z" format
}
