package models

// MarketType is the type of market.
type MarketType string

const (
	Stocks MarketType = "stocks"
	Forex  MarketType = "forex"
	Crypto MarketType = "crypto"
)

// Locale is the market location.
type Locale string

const (
	US     Locale = "us"
	Global Locale = "global"
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

// Order the results. asc will return results in ascending order (oldest at the top),
// desc will return results in descending order (newest at the top).
type Order string

const (
	Asc  Order = "asc"
	Desc Order = "desc"
)

// Ptr returns a pointer to a specified value.
func Ptr[T any](v T) *T {
	return &v
}