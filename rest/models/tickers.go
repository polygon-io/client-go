package models

import "time"

const (
	ListTickersPath      = "/v3/reference/tickers"
	GetTickerDetailsPath = "/v3/reference/tickers/{ticker}"
	GetTickerTypesPath   = "/v3/reference/tickers/types"
)

// ListTickersParams is the set of parameters for the ListTickers method.
type ListTickersParams struct {
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	// todo: which ones should be enums?
	Type     *string     `query:"type"`
	Market   *MarketType `query:"market"` // todo: this endpoint apparently expects fx instead of forex
	Exchange *string     `query:"exchange"`
	CUSIP    *string     `query:"cusip"`
	CIK      *string     `query:"cik"`
	Date     *string     `query:"date"` // todo: this is "2006-01-02" format, need to figure out the best way to encode this without interfering with the default
	Active   *bool       `query:"active"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`

	PageMarker *string `query:"page_marker"`
	Search     *string `query:"search"`
}

// ListTickersResponse is the response returned by the ListTickers method.
type ListTickersResponse struct {
	BaseResponse
	Results []Ticker `json:"results,omitempty"`
}

// GetTickerDetailsParams is the set of parameters for the GetTickerDetails method.
type GetTickerDetailsParams struct {
	Ticker string `validate:"required" path:"ticker"`

	Date *string `query:"date"` // todo: this is "2006-01-02" format
}

// GetTickerDetailsResponse is the response returned by the GetTickerDetails method.
type GetTickerDetailsResponse struct {
	BaseResponse
	Results Ticker `json:"results,omitempty"`
}

// GetTickerTypesParams is the set of parameters for the GetTickerTypes method.
type GetTickerTypesParams struct {
	AssetClass *string       `query:"asset_class"` // todo: this is similar but slightly different than market type (also we offer four options but only one returns results)
	Locale     *MarketLocale `query:"locale"`
}

// GetTickerTypesResponse is the response returned by the GetTickerTypes method.
type GetTickerTypesResponse struct {
	BaseResponse
	Results []TickerType `json:"results,omitempty"`
}

// Ticker contains detailed information on a specified ticker symbol.
type Ticker struct {
	Ticker                      string          `json:"ticker,omitempty"`
	Name                        string          `json:"name,omitempty"`
	Market                      string          `json:"market,omitempty"`
	Locale                      string          `json:"locale,omitempty"`
	PrimaryExchange             string          `json:"primary_exchange,omitempty"`
	Type                        string          `json:"type,omitempty"`
	Active                      bool            `json:"active,omitempty"`
	CurrencySymbol              string          `json:"currency_symbol,omitempty"`
	CurrencyName                string          `json:"currency_name,omitempty"`
	BaseCurrencySymbol          string          `json:"base_currency_symbol,omitempty"`
	BaseCurrencyName            string          `json:"base_currency_name,omitempty"`
	CUSIP                       string          `json:"cusip,omitempty"`
	CIK                         string          `json:"cik,omitempty"`
	CompositeFIGI               string          `json:"composite_figi,omitempty"`
	ShareClassFIGI              string          `json:"share_class_figi,omitempty"`
	LastUpdatedUTC              time.Time       `json:"last_updated_utc,omitempty"`
	DelistedUTC                 int64           `json:"delisted_utc,omitempty"`
	MarketCap                   float64         `json:"market_cap,omitempty"`
	PhoneNumber                 string          `json:"phone_number,omitempty"`
	Address                     *CompanyAddress `json:"address,omitempty"`
	Description                 string          `json:"description,omitempty"`
	SICCode                     string          `json:"sic_code,omitempty"`
	SICDescription              string          `json:"sic_description,omitempty"`
	TickerRoot                  string          `json:"ticker_root,omitempty"`
	TickerSuffix                string          `json:"ticker_suffix,omitempty"`
	HomepageURL                 string          `json:"homepage_url,omitempty"`
	TotalEmployees              int32           `json:"total_employees,omitempty"`
	ListDate                    string          `json:"list_date,omitempty"`
	Branding                    *Branding       `json:"branding,omitempty"`
	ShareClassSharesOutstanding int64           `json:"share_class_shares_outstanding,omitempty"`
	WeightedSharesOutstanding   int64           `json:"weighted_shares_outstanding,omitempty"`
}

// CompanyAddress contains information on the physical address of a company.
type CompanyAddress struct {
	Address1   string `json:"address1,omitempty"`
	Address2   string `json:"address2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"country,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
}

// Branding contains information related to a company's brand.
type Branding struct {
	LogoURL     string `json:"logo_url,omitempty"`
	IconURL     string `json:"icon_url,omitempty"`
	AccentColor string `json:"accent_color,omitempty"`
	LightColor  string `json:"light_color,omitempty"`
	DarkColor   string `json:"dark_color,omitempty"`
}

// TickerType represents a type of ticker with a code that the API understands.
type TickerType struct {
	AssetClass  string `json:"asset_class,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Locale      string `json:"locale,omitempty"`
}
