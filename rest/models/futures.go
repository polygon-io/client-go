package models

import (
	"time"
)

// ListFuturesAggsParams defines parameters for listing futures aggregates.
type ListFuturesAggsParams struct {
	Ticker         string `validate:"required" path:"ticker"`
	Resolution     string `validate:"required" query:"resolution"`
	WindowStart    *Nanos `query:"window_start"`
	WindowStartGT  *Nanos `query:"window_start.gt"`
	WindowStartGTE *Nanos `query:"window_start.gte"`
	WindowStartLT  *Nanos `query:"window_start.lt"`
	WindowStartLTE *Nanos `query:"window_start.lte"`
	Order          *Order `query:"order"`
	Limit          *int   `query:"limit"`
	Sort           *Sort  `query:"sort"`
}

// ListFuturesContractsParams defines parameters for listing futures contracts.
type ListFuturesContractsParams struct {
	ProductCode    *string `query:"product_code"`
	FirstTradeDate *Date   `query:"first_trade_date"`
	LastTradeDate  *Date   `query:"last_trade_date"`
	ExpirationDate *Date   `query:"expiration_date"`
	Active         *string `query:"active"`
	Type           *string `query:"type"`
	Order          *Order  `query:"order"`
	Limit          *int    `query:"limit"`
	Sort           *Sort   `query:"sort"`
}

// GetFuturesContractParams defines parameters for retrieving a specific futures contract.
type GetFuturesContractParams struct {
	Ticker string `validate:"required" path:"ticker"`
	AsOf   *Date  `query:"as_of"`
}

// ListFuturesMarketStatusesParams defines parameters for listing market statuses.
type ListFuturesMarketStatusesParams struct {
	ProductCode  *string `query:"product_code"`
	ExchangeCode *string `query:"exchange_code"`
	Order        *Order  `query:"order"`
	Limit        *int    `query:"limit"`
	Sort         *Sort   `query:"sort"`
}

// ListFuturesProductsParams defines parameters for listing futures products.
type ListFuturesProductsParams struct {
	Name         *string `query:"name"`
	AssetClass   *string `query:"asset_class"`
	ExchangeCode *string `query:"exchange_code"`
	Sector       *string `query:"sector"`
	SubSector    *string `query:"sub_sector"`
	Type         *string `query:"type"`
	Order        *Order  `query:"order"`
	Limit        *int    `query:"limit"`
	Sort         *Sort   `query:"sort"`
}

// GetFuturesProductParams defines parameters for retrieving a specific futures product.
type GetFuturesProductParams struct {
	ProductCode string `validate:"required" path:"product_code"`
	AsOf        *Date  `query:"as_of"`
}

// ListFuturesSchedulesParams defines parameters for listing futures schedules.
type ListFuturesSchedulesParams struct {
	SessionStartDate     *Date   `query:"session_start_date"`
	MarketIdentifierCode *string `query:"market_identifier_code"`
	Order                *Order  `query:"order"`
	Limit                *int    `query:"limit"`
	Sort                 *Sort   `query:"sort"`
}

// ListFuturesProductSchedulesParams defines parameters for listing schedules for a specific product.
type ListFuturesProductSchedulesParams struct {
	ProductCode       string `validate:"required" path:"product_code"`
	SessionEndDate    *Date  `query:"session_end_date"`
	SessionEndDateGT  *Date  `query:"session_end_date.gt"`
	SessionEndDateGTE *Date  `query:"session_end_date.gte"`
	SessionEndDateLT  *Date  `query:"session_end_date.lt"`
	SessionEndDateLTE *Date  `query:"session_end_date.lte"`
	Order             *Order `query:"order"`
	Limit             *int   `query:"limit"`
}

// ListFuturesTradesParams defines parameters for listing futures trades.
type ListFuturesTradesParams struct {
	Ticker       string `validate:"required" path:"ticker"`
	Timestamp    *Nanos `query:"timestamp"`
	TimestampGT  *Nanos `query:"timestamp.gt"`
	TimestampGTE *Nanos `query:"timestamp.gte"`
	TimestampLT  *Nanos `query:"timestamp.lt"`
	TimestampLTE *Nanos `query:"timestamp.lte"`
	Order        *Order `query:"order"`
	Limit        *int   `query:"limit"`
	Sort         *Sort  `query:"sort"`
}

// ListFuturesQuotesParams defines parameters for listing futures quotes.
type ListFuturesQuotesParams struct {
	Ticker       string `validate:"required" path:"ticker"`
	Timestamp    *Nanos `query:"timestamp"`
	TimestampGT  *Nanos `query:"timestamp.gt"`
	TimestampGTE *Nanos `query:"timestamp.gte"`
	TimestampLT  *Nanos `query:"timestamp.lt"`
	TimestampLTE *Nanos `query:"timestamp.lte"`
	Order        *Order `query:"order"`
	Limit        *int   `query:"limit"`
	Sort         *Sort  `query:"sort"`
}

// FuturesAggregate represents a single aggregate bar for futures.
type FuturesAggregate struct {
	Ticker      string  `json:"ticker"`
	Open        float64 `json:"open"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	Close       float64 `json:"close"`
	Volume      int64   `json:"volume"`
	WindowStart int64   `json:"window_start"`
	WindowEnd   int64   `json:"window_end"`
}

// ListFuturesAggsResponse defines the response for listing futures aggregates.
type ListFuturesAggsResponse struct {
	BaseResponse
	Results []FuturesAggregate `json:"results,omitempty"`
}

// FuturesContract represents a futures contract.
type FuturesContract struct {
	Ticker         string `json:"ticker"`
	ProductCode    string `json:"product_code"`
	ExpirationDate Date   `json:"expiration_date"`
	FirstTradeDate Date   `json:"first_trade_date"`
	LastTradeDate  Date   `json:"last_trade_date"`
	Active         bool   `json:"active"`
	Type           string `json:"type"`
}

// ListFuturesContractsResponse defines the response for listing futures contracts.
type ListFuturesContractsResponse struct {
	BaseResponse
	Results []FuturesContract `json:"results,omitempty"`
}

// GetFuturesContractResponse defines the response for retrieving a specific futures contract.
type GetFuturesContractResponse struct {
	BaseResponse
	Result FuturesContract `json:"result,omitempty"`
}

// FuturesMarketStatus represents the market status for a futures product.
type FuturesMarketStatus struct {
	ProductCode  string `json:"product_code"`
	ExchangeCode string `json:"exchange_code"`
	MarketStatus string `json:"market_status"`
	Timestamp    int64  `json:"timestamp"`
}

// ListFuturesMarketStatusesResponse defines the response for listing market statuses.
type ListFuturesMarketStatusesResponse struct {
	BaseResponse
	Results []FuturesMarketStatus `json:"results,omitempty"`
}

// FuturesProduct represents a futures product.
type FuturesProduct struct {
	ProductCode  string `json:"product_code"`
	Name         string `json:"name"`
	AssetClass   string `json:"asset_class"`
	ExchangeCode string `json:"exchange_code"`
	Sector       string `json:"sector"`
	SubSector    string `json:"sub_sector"`
	Type         string `json:"type"`
}

// ListFuturesProductsResponse defines the response for listing futures products.
type ListFuturesProductsResponse struct {
	BaseResponse
	Results []FuturesProduct `json:"results,omitempty"`
}

// GetFuturesProductResponse defines the response for retrieving a specific futures product.
type GetFuturesProductResponse struct {
	BaseResponse
	Result FuturesProduct `json:"result,omitempty"`
}

// FuturesSchedule represents a trading schedule for futures.
type FuturesSchedule struct {
	MarketIdentifierCode string          `json:"market_identifier_code"`
	ProductCode          string          `json:"product_code"`
	ProductName          string          `json:"product_name"`
	SessionEndDate       Date            `json:"session_end_date"`
	Schedule             []ScheduleEvent `json:"schedule"`
}

// ScheduleEvent represents a single event in a schedule.
type ScheduleEvent struct {
	Event     string `json:"event"`
	Timestamp string `json:"timestamp"`
}

// ListFuturesSchedulesResponse defines the response for listing futures schedules.
type ListFuturesSchedulesResponse struct {
	BaseResponse
	Results []FuturesSchedule `json:"results,omitempty"`
}

// ListFuturesProductSchedulesResponse defines the response for listing schedules for a specific product.
type ListFuturesProductSchedulesResponse struct {
	BaseResponse
	Results []FuturesSchedule `json:"results,omitempty"`
}

// FuturesTrade represents a trade event for futures.
type FuturesTrade struct {
	Price     float64 `json:"price"`
	Size      float64 `json:"size"`
	Ticker    string  `json:"ticker"`
	Timestamp int64   `json:"timestamp"`
}

// ListFuturesTradesResponse defines the response for listing futures trades.
type ListFuturesTradesResponse struct {
	BaseResponse
	Results []FuturesTrade `json:"results,omitempty"`
}

// FuturesQuote represents a quote event for futures.
type FuturesQuote struct {
	AskPrice  float64 `json:"ask_price"`
	AskSize   float64 `json:"ask_size"`
	BidPrice  float64 `json:"bid_price"`
	BidSize   float64 `json:"bid_size"`
	Ticker    string  `json:"ticker"`
	Timestamp int64   `json:"timestamp"`
}

// ListFuturesQuotesResponse defines the response for listing futures quotes.
type ListFuturesQuotesResponse struct {
	BaseResponse
	Results []FuturesQuote `json:"results,omitempty"`
}
