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

type SnapshotTickerFullBook struct {
	AskCount float64          `json:"askCount"`
	Asks     []OrderBookQuote `json:"asks"`
	BidCount float64          `json:"bidCount"`
	Bids     []OrderBookQuote `json:"bids"`
	Spread   float64          `json:"spread"`
	Ticker   string           `json:"ticker"`
	Updated  int64            `json:"updated"`
}

type OrderBookQuote struct {
	Price            float64            `json:"p"`
	ExchangeToShares map[string]float64 `json:"x"`
}
