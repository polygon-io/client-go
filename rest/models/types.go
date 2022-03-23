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

// Sort is a query param type that specifies how the results should be sorted.
type Sort string

const (
	Timestamp Sort = "timestamp"
)

// Order the results. asc will return results in ascending order (oldest at the top),
// desc will return results in descending order (newest at the top).
type Order string

const (
	Asc  Order = "asc"
	Desc Order = "desc"
)

// todo: replace these with a single generic method once the linter supports it

// Bool returns a pointer to a bool value.
func Bool(v bool) *bool {
	return &v
}

// Int returns a pointer to an int value.
func Int(v int) *int {
	return &v
}

// String returns a pointer to a string value.
func String(v string) *string {
	return &v
}

// OrderBy returns a pointer to an order value.
func OrderBy(v Order) *Order {
	return &v
}

// SortOn returns a pointer to a sort value.
func SortOn(v Sort) *Sort {
	return &v
}
