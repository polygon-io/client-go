package models

// ListStockFinancialsParams is the set of parameters for the ListFinancials method.
type ListStockFinancialsParams struct {
	Ticker *string `query:"ticker"`

	CIK *string `query:"cik"`

	CompanyNameFull   *string `query:"company_name"`
	CompanyNameSearch *string `query:"company_name.search"`

	SIC *string `query:"sic"`

	FilingDateEQ  *Date `query:"filing_dividend_date"`
	FilingDateLT  *Date `query:"filing_dividend_date.lt"`
	FilingDateLTE *Date `query:"filing_dividend_date.lte"`
	FilingDateGT  *Date `query:"filing_dividend_date.gt"`
	FilingDateGTE *Date `query:"filing_dividend_date.gte"`

	PeriodOfReportDateEQ  *Date `query:"period_of_report_date"`
	PeriodOfReportDateLT  *Date `query:"period_of_report_date.lt"`
	PeriodOfReportDateLTE *Date `query:"period_of_report_date.lte"`
	PeriodOfReportDateGT  *Date `query:"period_of_report_date.gt"`
	PeriodOfReportDateGTE *Date `query:"period_of_report_date.gte"`

	Timeframe *Timeframe `query:"timeframe"`

	IncludeSources *bool `query:"include_sources"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
}

func (p ListStockFinancialsParams) WithTicker(q string) *ListStockFinancialsParams {
	p.Ticker = &q
	return &p
}

func (p ListStockFinancialsParams) WithCIK(q string) *ListStockFinancialsParams {
	p.CIK = &q
	return &p
}

func (p ListStockFinancialsParams) WithCompanyName(c NameComparator, q string) *ListStockFinancialsParams {
	if c == Full {
		p.CompanyNameFull = &q
	} else if c == Search {
		p.CompanyNameSearch = &q
	}
	return &p
}

func (p ListStockFinancialsParams) WithSIC(q string) *ListStockFinancialsParams {
	p.SIC = &q
	return &p
}

func (p ListStockFinancialsParams) WithFilingDate(c Comparator, q Date) *ListStockFinancialsParams {
	if c == EQ {
		p.FilingDateEQ = &q
	} else if c == LT {
		p.FilingDateLT = &q
	} else if c == LTE {
		p.FilingDateLTE = &q
	} else if c == GT {
		p.FilingDateGT = &q
	} else if c == GTE {
		p.FilingDateGTE = &q
	}
	return &p
}

func (p ListStockFinancialsParams) WithPeriodOfReportDate(c Comparator, q Date) *ListStockFinancialsParams {
	if c == EQ {
		p.PeriodOfReportDateEQ = &q
	} else if c == LT {
		p.PeriodOfReportDateLT = &q
	} else if c == LTE {
		p.PeriodOfReportDateLTE = &q
	} else if c == GT {
		p.PeriodOfReportDateGT = &q
	} else if c == GTE {
		p.PeriodOfReportDateGTE = &q
	}
	return &p
}

func (p ListStockFinancialsParams) WithTimeframe(q Timeframe) *ListStockFinancialsParams {
	p.Timeframe = &q
	return &p
}

func (p ListStockFinancialsParams) WithIncludeSources(q bool) *ListStockFinancialsParams {
	p.IncludeSources = &q
	return &p
}

func (p ListStockFinancialsParams) WithSort(q Sort) *ListStockFinancialsParams {
	p.Sort = &q
	return &p
}

func (p ListStockFinancialsParams) WithOrder(q Order) *ListStockFinancialsParams {
	p.Order = &q
	return &p
}

func (p ListStockFinancialsParams) WithLimit(q int) *ListStockFinancialsParams {
	p.Limit = &q
	return &p
}

// ListStockFinancialsResponse is the response returned by the ListFinancials method.
type ListStockFinancialsResponse struct {
	BaseResponse
	Results []StockFinancial `json:"results,omitempty"`
}

// StockFinancial contains detailed information on a specified stock financial.
type StockFinancial struct {
	CIK                 string               `json:"cik,omitempty"`
	CompanyName         string               `json:"company_name,omitempty"`
	EndDate             string               `json:"end_date,omitempty"`
	FilingDate          string               `json:"filing_date,omitempty"`
	Financials          map[string]Financial `json:"financials,omitempty"`
	FiscalPeriod        string               `json:"fiscal_period,omitempty"`
	FiscalYear          string               `json:"fiscal_year,omitempty"`
	SourceFilingFileUrl string               `json:"source_filing_file_url,omitempty"`
	SourceFilingUrl     string               `json:"source_filing_url,omitempty"`
	StartDate           string               `json:"start_date,omitempty"`
}

type DataPoint struct {
	Formula string  `json:"formula,omitempty"`
	Label   string  `json:"label,omitempty"`
	Order   int32   `json:"order,omitempty"`
	Unit    string  `json:"unit,omitempty"`
	Value   float64 `json:"value,omitempty"`
	Xpath   string  `json:"xpath,omitempty"`
}

type Financial map[string]DataPoint
