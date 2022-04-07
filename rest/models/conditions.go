package models

const ListConditionsPath = "/v3/reference/conditions"

// ListConditionsParams is the set of parameters for the ListConditions method.
type ListConditionsParams struct {
	AssetClass *string `query:"asset_class,omitempty"` // todo: enum value
	DataType   *string `query:"data_type,omitempty"`   // todo: enum value
	ID         *int64  `query:"id,omitempty"`
	SIP        *string `query:"sip,omitempty"` // todo: enum value

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
	AssetClass   string                     `json:"asset_class,omitempty"` // todo: enum
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
