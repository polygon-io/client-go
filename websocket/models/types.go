package models

type MarketType string

const (
	Stocks  MarketType = "stocks"
	Options MarketType = "options"
	Forex   MarketType = "forex"
	Crypto  MarketType = "crypto"
)
