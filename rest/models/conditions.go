package models

const ListConditionsPath = "/v3/reference/conditions"

// ListConditionsParams is the set of parameters for the ListConditions method.
type ListConditionsParams struct {
	AssetClass *AssetClass `query:"asset_class,omitempty"`
	DataType   *DataType   `query:"data_type,omitempty"`
	ID         *int64      `query:"id,omitempty"`
	SIP        *SIP        `query:"sip,omitempty"`

	Sort  *Sort  `query:"sort"`
	Order *Order `query:"order"`
	Limit *int   `query:"limit"`
}

// ListConditionsResponse is the response returned by the ListConditions method.
type ListConditionsResponse struct {
	BaseResponse
	Results []Condition `json:"results,omitempty"`
}

// Condition contains detailed information on a specified condition.
type Condition struct {
	Abbreviation string                     `json:"abbreviation,omitempty"`
	AssetClass   string                     `json:"asset_class,omitempty"`
	DataTypes    []string                   `json:"data_types,omitempty"`
	Description  string                     `json:"description,omitempty"`
	Exchange     int64                      `json:"exchange,omitempty"`
	ID           int64                      `json:"id,omitempty"`
	Legacy       bool                       `json:"legacy"`
	Name         string                     `json:"name,omitempty"`
	SIPMapping   map[string]string          `json:"sip_mapping,omitempty"`
	Type         string                     `json:"type,omitempty"`
	UpdateRules  map[string]map[string]bool `json:"update_rules,omitempty"`
}
