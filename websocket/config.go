package polygonws

import "time"

// todo: add more config options with validation (waits, feed/market, etc)

const (
	writeWait      = 10 * time.Second
	pongWait       = 30 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1000000 // todo: 1MB, what's the limit on the server side?
)

type Logger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}

type nopLogger struct{}

func (l *nopLogger) Debugf(template string, args ...interface{}) {}
func (l *nopLogger) Infof(template string, args ...interface{})  {}
func (l *nopLogger) Errorf(template string, args ...interface{}) {}

type Config struct {
	APIKey string
	Feed   Feed
	Market Market
	Log    Logger
}

type Feed string

const (
	RealTime Feed = "socket"
	Delayed  Feed = "delayed"
	// todo: polyfeed, etc
)

type Market string

const (
	Stocks  Market = "stocks"
	Options Market = "options"
	Forex   Market = "forex"
	Crypto  Market = "crypto"
)

type Topic uint8

const (
	stocksMin        Topic = 10
	StocksSecAggs    Topic = 11
	StocksMinAggs    Topic = 12
	StocksTrades     Topic = 13
	StocksQuotes     Topic = 14
	StocksImbalances Topic = 15
	StocksLULD       Topic = 16
	stocksMax        Topic = 17

	optionsMin     Topic = 30
	OptionsSecAggs Topic = 31
	OptionsMinAggs Topic = 32
	OptionsTrades  Topic = 33
	OptionsQuotes  Topic = 34
	optionsMax     Topic = 35

	forexMin     Topic = 50
	ForexMinAggs Topic = 51
	ForexQuotes  Topic = 52
	forexMax     Topic = 53

	cryptoMin     Topic = 70
	CryptoMinAggs Topic = 71
	CryptoTrades  Topic = 72
	CryptoQuotes  Topic = 73
	CryptoL2Book  Topic = 74
	cryptoMax     Topic = 75
)

func (t Topic) prefix() string {
	switch t {
	case StocksSecAggs:
		return "A"
	case StocksMinAggs:
		return "AM"
	case StocksTrades:
		return "T"
	case StocksQuotes:
		return "Q"
	case StocksImbalances:
		return "NOI"
	case StocksLULD:
		return "LULD"
	case OptionsSecAggs:
		return "A"
	case OptionsMinAggs:
		return "AM"
	case OptionsTrades:
		return "T"
	case OptionsQuotes:
		return "Q"
	case ForexMinAggs:
		return "CA"
	case ForexQuotes:
		return "C"
	case CryptoMinAggs:
		return "XA"
	case CryptoTrades:
		return "XT"
	case CryptoQuotes:
		return "XQ"
	case CryptoL2Book:
		return "XL2"
	}

	return ""
}
