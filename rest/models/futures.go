// models/futures.go
package models

// ListFuturesAggsParams defines parameters for the ListFuturesAggs endpoint.
type ListFuturesAggsParams struct {
	Ticker         string  `validate:"required" path:"ticker"`
	Resolution     string  `query:"resolution"`
	WindowStart    *Nanos  `query:"window_start"`
	WindowStartLT  *Nanos  `query:"window_start.lt"`
	WindowStartLTE *Nanos  `query:"window_start.lte"`
	WindowStartGT  *Nanos  `query:"window_start.gt"`
	WindowStartGTE *Nanos  `query:"window_start.gte"`
	Limit          *int    `query:"limit"`
	Sort           *string `query:"sort"`
}

func (p ListFuturesAggsParams) WithWindowStart(c Comparator, q Nanos) *ListFuturesAggsParams {
	switch c {
	case EQ:
		p.WindowStart = &q
	case LT:
		p.WindowStartLT = &q
	case LTE:
		p.WindowStartLTE = &q
	case GT:
		p.WindowStartGT = &q
	case GTE:
		p.WindowStartGTE = &q
	}
	return &p
}

func (p ListFuturesAggsParams) WithLimit(q int) *ListFuturesAggsParams {
	p.Limit = &q
	return &p
}

func (p ListFuturesAggsParams) WithSort(q string) *ListFuturesAggsParams {
	p.Sort = &q
	return &p
}

// ListFuturesAggsResponse defines the response for the ListFuturesAggs endpoint.
type ListFuturesAggsResponse struct {
	BaseResponse
	Results []FuturesAggregate `json:"results,omitempty"`
}

// FuturesAggregate represents an aggregate for a futures contract.
type FuturesAggregate struct {
	Close           float64 `json:"close,omitempty"`
	DollarVolume    float64 `json:"dollar_volume,omitempty"`
	High            float64 `json:"high,omitempty"`
	Low             float64 `json:"low,omitempty"`
	Open            float64 `json:"open,omitempty"`
	SessionEndDate  string  `json:"session_end_date,omitempty"`
	SettlementPrice float64 `json:"settlement_price,omitempty"`
	Ticker          string  `json:"ticker,omitempty"`
	Transactions    int64   `json:"transaction_count,omitempty"`
	Volume          int64   `json:"volume,omitempty"`
	WindowStart     Nanos   `json:"window_start,omitempty"`
}

// ListFuturesContractsParams defines parameters for the ListFuturesContracts endpoint.
type ListFuturesContractsParams struct {
	ProductCode    *string `query:"product_code"`
	FirstTradeDate *Date   `query:"first_trade_date"`
	LastTradeDate  *Date   `query:"last_trade_date"`
	AsOf           *Date   `query:"as_of"`
	Active         *string `query:"active"`
	Type           *string `query:"type"`
	Limit          *int    `query:"limit"`
	Sort           *string `query:"sort"`
}

// ListFuturesContractsResponse defines the response for the ListFuturesContracts endpoint.
type ListFuturesContractsResponse struct {
	BaseResponse
	Results []FuturesContract `json:"results,omitempty"`
}

// FuturesContract represents a futures contract.
type FuturesContract struct {
	Active             bool    `json:"active,omitempty"`
	AsOf               Date    `json:"as_of,omitempty"`
	Maturity           string  `json:"maturity,omitempty"`
	DaysToMaturity     int     `json:"days_to_maturity,omitempty"`
	FirstTradeDate     Date    `json:"first_trade_date,omitempty"`
	LastTradeDate      Date    `json:"last_trade_date,omitempty"`
	MaxOrderQuantity   int     `json:"max_order_quantity,omitempty"`
	MinOrderQuantity   int     `json:"min_order_quantity,omitempty"`
	Name               string  `json:"name,omitempty"`
	ProductCode        string  `json:"product_code,omitempty"`
	SettlementDate     Date    `json:"settlement_date,omitempty"`
	SettlementTickSize float64 `json:"settlement_tick_size,omitempty"`
	SpreadTickSize     float64 `json:"spread_tick_size,omitempty"`
	Ticker             string  `json:"ticker,omitempty"`
	TradeTickSize      float64 `json:"trade_tick_size,omitempty"`
	TradingVenue       string  `json:"trading_venue,omitempty"`
	Type               string  `json:"type,omitempty"`
}

// GetFuturesContractParams defines parameters for the GetFuturesContract endpoint.
type GetFuturesContractParams struct {
	Ticker string `validate:"required" path:"ticker"`
	AsOf   *Date  `query:"as_of"`
}

// GetFuturesContractResponse defines the response for the GetFuturesContract endpoint.
type GetFuturesContractResponse struct {
	BaseResponse
	Results FuturesContract `json:"results,omitempty"`
}

// ListFuturesMarketStatusesParams defines parameters for the ListFuturesMarketStatuses endpoint.
type ListFuturesMarketStatusesParams struct {
	ProductCodeAnyOf *string `query:"product_code.any_of"`
	ProductCode      *string `query:"product_code"`
	Limit            *int    `query:"limit"`
	Sort             *string `query:"sort"`
}

// ListFuturesMarketStatusesResponse defines the response for the ListFuturesMarketStatuses endpoint.
type ListFuturesMarketStatusesResponse struct {
	BaseResponse
	Results   []FuturesMarketStatus `json:"results,omitempty"`
	Timestamp string                `json:"timestamp,omitempty"`
}

// FuturesMarketStatus represents the market status for a futures product.
type FuturesMarketStatus struct {
	MarketStatus string `json:"market_status,omitempty"`
	ProductCode  string `json:"product_code,omitempty"`
	TradingVenue string `json:"trading_venue,omitempty"`
}

// ListFuturesProductsParams defines parameters for the ListFuturesProducts endpoint.
type ListFuturesProductsParams struct {
	Name          *string `query:"name"`
	AsOf          *Date   `query:"as_of"`
	TradingVenue  *string `query:"trading_venue"`
	Sector        *string `query:"sector"`
	SubSector     *string `query:"sub_sector"`
	AssetClass    *string `query:"asset_class"`
	AssetSubClass *string `query:"asset_sub_class"`
	Type          *string `query:"type"`
	Limit         *int    `query:"limit"`
	NameSearch    *string `query:"name.search"`
	Sort          *string `query:"sort"`
}

// ListFuturesProductsResponse defines the response for the ListFuturesProducts endpoint.
type ListFuturesProductsResponse struct {
	BaseResponse
	Results []FuturesProduct `json:"results,omitempty"`
}

// FuturesProduct represents a futures product.
type FuturesProduct struct {
	AsOf                   Date    `json:"as_of,omitempty"`
	AssetClass             string  `json:"asset_class,omitempty"`
	AssetSubClass          string  `json:"asset_sub_class,omitempty"`
	ClearingChannel        string  `json:"clearing_channel,omitempty"`
	LastUpdated            string  `json:"last_updated,omitempty"`
	Name                   string  `json:"name,omitempty"`
	PriceQuotation         string  `json:"price_quotation,omitempty"`
	ProductCode            string  `json:"product_code,omitempty"`
	Sector                 string  `json:"sector,omitempty"`
	SettlementCurrencyCode string  `json:"settlement_currency_code,omitempty"`
	SettlementMethod       string  `json:"settlement_method,omitempty"`
	SettlementType         string  `json:"settlement_type,omitempty"`
	SubSector              string  `json:"sub_sector,omitempty"`
	TradeCurrencyCode      string  `json:"trade_currency_code,omitempty"`
	TradingVenue           string  `json:"trading_venue,omitempty"`
	Type                   string  `json:"type,omitempty"`
	UnitOfMeasure          string  `json:"unit_of_measure,omitempty"`
	UnitOfMeasureQuantity  float64 `json:"unit_of_measure_quantity,omitempty"`
}

// GetFuturesProductParams defines parameters for the GetFuturesProduct endpoint.
type GetFuturesProductParams struct {
	ProductCode string  `validate:"required" path:"product_code"`
	Type        *string `query:"type"`
	AsOf        *Date   `query:"as_of"`
}

// GetFuturesProductResponse defines the response for the GetFuturesProduct endpoint.
type GetFuturesProductResponse struct {
	BaseResponse
	Results FuturesProduct `json:"results,omitempty"`
}

// ListFuturesProductSchedulesParams defines parameters for the ListFuturesProductSchedules endpoint.
type ListFuturesProductSchedulesParams struct {
	ProductCode       string  `validate:"required" path:"product_code"`
	SessionEndDate    *Date   `query:"session_end_date"`
	SessionEndDateLT  *Date   `query:"session_end_date.lt"`
	SessionEndDateLTE *Date   `query:"session_end_date.lte"`
	SessionEndDateGT  *Date   `query:"session_end_date.gt"`
	SessionEndDateGTE *Date   `query:"session_end_date.gte"`
	Limit             *int    `query:"limit"`
	Sort              *string `query:"sort"`
}

func (p ListFuturesProductSchedulesParams) WithSessionEndDate(c Comparator, q Date) *ListFuturesProductSchedulesParams {
	switch c {
	case EQ:
		p.SessionEndDate = &q
	case LT:
		p.SessionEndDateLT = &q
	case LTE:
		p.SessionEndDateLTE = &q
	case GT:
		p.SessionEndDateGT = &q
	case GTE:
		p.SessionEndDateGTE = &q
	}
	return &p
}

// ListFuturesProductSchedulesResponse defines the response for the ListFuturesProductSchedules endpoint.
type ListFuturesProductSchedulesResponse struct {
	BaseResponse
	Results []FuturesSchedule `json:"results,omitempty"`
}

// FuturesSchedule represents a trading schedule for a futures product.
type FuturesSchedule struct {
	ProductCode    string          `json:"product_code,omitempty"`
	ProductName    string          `json:"product_name,omitempty"`
	Schedule       []ScheduleEvent `json:"schedule,omitempty"`
	SessionEndDate Date            `json:"session_end_date,omitempty"`
	TradingVenue   string          `json:"trading_venue,omitempty"`
}

type ScheduleEvent struct {
	Event     string `json:"event,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

// ListFuturesQuotesParams defines parameters for the ListFuturesQuotes endpoint.
type ListFuturesQuotesParams struct {
	Ticker            string  `validate:"required" path:"ticker"`
	Timestamp         *Nanos  `query:"timestamp"`
	TimestampLT       *Nanos  `query:"timestamp.lt"`
	TimestampLTE      *Nanos  `query:"timestamp.lte"`
	TimestampGT       *Nanos  `query:"timestamp.gt"`
	TimestampGTE      *Nanos  `query:"timestamp.gte"`
	SessionEndDate    *string `query:"session_end_date"`
	SessionEndDateLT  *string `query:"session_end_date.lt"`
	SessionEndDateLTE *string `query:"session_end_date.lte"`
	SessionEndDateGT  *string `query:"session_end_date.gt"`
	SessionEndDateGTE *string `query:"session_end_date.gte"`
	Limit             *int    `query:"limit"`
	Sort              *string `query:"sort"`
}

func (p ListFuturesQuotesParams) WithTimestamp(c Comparator, q Nanos) *ListFuturesQuotesParams {
	switch c {
	case EQ:
		p.Timestamp = &q
	case LT:
		p.TimestampLT = &q
	case LTE:
		p.TimestampLTE = &q
	case GT:
		p.TimestampGT = &q
	case GTE:
		p.TimestampGTE = &q
	}
	return &p
}

func (p ListFuturesQuotesParams) WithSessionEndDate(c Comparator, q string) *ListFuturesQuotesParams {
	switch c {
	case EQ:
		p.SessionEndDate = &q
	case LT:
		p.SessionEndDateLT = &q
	case LTE:
		p.SessionEndDateLTE = &q
	case GT:
		p.SessionEndDateGT = &q
	case GTE:
		p.SessionEndDateGTE = &q
	}
	return &p
}

// ListFuturesQuotesResponse defines the response for the ListFuturesQuotes endpoint.
type ListFuturesQuotesResponse struct {
	BaseResponse
	Results []FuturesQuote `json:"results,omitempty"`
}

// FuturesQuote represents a quote for a futures contract.
type FuturesQuote struct {
	AskPrice       float64 `json:"ask_price,omitempty"`
	AskSize        float64 `json:"ask_size,omitempty"`
	AskTimestamp   Nanos   `json:"ask_timestamp,omitempty"`
	BidPrice       float64 `json:"bid_price,omitempty"`
	BidSize        float64 `json:"bid_size,omitempty"`
	BidTimestamp   Nanos   `json:"bid_timestamp,omitempty"`
	SessionEndDate string  `json:"session_end_date,omitempty"`
	Ticker         string  `json:"ticker,omitempty"`
	Timestamp      Nanos   `json:"timestamp,omitempty"`
}

// ListFuturesSchedulesParams defines parameters for the ListFuturesSchedules endpoint.
type ListFuturesSchedulesParams struct {
	SessionEndDate *Date   `query:"session_end_date"`
	TradingVenue   *string `query:"trading_venue"`
	Limit          *int    `query:"limit"`
	Sort           *string `query:"sort"`
}

// ListFuturesSchedulesResponse defines the response for the ListFuturesSchedules endpoint.
type ListFuturesSchedulesResponse struct {
	BaseResponse
	Results []FuturesSchedule `json:"results,omitempty"`
}

// ListFuturesTradesParams defines parameters for the ListFuturesTrades endpoint.
type ListFuturesTradesParams struct {
	Ticker            string  `validate:"required" path:"ticker"`
	Timestamp         *Nanos  `query:"timestamp"`
	TimestampLT       *Nanos  `query:"timestamp.lt"`
	TimestampLTE      *Nanos  `query:"timestamp.lte"`
	TimestampGT       *Nanos  `query:"timestamp.gt"`
	TimestampGTE      *Nanos  `query:"timestamp.gte"`
	SessionEndDate    *string `query:"session_end_date"`
	SessionEndDateLT  *string `query:"session_end_date.lt"`
	SessionEndDateLTE *string `query:"session_end_date.lte"`
	SessionEndDateGT  *string `query:"session_end_date.gt"`
	SessionEndDateGTE *string `query:"session_end_date.gte"`
	Limit             *int    `query:"limit"`
	Sort              *string `query:"sort"`
}

func (p ListFuturesTradesParams) WithTimestamp(c Comparator, q Nanos) *ListFuturesTradesParams {
	switch c {
	case EQ:
		p.Timestamp = &q
	case LT:
		p.TimestampLT = &q
	case LTE:
		p.TimestampLTE = &q
	case GT:
		p.TimestampGT = &q
	case GTE:
		p.TimestampGTE = &q
	}
	return &p
}

func (p ListFuturesTradesParams) WithSessionEndDate(c Comparator, q string) *ListFuturesTradesParams {
	switch c {
	case EQ:
		p.SessionEndDate = &q
	case LT:
		p.SessionEndDateLT = &q
	case LTE:
		p.SessionEndDateLTE = &q
	case GT:
		p.SessionEndDateGT = &q
	case GTE:
		p.SessionEndDateGTE = &q
	}
	return &p
}

// ListFuturesTradesResponse defines the response for the ListFuturesTrades endpoint.
type ListFuturesTradesResponse struct {
	BaseResponse
	Results []FuturesTrade `json:"results,omitempty"`
}

// FuturesTrade represents a trade for a futures contract.
type FuturesTrade struct {
	Price          float64 `json:"price,omitempty"`
	SessionEndDate string  `json:"session_end_date,omitempty"`
	Size           float64 `json:"size,omitempty"`
	Ticker         string  `json:"ticker,omitempty"`
	Timestamp      Nanos   `json:"timestamp,omitempty"`
}
