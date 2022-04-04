package models

const (
	GetAllTickersSnapshotPath     = "/v2/snapshot/locale/{locale}/markets/{marketType}/tickers"
	GetTickerSnapshotPath         = "/v2/snapshot/locale/{locale}/markets/{marketType}/tickers/{ticker}"
	GetGainersLosersSnapshotPath  = "/v2/snapshot/locale/{locale}/markets/{marketType}/{direction}"
	GetOptionContractSnapshotPath = "/v3/snapshot/options/{underlyingAsset}/{optionContract}"
	GetCryptoFullBookSnapshotPath = "/v2/snapshot/locale/global/markets/crypto/tickers/{ticker}/book"
)

// GetAllTickersSnapshotParams is the set of parameters for the GetAllTickersSnapshot method.
type GetAllTickersSnapshotParams struct {
	Locale     MarketLocale `validate:"required" path:"locale"`
	MarketType MarketType   `validate:"required" path:"marketType"`

	Tickers *string `query:"tickers"`
}

// GetAllTickersSnapshotResponse is the response returned by the GetAllTickersSnapshot method.
type GetAllTickersSnapshotResponse struct {
	BaseResponse
	Tickers []TickerSnapshot `json:"tickers,omitempty"`
}

// GetTickerSnapshotParams is the set of parameters for the GetTickerSnapshot method.
type GetTickerSnapshotParams struct {
	Locale     MarketLocale `validate:"required" path:"locale"`
	MarketType MarketType   `validate:"required" path:"marketType"`
	Ticker     string       `validate:"required" path:"ticker"`
}

// GetTickerSnapshotResponse is the response returned by the GetTickerSnapshot method.
type GetTickerSnapshotResponse struct {
	BaseResponse
	Snapshot TickerSnapshot `json:"ticker,omitempty"`
}

// GetGainersLosersSnapshotParams is the set of parameters for the GetGainersLosersSnapshot method.
type GetGainersLosersSnapshotParams struct {
	Locale     MarketLocale `validate:"required" path:"locale"`
	MarketType MarketType   `validate:"required" path:"marketType"`
	Direction  Direction    `validate:"required" path:"direction"`
}

// GetGainersLosersSnapshotResponse is the response returned by the GetGainersLosersSnapshot method.
type GetGainersLosersSnapshotResponse struct {
	BaseResponse
	Tickers []TickerSnapshot `json:"tickers,omitempty"`
}

// GetOptionContractSnapshotParams is the set of parameters for the GetOptionContractSnapshot method.
type GetOptionContractSnapshotParams struct {
	UnderlyingAsset string `validate:"required" path:"underlyingAsset"`
	OptionContract  string `validate:"required" path:"optionContract"`
}

// GetOptionContractSnapshotResponse is the response returned by the GetOptionContractSnapshot method.
type GetOptionContractSnapshotResponse struct {
	BaseResponse
	Results OptionContractSnapshot `json:"results,omitempty"`
}

// GetCryptoFullBookSnapshotParams is the set of parameters for the GetCryptoFullBookSnapshot method.
type GetCryptoFullBookSnapshotParams struct {
	Ticker string `validate:"required" path:"ticker"`
}

// GetCryptoFullBookSnapshotResponse is the response returned by the GetCryptoFullBookSnapshot method.
type GetCryptoFullBookSnapshotResponse struct {
	BaseResponse
	Data SnapshotTickerFullBook `json:"data,omitempty"`
}

// TickerSnapshot is a collection of data for a ticker including the current minute, day, and previous day's aggregate, as well as the last trade and quote.
type TickerSnapshot struct {
	Day              DaySnapshot       `json:"day,omitempty"`
	LastQuote        LastQuoteSnapshot `json:"lastQuote,omitempty"`
	LastTrade        LastTradeSnapshot `json:"lastTrade,omitempty"`
	Minute           MinuteSnapshot    `json:"min,omitempty"`
	PrevDay          DaySnapshot       `json:"prevDay,omitempty"`
	Ticker           string            `json:"ticker,omitempty"`
	TodaysChange     float64           `json:"todaysChange,omitempty"`
	TodaysChangePerc float64           `json:"todaysChangePerc,omitempty"`
	Updated          int64             `json:"updated,omitempty"`
}

type DaySnapshot struct {
	Close                 float64 `json:"c,omitempty"`
	High                  float64 `json:"h,omitempty"`
	Low                   float64 `json:"l,omitempty"`
	Open                  float64 `json:"o,omitempty"`
	Volume                float64 `json:"v,omitempty"`
	VolumeWeightedAverage float64 `json:"vw,omitempty"`
}

type LastQuoteSnapshot struct {
	AskPrice  float64 `json:"P,omitempty"`
	BidPrice  float64 `json:"p,omitempty"`
	AskSize   float64 `json:"S,omitempty"`
	BidSize   float64 `json:"s,omitempty"`
	Timestamp int64   `json:"t,omitempty"`
}

type LastTradeSnapshot struct {
	Conditions []int   `json:"c,omitempty"`
	TradeID    string  `json:"i,omitempty"`
	Price      float64 `json:"p,omitempty"`
	Size       int     `json:"s,omitempty"`
	Timestamp  int64   `json:"t,omitempty"`
	ExchangeID int     `json:"x,omitempty"`
}

type MinuteSnapshot struct {
	AccumulatedVolume     int     `json:"av,omitempty"`
	Close                 float64 `json:"c,omitempty"`
	High                  float64 `json:"h,omitempty"`
	Low                   float64 `json:"l,omitempty"`
	Open                  float64 `json:"o,omitempty"`
	Volume                float64 `json:"v,omitempty"`
	VolumeWeightedAverage float64 `json:"vw,omitempty"`
}

type OptionContractSnapshot struct {
	BreakEvenPrice    float64                         `json:"break_even_price,omitempty"`
	Day               DayOptionContractSnapshot       `json:"day,omitempty"`
	Details           OptionDetails                   `json:"details,omitempty"`
	Greeks            Greeks                          `json:"greeks,omitempty"`
	ImpliedVolatility float64                         `json:"implied_volatility,omitempty"`
	LastQuote         LastQuoteOptionContractSnapshot `json:"last_quote,omitempty"`
	OpenInterest      float64                         `json:"open_interest,omitempty"`
	UnderlyingAsset   UnderlyingAsset                 `json:"underlying_asset,omitempty"`
}

type DayOptionContractSnapshot struct {
	Change        float64 `json:"change,omitempty"`
	ChangePercent float64 `json:"change_percent,omitempty"`
	Close         float64 `json:"close,omitempty"`
	High          float64 `json:"high,omitempty"`
	LastUpdated   int64   `json:"last_updated,omitempty"`
	Low           float64 `json:"low,omitempty"`
	Open          float64 `json:"open,omitempty"`
	PreviousClose float64 `json:"previous_close,omitempty"`
	Volume        float64 `json:"volume,omitempty"`
	VWAP          float64 `json:"vwap,omitempty"`
}

type OptionDetails struct {
	ContractType      string  `json:"contract_type,omitempty"`
	ExerciseStyle     string  `json:"exercise_style,omitempty"`
	ExpirationDate    string  `json:"expiration_date,omitempty"`
	SharesPerContract float64 `json:"shares_per_contract,omitempty"`
	StrikePrice       float64 `json:"strike_price,omitempty"`
	Ticker            string  `json:"ticker,omitempty"`
}

type Greeks struct {
	Delta float64 `json:"delta,omitempty"`
	Gamma float64 `json:"gamma,omitempty"`
	Theta float64 `json:"theta,omitempty"`
	Vega  float64 `json:"vega,omitempty"`
}

type LastQuoteOptionContractSnapshot struct {
	Ask         float64 `json:"ask,omitempty"`
	AskSize     float64 `json:"ask_size,omitempty"`
	Bid         float64 `json:"bid,omitempty"`
	BidSize     float64 `json:"bid_size,omitempty"`
	LastUpdated int64   `json:"last_updated,omitempty"`
	Midpoint    float64 `json:"midpoint,omitempty"`
	Timeframe   string  `json:"timeframe,omitempty"`
}

type UnderlyingAsset struct {
	ChangeToBreakEven float64 `json:"change_to_break_even,omitempty"`
	LastUpdated       int64   `json:"last_updated,omitempty"`
	Price             float64 `json:"price,omitempty"`
	Ticker            string  `json:"ticker,omitempty"`
	Timeframe         string  `json:"timeframe,omitempty"`
}

type SnapshotTickerFullBook struct {
	AskCount float64          `json:"askCount,omitempty"`
	Asks     []OrderBookQuote `json:"asks,omitempty"`
	BidCount float64          `json:"bidCount,omitempty"`
	Bids     []OrderBookQuote `json:"bids,omitempty"`
	Spread   float64          `json:"spread,omitempty"`
	Ticker   string           `json:"ticker,omitempty"`
	Updated  int64            `json:"updated,omitempty"`
}

type OrderBookQuote struct {
	Price            float64            `json:"p,omitempty"`
	ExchangeToShares map[string]float64 `json:"x,omitempty"`
}
