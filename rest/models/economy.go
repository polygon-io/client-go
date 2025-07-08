package models

// ListTreasuryYieldsParams is the set of parameters for the ListTreasuryYields method.
type ListTreasuryYieldsParams struct {
	// Calendar date of the yield observation (YYYY-MM-DD).
	DateEQ  *string `query:"date"`
	DateLT  *string `query:"date.lt"`
	DateLTE *string `query:"date.lte"`
	DateGT  *string `query:"date.gt"`
	DateGTE *string `query:"date.gte"`

	// Sort field used for ordering. Default is date.
	Sort *Sort `query:"sort"`

	// Order results based on the sort field. Default is asc.
	Order *Order `query:"order"`

	// Limit the number of results returned, default is 100 and max is 50000.
	Limit *int `query:"limit"`
}

func (p ListTreasuryYieldsParams) WithDate(c Comparator, q string) *ListTreasuryYieldsParams {
	switch c {
	case EQ:
		p.DateEQ = &q
	case LT:
		p.DateLT = &q
	case LTE:
		p.DateLTE = &q
	case GT:
		p.DateGT = &q
	case GTE:
		p.DateGTE = &q
	}
	return &p
}

func (p ListTreasuryYieldsParams) WithSort(q Sort) *ListTreasuryYieldsParams {
	p.Sort = &q
	return &p
}

func (p ListTreasuryYieldsParams) WithOrder(q Order) *ListTreasuryYieldsParams {
	p.Order = &q
	return &p
}

func (p ListTreasuryYieldsParams) WithLimit(q int) *ListTreasuryYieldsParams {
	p.Limit = &q
	return &p
}

// ListTreasuryYieldsResponse is the response returned by the ListTreasuryYields method.
type ListTreasuryYieldsResponse struct {
	BaseResponse

	// An array of treasury yields that match your query.
	Results []TreasuryYield `json:"results,omitempty"`
}

// TreasuryYield contains treasury yield data for a specific date.
type TreasuryYield struct {
	Date        string   `json:"date,omitempty"`
	Yield1Month *float64 `json:"yield_1_month,omitempty"`
	Yield3Month *float64 `json:"yield_3_month,omitempty"`
	Yield6Month *float64 `json:"yield_6_month,omitempty"`
	Yield1Year  *float64 `json:"yield_1_year,omitempty"`
	Yield2Year  *float64 `json:"yield_2_year,omitempty"`
	Yield3Year  *float64 `json:"yield_3_year,omitempty"`
	Yield5Year  *float64 `json:"yield_5_year,omitempty"`
	Yield7Year  *float64 `json:"yield_7_year,omitempty"`
	Yield10Year *float64 `json:"yield_10_year,omitempty"`
	Yield20Year *float64 `json:"yield_20_year,omitempty"`
	Yield30Year *float64 `json:"yield_30_year,omitempty"`
}
