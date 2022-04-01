package models

import (
	"github.com/polygon-io/client-go/rest/client"
)

const (
	ListSnapshotAllTickersPath     = "/v2/snapshot/locale/{locale}/markets/{marketType}/tickers"
	GetSnapshotTickerPath          = "/v2/snapshot/locale/{locale}/markets/{marketType}/tickers/{ticker}"
	ListSnapshotGainersLosersPath  = "/v2/snapshot/locale/{locale}/markets/{marketType}/{direction}"
	GetSnapshotOptionContractPath  = "/v3/snapshot/options/{underlyingAsset}/{optionContract}"
	ListSnapshotTickerFullBookPath = "/v2/snapshot/locale/global/markets/crypto/tickers/{ticker}/book"
)

// TickerSnapshot is a collection of data for a ticker including the current minute, day, and previous day's aggregate, as well as the last trade and quote.
type TickerSnapshot struct {
	Day              DaySnapshot       `json:"day"`
	LastQuote        LastQuoteSnapshot `json:"lastQuote"`
	LastTrade        LastTradeSnapshot `json:"lastTrade"`
	Minute           MinuteSnapshot    `json:"min"`
	PrevDay          DaySnapshot       `json:"prevDay"`
	Ticker           string            `json:"ticker"`
	TodaysChange     float64           `json:"todaysChange"`
	TodaysChangePerc float64           `json:"todaysChangePerc"`
	Updated          int64             `json:"updated"`
}

type DaySnapshot struct {
	Close                 float64 `json:"c"`
	High                  float64 `json:"h"`
	Low                   float64 `json:"l"`
	Open                  float64 `json:"o"`
	Volume                float64 `json:"v"`
	VolumeWeightedAverage float64 `json:"vw"`
}

type LastQuoteSnapshot struct {
	AskPrice  float64 `json:"P"`
	BidPrice  float64 `json:"p"`
	AskSize   float64 `json:"S"`
	BidSize   float64 `json:"s"`
	Timestamp int64   `json:"t"`
}

type LastTradeSnapshot struct {
	Conditions []int   `json:"c,omitempty"`
	TradeID    string  `json:"i"`
	Price      float64 `json:"p"`
	Size       int     `json:"s"`
	Timestamp  int64   `json:"t"`
	ExchangeID int     `json:"x"`
}

type MinuteSnapshot struct {
	AccumulatedVolume     int     `json:"av"`
	Close                 float64 `json:"c"`
	High                  float64 `json:"h"`
	Low                   float64 `json:"l"`
	Open                  float64 `json:"o"`
	Volume                float64 `json:"v"`
	VolumeWeightedAverage float64 `json:"vw"`
}

// GetSnapshotTickerResponse is returned by the Snapshot - Ticker API. It contains a snapshot for the specified ticker.
type GetSnapshotTickerResponse struct {
	client.BaseResponse
	Snapshot TickerSnapshot `json:"ticker,omitempty"`
}

// GetSnapshotTickerParams is the set of path and query parameters that can be used when requesting a snapshot for a ticker through the GetSnapshotTicker method.
type GetSnapshotTickerParams struct {
	Locale     MarketLocale
	MarketType MarketType
	Ticker     string
}

// Path maps the input GetSnapshotTickerParams path parameters to their respective keys.
func (p GetSnapshotTickerParams) Path() map[string]string {
	return map[string]string{
		"locale":     string(p.Locale),
		"marketType": string(p.MarketType),
		"ticker":     p.Ticker,
	}
}

// ListSnapshotAllTickersResponse is returned by the Snapshot - All Tickers API. It contains a snapshot for all the tickers of a specified market type.
type ListSnapshotAllTickersResponse struct {
	client.BaseResponse
	Snapshots []TickerSnapshot `json:"ticker,omitempty"`
}

// ListSnapshotAllTickersParams is the set of path and query parameters that can be used when requesting the snapshots for tickers through the ListSnapshotAllTickers method.
type ListSnapshotAllTickersParams struct {
	Locale     MarketLocale
	MarketType MarketType
}

// Path maps the input ListSnapshotAllTickersParams path parameters to their respective keys.
func (p ListSnapshotAllTickersParams) Path() map[string]string {
	return map[string]string{
		"locale":     string(p.Locale),
		"marketType": string(p.MarketType),
	}
}

// ListSnapshotGainersLosersResponse is returned by the Snapshot - Gainers/Losers API. It contains a snapshot of the top gainers or losers of a specified market type.
type ListSnapshotGainersLosersResponse struct {
	client.BaseResponse
	Snapshots []TickerSnapshot `json:"ticker,omitempty"`
}

// ListSnapshotGainersLosersParams is the set of path and query parameters that can be used when requesting a snapshot for a ticker through the ListSnapshotGainersLosers method.
type ListSnapshotGainersLosersParams struct {
	Locale     MarketLocale
	MarketType MarketType
	Direction  Direction
}

// Path maps the input ListSnapshotAllTickersParams path parameters to their respective keys.
func (p ListSnapshotGainersLosersParams) Path() map[string]string {
	return map[string]string{
		"locale":     string(p.Locale),
		"marketType": string(p.MarketType),
		"direction":  string(p.Direction),
	}
}

type OptionContractSnapshot struct {
	BreakEvenPrice    float64                         `json:"break_even_price"`
	Day               DayOptionContractSnapshot       `json:"day"`
	Details           OptionDetails                   `json:"details"`
	Greeks            Greeks                          `json:"greeks"`
	ImpliedVolatility float64                         `json:"implied_volatility"`
	LastQuote         LastQuoteOptionContractSnapshot `json:"last_quote"`
	OpenInterest      float64                         `json:"open_interest"`
	UnderlyingAsset   UnderlyingAsset                 `json:"underlying_asset"`
}

type DayOptionContractSnapshot struct {
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"change_percent"`
	Close         float64 `json:"close"`
	High          float64 `json:"high"`
	LastUpdated   int64   `json:"last_updated"`
	Low           float64 `json:"low"`
	Open          float64 `json:"open"`
	PreviousClose float64 `json:"previous_close"`
	Volume        float64 `json:"volume"`
	VWAP          float64 `json:"vwap"`
}

type OptionDetails struct {
	ContractType      string  `json:"contract_type"`
	ExerciseStyle     string  `json:"exercise_style"`
	ExpirationDate    string  `json:"expiration_date"`
	SharesPerContract float64 `json:"shares_per_contract"`
	StrikePrice       float64 `json:"strike_price"`
	Ticker            string  `json:"ticker"`
}

type Greeks struct {
	Delta float64 `json:"delta"`
	Gamma float64 `json:"gamma"`
	Theta float64 `json:"theta"`
	Vega  float64 `json:"vega"`
}

type LastQuoteOptionContractSnapshot struct {
	Ask         float64 `json:"ask"`
	AskSize     float64 `json:"ask_size"`
	Bid         float64 `json:"bid"`
	BidSize     float64 `json:"bid_size"`
	LastUpdated int64   `json:"last_updated"`
	Midpoint    float64 `json:"midpoint"`
	Timeframe   string  `json:"timeframe"`
}

type UnderlyingAsset struct {
	ChangeToBreakEven float64 `json:"change_to_break_even"`
	LastUpdated       int64   `json:"last_updated"`
	Price             float64 `json:"price"`
	Ticker            string  `json:"ticker"`
	Timeframe         string  `json:"timeframe"`
}

// GetSnapshotOptionContractResponse is returned by the Snapshot - Option Contract. It contains a snapshot of an option contract for a stock equity.
type GetSnapshotOptionContractResponse struct {
	client.BaseResponse
	Results OptionContractSnapshot `json:"results,omitempty"`
}

// GetSnapshotOptionContractParams is the set of path and query parameters that can be used when requesting a snapshot for a ticker through the GetSnapshotOptionContract method.
type GetSnapshotOptionContractParams struct {
	UnderlyingAsset string
	OptionContract  string
}

// Path maps the input GetSnapshotOptionContractParams path parameters to their respective keys.
func (p GetSnapshotOptionContractParams) Path() map[string]string {
	return map[string]string{
		"underlyingAsset": p.UnderlyingAsset,
		"optionContract":  p.OptionContract,
	}
}

type SnapshotTickerFullBook struct {
	AskCount float64 `json:"askCount"`
	Asks     []Ask   `json:"asks"`
	BidCount float64 `json:"bidCount"`
	Bids     []Bid   `json:"bids"`
	Spread   float64 `json:"spread"`
	Ticker   string  `json:"ticker"`
	Updated  int64   `json:"updated"`
}

type OrderBookQuote struct {
	Price            float64        `json:"p"`
	ExchangeToShares map[string]int `json:"x"`
}

type Ask OrderBookQuote
type Bid OrderBookQuote

// ListSnapshotTickerFullBookResponse is returned by the Snapshot - Ticker Full Book (L2). It contains the current level 2 book of a single ticker. This is the combined book from all of the exchanges.
type ListSnapshotTickerFullBookResponse struct {
	client.BaseResponse
	Data SnapshotTickerFullBook `json:"data,omitempty"`
}

// ListSnapshotTickerFullBookParams is the set of path and query parameters that can be used when requesting the full book of a single ticker through the ListSnapshotTickerFullBook method.
type ListSnapshotTickerFullBookParams struct {
	Ticker string
}

// Path maps the input ListSnapshotTickerFullBookParams path parameters to their respective keys.
func (p ListSnapshotTickerFullBookParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}
