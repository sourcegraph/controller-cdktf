package googlecestool


type GoogleCesToolDataStoreToolBoostSpecsSpecConditionBoostSpecsBoostControlSpecControlPoints struct {
	// Can be one of: 1.
	//
	// The numerical field value.
	// 2. The duration spec for freshness:
	// The value must be formatted as an XSD 'dayTimeDuration' value (a
	// restricted subset of an ISO 8601 duration value). The pattern for
	// this is: 'nDnM]'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#attribute_value GoogleCesTool#attribute_value}
	AttributeValue *string `field:"optional" json:"attributeValue" yaml:"attributeValue"`
	// The value between -1 to 1 by which to boost the score if the attribute_value evaluates to the value specified above.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#boost_amount GoogleCesTool#boost_amount}
	BoostAmount *float64 `field:"optional" json:"boostAmount" yaml:"boostAmount"`
}

