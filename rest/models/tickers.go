package models

import (
	"github.com/polygon-io/client-go/rest/client"
)

const (
	ListTickersPath      = "/v3/reference/tickers"
	GetTickerDetailsPath = "/v3/reference/tickers/{ticker}"
	GetTickerTypesPath   = "/v3/reference/tickers/types"
)

// TickerDetails contains detailed information on a specified ticker symbol.
// For more details see https://polygon.io/docs/stocks/get_v3_reference_tickers__ticker.
type TickerDetails struct {
	Ticker                      string         `json:"ticker"`
	Name                        string         `json:"name"`
	Market                      string         `json:"market"`
	Locale                      string         `json:"locale"`
	PrimaryExchange             string         `json:"primary_exchange,omitempty"`
	Type                        string         `json:"type,omitempty"`
	Active                      bool           `json:"active"`
	CurrencySymbol              string         `json:"currency_symbol,omitempty"`
	CurrencyName                string         `json:"currency_name,omitempty"`
	BaseCurrencySymbol          string         `json:"base_currency_symbol,omitempty"`
	BaseCurrencyName            string         `json:"base_currency_name,omitempty"`
	CUSIP                       string         `json:"cusip,omitempty"`
	CIK                         string         `json:"cik,omitempty"`
	CompositeFIGI               string         `json:"composite_figi,omitempty"`
	ShareClassFIGI              string         `json:"share_class_figi,omitempty"`
	LastUpdatedUTC              int64          `json:"last_updated_utc,omitempty"`
	DelistedUTC                 int64          `json:"delisted_utc,omitempty"`
	MarketCap                   float64        `json:"market_cap,omitempty"`
	PhoneNumber                 string         `json:"phone_number,omitempty"`
	Address                     CompanyAddress `json:"address,omitempty"` // todo: ptr?
	Description                 string         `json:"description,omitempty"`
	SICCode                     string         `json:"sic_code,omitempty"`
	SICDescription              string         `json:"sic_description,omitempty"`
	TickerRoot                  string         `json:"ticker_root,omitempty"`
	TickerSuffix                string         `json:"ticker_suffix,omitempty"`
	HomepageURL                 string         `json:"homepage_url,omitempty"`
	TotalEmployees              int32          `json:"total_employees,omitempty"`
	ListDate                    string         `json:"list_date,omitempty"`
	Branding                    Branding       `json:"branding,omitempty"` // todo: ptr?
	ShareClassSharesOutstanding int64          `json:"share_class_shares_outstanding,omitempty"`
	WeightedSharesOutstanding   int64          `json:"weighted_shares_outstanding,omitempty"`
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

// TickersResponse contains a list of reference tickers.
type TickersResponse struct {
	client.BaseResponse
	Results []TickerDetails `json:"results,omitempty"`
}

// ListTickersParams is the set of path and query parameters that are used to request reference tickers.
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

// Path maps the path parameters to their respective keys.
func (p ListTickersParams) Path() map[string]string {
	return map[string]string{}
}

// GetTickerDetailsParams is the set of path and query parameters that are used to request reference ticker details.
type GetTickerDetailsParams struct {
	Ticker string `validate:"required"`

	Date *string `query:"date"` // todo: this is "2006-01-02" format
}

// Path maps the path parameters to their respective keys.
func (p GetTickerDetailsParams) Path() map[string]string {
	return map[string]string{
		"ticker": p.Ticker,
	}
}

// TickerType represents a type of ticker with a code that the API understands.
type TickerType struct {
	AssetClass  string `json:"asset_class"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Locale      string `json:"locale"`
}

// TickerTypesResponse contains a list of ticker types.
type TickerTypesResponse struct {
	client.BaseResponse
	Results []*TickerType `json:"results,omitempty"`
}

// GetTickerTypesParams is the set of path and query parameters that are used to request ticker types.
type GetTickerTypesParams struct {
	AssetClass *string       `query:"asset_class"` // todo: this is similar but slightly different than market type (also we offer four options but only one returns results)
	Locale     *MarketLocale `query:"locale"`
}

// Path maps the path parameters to their respective keys.
func (p GetTickerTypesParams) Path() map[string]string {
	return map[string]string{}
}
