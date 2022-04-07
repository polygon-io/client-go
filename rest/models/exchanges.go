package models

const ListExchangesPath = "/v3/reference/exchanges"

// ListExchangesParams is the set of parameters for the ListExchanges method.
type ListExchangesParams struct {
	AssetClass *string `query:"asset_class,omitempty"` // todo: enum value
	Locale     *string `query:"locale,omitempty"`      // todo: enum value
}

// ListExchangesResponse is the response returned by the ListExchanges method.
type ListExchangesResponse struct {
	BaseResponse
	Results []Exchange `json:"results,omitempty"`
}

// Exchange contains detailed information on a specified stock Exchange.
type Exchange struct {
	Acronym       string `json:"acronym,omitempty"`
	AssetClass    string `json:"asset_class,omitempty"`
	ID            int64  `json:"id,omitempty"`
	Locale        string `json:"locale,omitempty"`
	MIC           string `json:"mic,omitempty"`
	Name          string `json:"name,omitempty"`
	OperatingMIC  string `json:"operating_mic,omitempty"`
	ParticipantID string `json:"participant_id,omitempty"`
	Type          string `json:"type,omitempty"`
	URL           string `json:"url,omitempty"`
}
