package models

// GetAllTickersSnapshotParams is the set of parameters for the GetAllTickersSnapshot method.
type GetAllTickersSnapshotParams struct {
	// The locale of the market.
	Locale MarketLocale `validate:"required" path:"locale"`

	// The type of market to query.
	MarketType MarketType `validate:"required" path:"marketType"`

	// A comma separated list of tickers to get snapshots for.
	Tickers *string `query:"tickers"`

	// Include OTC securities in the response. Default is false (don't include OTC securities).
	IncludeOTC *bool `query:"include_otc"`
}

func (p GetAllTickersSnapshotParams) WithTickers(q string) *GetAllTickersSnapshotParams {
	p.Tickers = &q
	return &p
}

func (p GetAllTickersSnapshotParams) WithIncludeOTC(q bool) *GetAllTickersSnapshotParams {
	p.IncludeOTC = &q
	return &p
}

// GetAllTickersSnapshotResponse is the response returned by the GetAllTickersSnapshot method.
type GetAllTickersSnapshotResponse struct {
	BaseResponse
	Tickers []TickerSnapshot `json:"tickers,omitempty"`
}

// GetTickerSnapshotParams is the set of parameters for the GetTickerSnapshot method.
type GetTickerSnapshotParams struct {
	// The locale of the market.
	Locale MarketLocale `validate:"required" path:"locale"`

	// The type of market to query.
	MarketType MarketType `validate:"required" path:"marketType"`

	// The ticker symbol of the stock/equity.
	Ticker string `validate:"required" path:"ticker"`
}

// GetTickerSnapshotResponse is the response returned by the GetTickerSnapshot method.
type GetTickerSnapshotResponse struct {
	BaseResponse
	Snapshot TickerSnapshot `json:"ticker,omitempty"`
}

// GetGainersLosersSnapshotParams is the set of parameters for the GetGainersLosersSnapshot method.
type GetGainersLosersSnapshotParams struct {
	// The locale of the market.
	Locale MarketLocale `validate:"required" path:"locale"`

	// The type of market to query.
	MarketType MarketType `validate:"required" path:"marketType"`

	// The direction of the snapshot results to return.
	Direction Direction `validate:"required" path:"direction"`

	// Include OTC securities in the response. Default is false (don't include OTC securities).
	IncludeOTC *bool `query:"include_otc"`
}

func (p GetGainersLosersSnapshotParams) WithIncludeOTC(q bool) *GetGainersLosersSnapshotParams {
	p.IncludeOTC = &q
	return &p
}

// GetOptionsChainParams is a set of parameters for the GetOptionsChainSnapshot method.
type GetOptionsChainParams struct {
	// The ticker for the option contract.
	Ticker *string `validate:"required" path:"underlyingAsset"`

	// The strike price of the option contract.
	StrikePrice *string `query:"string_price"`

	// The type of contract. Can be "put", "call", or in some rare cases, "other".
	ContractType *ContractType `query:"contract_type"`

	// The contract's expiration date in YYYY-MM-DD format.
	ExpirationDateEQ  *string `query:"expiration_date"`
	ExpirationDateLT  *string `query:"expiration_date.lt"`
	ExpirationDateLTE *string `query:"expiration_date.lte"`
	ExpirationDateGT  *string `query:"expiration_date.gt"`
	ExpirationDateGTE *string `query:"expiration_date.gte"`

	Limit *int   `query:"Limit"`
	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
}

// NewGetOptionsChainParams will return a pointer to GetOptionsChainParams.
func NewGetOptionsChainParams(ticker string) *GetOptionsChainParams {
	return &GetOptionsChainParams{Ticker: &ticker}
}

// WithTicker sets underlying ticker symbol to params
func (o *GetOptionsChainParams) WithTicker(ticker string) *GetOptionsChainParams {
	o.Ticker = &ticker
	return o
}

// WithStrikePrice sets strike price to params. Strike Price is the price at which a put or call option can be exercised.
func (o *GetOptionsChainParams) WithStrikePrice(strikePrice string) *GetOptionsChainParams {
	o.StrikePrice = &strikePrice
	return o
}

// WithContractType sets contract type to params.
// contractType options include `call` and `put`.
func (o *GetOptionsChainParams) WithContractType(contractType ContractType) *GetOptionsChainParams {
	o.ContractType = &contractType
	return o
}

// WithLimit sets number of results returned. Limit default is 10. Limit must fall in range of 0-1000.
func (o *GetOptionsChainParams) WithLimit(limit int) *GetOptionsChainParams {
	o.Limit = &limit
	return o
}

// WithExpirationDate sets expiration_date query parameter.
// comparator options include EQ, LT, LTE, GT, and GTE.
// expirationDate should be in YYYY-MM-DD format
func (o *GetOptionsChainParams) WithExpirationDate(comparator Comparator, expirationDate string) *GetOptionsChainParams {
	switch comparator {
	case EQ:
		o.ExpirationDateEQ = &expirationDate
	case LT:
		o.ExpirationDateLT = &expirationDate
	case LTE:
		o.ExpirationDateLTE = &expirationDate
	case GT:
		o.ExpirationDateGT = &expirationDate
	case GTE:
		o.ExpirationDateGTE = &expirationDate
	default:
		o.ExpirationDateEQ = &expirationDate
	}
	return o
}

// WithOrder sets order of results based on the Sort field.
func (o *GetOptionsChainParams) WithOrder(order Order) *GetOptionsChainParams {
	switch order {
	case Asc:
		o.Order = &order
	case Desc:
		o.Order = &order
	}
	return o
}

// WithSort sets sort field. Sort expects to receive TickerSymbol, ExpirationDate, or StrikePrice as an argument.
func (o *GetOptionsChainParams) WithSort(sort Sort) *GetOptionsChainParams {
	switch sort {
	case TickerSymbol:
		o.Sort = &sort
	case ExpirationDate:
		o.Sort = &sort
	case StrikePrice:
		o.Sort = &sort
	}
	return o
}

type GetOptionsChainSnapshotResponse struct {
	BaseResponse
	Results []OptionsChainSnapshot `json:"results,omitempty"`
}

type OptionsChainSnapshot struct {
	Day               DaySnapshot     `json:"day,omitempty"`
	Details           OptionDetails   `json:"details,omitempty"`
	Greeks            Greeks          `json:"greeks,omitempty"`
	LastQuote         LastQuote       `json:"last_quote,omitempty"`
	UnderlyingAsset   UnderlyingAsset `json:"underlying_asset,omitempty"`
	ImpliedVolatility *float64        `json:"implied_volatility,omitempty"`
	OpenInterest      *float64        `json:"open_interest,omitempty"`
	BreakEvenPrice    *float64        `json:"break_even_price,omitempty"`
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
	Data FullBookSnapshot `json:"data,omitempty"`
}

// TickerSnapshot is a collection of data for a ticker including the current minute, day, and previous day's aggregate,
// as well as the last trade and quote.
type TickerSnapshot struct {
	Day              DaySnapshot       `json:"day,omitempty"`
	LastQuote        LastQuoteSnapshot `json:"lastQuote,omitempty"`
	LastTrade        LastTradeSnapshot `json:"lastTrade,omitempty"`
	Minute           MinuteSnapshot    `json:"min,omitempty"`
	PrevDay          DaySnapshot       `json:"prevDay,omitempty"`
	Ticker           string            `json:"ticker,omitempty"`
	TodaysChange     float64           `json:"todaysChange,omitempty"`
	TodaysChangePerc float64           `json:"todaysChangePerc,omitempty"`
	Updated          Nanos             `json:"updated,omitempty"`
}

// DaySnapshot is the most recent day agg for a ticker.
type DaySnapshot struct {
	Close                 float64 `json:"c,omitempty"`
	High                  float64 `json:"h,omitempty"`
	Low                   float64 `json:"l,omitempty"`
	Open                  float64 `json:"o,omitempty"`
	Volume                float64 `json:"v,omitempty"`
	VolumeWeightedAverage float64 `json:"vw,omitempty"`
	OTC                   bool    `json:"otc,omitempty"`
}

// LastQuoteSnapshot is the most recent quote for a ticker.
type LastQuoteSnapshot struct {
	AskPrice  float64 `json:"P,omitempty"`
	BidPrice  float64 `json:"p,omitempty"`
	AskSize   float64 `json:"S,omitempty"`
	BidSize   float64 `json:"s,omitempty"`
	Timestamp Nanos   `json:"t,omitempty"`
}

// LastQuoteSnapshot is the most recent trade for a ticker.
type LastTradeSnapshot struct {
	Conditions []int   `json:"c,omitempty"`
	TradeID    string  `json:"i,omitempty"`
	Price      float64 `json:"p,omitempty"`
	Size       int     `json:"s,omitempty"`
	Timestamp  Nanos   `json:"t,omitempty"`
	ExchangeID int     `json:"x,omitempty"`
}

// DaySnapshot is the most recent minute agg for a ticker.
type MinuteSnapshot struct {
	AccumulatedVolume     int     `json:"av,omitempty"`
	Close                 float64 `json:"c,omitempty"`
	High                  float64 `json:"h,omitempty"`
	Low                   float64 `json:"l,omitempty"`
	Open                  float64 `json:"o,omitempty"`
	Volume                float64 `json:"v,omitempty"`
	VolumeWeightedAverage float64 `json:"vw,omitempty"`
	OTC                   bool    `json:"otc,omitempty"`
}

// OptionContractSnapshot is a collection of data for an option contract ticker including the current day aggregate and
// the most recent quote.
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

// DayOptionContractSnapshot contains the most recent day agg for an option contract.
type DayOptionContractSnapshot struct {
	Change        float64 `json:"change,omitempty"`
	ChangePercent float64 `json:"change_percent,omitempty"`
	Close         float64 `json:"close,omitempty"`
	High          float64 `json:"high,omitempty"`
	LastUpdated   Nanos   `json:"last_updated,omitempty"`
	Low           float64 `json:"low,omitempty"`
	Open          float64 `json:"open,omitempty"`
	PreviousClose float64 `json:"previous_close,omitempty"`
	Volume        float64 `json:"volume,omitempty"`
	VWAP          float64 `json:"vwap,omitempty"`
}

// OptionDetails contains more detailed information about an option contract.
type OptionDetails struct {
	ContractType      string  `json:"contract_type,omitempty"`
	ExerciseStyle     string  `json:"exercise_style,omitempty"`
	ExpirationDate    Date    `json:"expiration_date,omitempty"`
	SharesPerContract float64 `json:"shares_per_contract,omitempty"`
	StrikePrice       float64 `json:"strike_price,omitempty"`
	Ticker            string  `json:"ticker,omitempty"`
}

// Greeks contains the delta, gamma, vega, and theta of an option contract.
type Greeks struct {
	Delta float64 `json:"delta,omitempty"`
	Gamma float64 `json:"gamma,omitempty"`
	Theta float64 `json:"theta,omitempty"`
	Vega  float64 `json:"vega,omitempty"`
}

// LastQuoteOptionContractSnapshot contains the most recent quote of an option contract.
type LastQuoteOptionContractSnapshot struct {
	Ask         float64 `json:"ask,omitempty"`
	AskSize     float64 `json:"ask_size,omitempty"`
	Bid         float64 `json:"bid,omitempty"`
	BidSize     float64 `json:"bid_size,omitempty"`
	LastUpdated Nanos   `json:"last_updated,omitempty"`
	Midpoint    float64 `json:"midpoint,omitempty"`
	Timeframe   string  `json:"timeframe,omitempty"`
}

// UnderlyingAsset contains information on the underlying stock for this options contract.
type UnderlyingAsset struct {
	ChangeToBreakEven float64 `json:"change_to_break_even,omitempty"`
	LastUpdated       int64   `json:"last_updated,omitempty"`
	Price             float64 `json:"price,omitempty"`
	Ticker            string  `json:"ticker,omitempty"`
	Timeframe         string  `json:"timeframe,omitempty"`
}

// FullBookSnapshot is the level 2 book of a single crypto ticker.
type FullBookSnapshot struct {
	AskCount float64          `json:"askCount,omitempty"`
	Asks     []OrderBookQuote `json:"asks,omitempty"`
	BidCount float64          `json:"bidCount,omitempty"`
	Bids     []OrderBookQuote `json:"bids,omitempty"`
	Spread   float64          `json:"spread,omitempty"`
	Ticker   string           `json:"ticker,omitempty"`
	Updated  Nanos            `json:"updated,omitempty"`
}

// OrderBookQuote contains quote information for a crypto ticker.
type OrderBookQuote struct {
	Price            float64            `json:"p,omitempty"`
	ExchangeToShares map[string]float64 `json:"x,omitempty"`
}
