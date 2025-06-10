package models

import "strings"

type GetSummaryParams struct {
	// The ticker list to get summaries for
	TickerAnyOf *string `query:"ticker.any_of"`
}

func (p GetSummaryParams) WithTickerAnyOf(tickers ...string) *GetSummaryParams {
	q := strings.Join(tickers, ",")
	p.TickerAnyOf = &q
	return &p
}

type GetSummaryResponse struct {
	BaseResponse
	Results []SummaryResult `json:"results,omitempty"`
}

type SummaryResult struct {
	Price        float64  `json:"price,omitempty"`
	Name         string   `json:"name,omitempty"`
	Ticker       string   `json:"ticker,omitempty"`
	Branding     Branding `json:"branding,omitempty"`
	MarketStatus string   `json:"market_status,omitempty"`
	Type         string   `json:"type,omitempty"`
	Session      Session  `json:"session,omitempty"`
	Options      Options  `json:"options,omitempty"`
	Message      string   `json:"message,omitempty"`
	Error        string   `json:"error,omitempty"`
}

//easyjson:json
type Options struct {
	ContractType      string  `json:"contract_type,omitempty"`
	ExerciseStyle     string  `json:"exercise_style,omitempty"`
	ExpirationDate    Date    `json:"expiration_date,omitempty"`
	SharesPerContract float64 `json:"shares_per_contract,omitempty"`
	StrikePrice       float64 `json:"strike_price,omitempty"`
	UnderlyingTicker  string  `json:"underlying_ticker,omitempty"`
}
