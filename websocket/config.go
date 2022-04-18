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

type Config struct {
	APIKey string
	Feed   Feed
	Market Market
	Log    Logger
}
