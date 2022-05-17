package models

// GetOptionsContract is the set of parameters for the GetOptionsContract method.
type GetOptionsContractParams struct {
	// Return the contract that contains this options ticker.
	OptionsTicker string `validate:"required" path:"options_ticker"`

	// Specify a point in time for the contract as of this date.
	AsOf *Date `query:"as_of"`
}

func (p GetOptionsContractParams) WithAsOf(q Date) *GetOptionsContractParams {
	p.AsOf = &q
	return &p
}

// GetOptionsContractResponse is the response returned by the GetOptionsContract method.
type GetOptionsContractResponse struct {
	BaseResponse
	Results OptionsContract `json:"results,omitempty"`
}

// ListOptionsContracts is the set of parameters for the ListOptionsContracts method.
type ListOptionsContractsParams struct {
	// Return contracts relating to this underlying stock ticker.
	UnderlyingTickerEQ  *string `query:"underlying_ticker"`
	UnderlyingTickerLT  *string `query:"underlying_ticker.lt"`
	UnderlyingTickerLTE *string `query:"underlying_ticker.lte"`
	UnderlyingTickerGT  *string `query:"underlying_ticker.GT"`
	UnderlyingTickerGTE *string `query:"underlying_ticker.GTE"`

	// Specify the type of contract.
	ContractType *string `query:"contract_type"`

	// Specify the expiration date.
	ExpirationDateEQ  *Date `query:"expiration_date"`
	ExpirationDateLT  *Date `query:"expiration_date.lt"`
	ExpirationDateLTE *Date `query:"expiration_date.lte"`
	ExpirationDateGT  *Date `query:"expiration_date.GT"`
	ExpirationDateGTE *Date `query:"expiration_date.GTE"`

	// Specify the strike price.
	StrikePriceEQ  *float64 `query:"strike_price"`
	StrikePriceLT  *float64 `query:"strike_price.lt"`
	StrikePriceLTE *float64 `query:"strike_price.lte"`
	StrikePriceGT  *float64 `query:"strike_price.GT"`
	StrikePriceGTE *float64 `query:"strike_price.GTE"`

	// Specify whether to query for expired contracts.
	Expired *bool `query:"expired"`

	// Sort field used for ordering.
	Sort *Sort `query:"sort"`

	// Order results based on the sort field.
	Order *Order `query:"order"`

	// Limit the number of results returned, default is 10 and max is 1000.
	Limit *int `query:"limit"`
}

func (p ListOptionsContractsParams) WithUnderlyingTicker(c Comparator, q string) *ListOptionsContractsParams {
	if c == EQ {
		p.UnderlyingTickerEQ = &q
	} else if c == LT {
		p.UnderlyingTickerLT = &q
	} else if c == LTE {
		p.UnderlyingTickerLTE = &q
	} else if c == GT {
		p.UnderlyingTickerGT = &q
	} else if c == GTE {
		p.UnderlyingTickerGTE = &q
	}
	return &p
}

func (p ListOptionsContractsParams) WithContractType(q string) *ListOptionsContractsParams {
	p.ContractType = &q
	return &p
}

func (p ListOptionsContractsParams) WithExpirationDate(c Comparator, q Date) *ListOptionsContractsParams {
	if c == EQ {
		p.ExpirationDateEQ = &q
	} else if c == LT {
		p.ExpirationDateLT = &q
	} else if c == LTE {
		p.ExpirationDateLTE = &q
	} else if c == GT {
		p.ExpirationDateGT = &q
	} else if c == GTE {
		p.ExpirationDateGTE = &q
	}
	return &p
}

func (p ListOptionsContractsParams) WithStrikePrice(c Comparator, q float64) *ListOptionsContractsParams {
	if c == EQ {
		p.StrikePriceEQ = &q
	} else if c == LT {
		p.StrikePriceLT = &q
	} else if c == LTE {
		p.StrikePriceLTE = &q
	} else if c == GT {
		p.StrikePriceGT = &q
	} else if c == GTE {
		p.StrikePriceGTE = &q
	}
	return &p
}

func (p ListOptionsContractsParams) WithExpired(q bool) *ListOptionsContractsParams {
	p.Expired = &q
	return &p
}

func (p ListOptionsContractsParams) WithSort(q Sort) *ListOptionsContractsParams {
	p.Sort = &q
	return &p
}

func (p ListOptionsContractsParams) WithOrder(q Order) *ListOptionsContractsParams {
	p.Order = &q
	return &p
}

func (p ListOptionsContractsParams) WithLimit(q int) *ListOptionsContractsParams {
	p.Limit = &q
	return &p
}

type ListOptionsContractsResponse struct {
	BaseResponse
	Results []OptionsContract `json:"results,omitempty"`
}

// OptionsContract contains detailed information on a specified options contract.
type OptionsContract struct {
	AdditionalUnderlyings []Underlying `json:"additional_underlyings,omitempty"`
	CFI                   string       `json:"cfi,omitempty"`
	ContractType          string       `json:"contract_type,omitempty"`
	Correction            int32        `json:"correction,omitempty"`
	ExerciseStyle         string       `json:"exercise_style,omitempty"`
	ExpirationDate        Date         `json:"expiration_date,omitempty"`
	PrimaryExchange       string       `json:"primary_exchange,omitempty"`
	SharesPerContract     float64      `json:"shares_per_contract,omitempty"`
	StrikePrice           float64      `json:"strike_price,omitempty"`
	Ticker                string       `json:"ticker,omitempty"`
	UnderlyingTicker      string       `json:"underlying_ticker,omitempty"`
}

// An underlying or deliverable associated with an option contract.
type Underlying struct {
	Amount     float64 `json:"amount,omitempty"`
	Type       string  `json:"type,omitempty"`
	Underlying string  `json:"underlying,omitempty"`
}
