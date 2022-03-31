package models

// Ptr returns a pointer to any value.
func Ptr[T any](v T) *T {
	return &v
}

// MarketType is the type of market.
type MarketType string

const (
	Stocks MarketType = "stocks"
	Forex  MarketType = "forex"
	Crypto MarketType = "crypto"
)

// Locale is the market location.
type MarketLocale string

const (
	US     MarketLocale = "us"
	Global MarketLocale = "global"
)

// Resolution is the size of the time window.
type Resolution string

const (
	Minute  Resolution = "minute"
	Hour    Resolution = "hour"
	Day     Resolution = "day"
	Week    Resolution = "week"
	Month   Resolution = "month"
	Quarter Resolution = "quarter"
	Year    Resolution = "year"
)

// Sort is a query param type that specifies how the results should be sorted.
type Sort string

const (
	Ticker             Sort = "ticker"
	Name               Sort = "name"
	Market             Sort = "market"
	Locale             Sort = "locale"
	PrimaryExchange    Sort = "primary_exchange"
	Type               Sort = "type"
	CurrencySymbol     Sort = "currency_symbol"
	CurrencyName       Sort = "currency_name"
	BaseCurrencySymbol Sort = "base_currency_symbol"
	BaseCurrencyName   Sort = "base_currency_name"
	CIK                Sort = "cik"
	CompositeFIGI      Sort = "composite_figi"
	ShareClassFIGI     Sort = "share_class_figi"
	LastUpdatedUTC     Sort = "last_updated_utc"
	DelistedUTC        Sort = "delisted_utc"
	Timestamp          Sort = "timestamp"
)

// Order the results. asc will return results in ascending order (oldest at the top),
// desc will return results in descending order (newest at the top).
type Order string

const (
	Asc  Order = "asc"
	Desc Order = "desc"
)

// Direction is the direction of the snapshot results to return
type Direction string

const (
	Gainers Direction = "gainers"
	Losers  Direction = "losers"
)
