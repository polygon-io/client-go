package models

import (
	"encoding/json"
	"strconv"
	"time"
)

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
	TickerSymbol       Sort = "ticker"
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

// todo: godoc

type Date time.Time

func (d *Date) UnmarshalJSON(data []byte) error {
	unquoteData, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", unquoteData)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

type Millis time.Time

func (m *Millis) UnmarshalJSON(data []byte) error {
	d, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*m = Millis(time.UnixMilli(d))
	return nil
}

func (m *Millis) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*m).UnixMilli())
}

type Nanos time.Time

func (n *Nanos) UnmarshalJSON(data []byte) error {
	d, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	// Go Time package does not include a method to convert UnixNano to a time.
	timeNano := time.Unix(d/1_000_000_000, d%1_000_000_000)
	*n = Nanos(timeNano)
	return nil
}
