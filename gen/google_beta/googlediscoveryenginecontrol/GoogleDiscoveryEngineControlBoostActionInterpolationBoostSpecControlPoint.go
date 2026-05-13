package googlediscoveryenginecontrol


type GoogleDiscoveryEngineControlBoostActionInterpolationBoostSpecControlPoint struct {
	// The attribute value of the control point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#attribute_value GoogleDiscoveryEngineControl#attribute_value}
	AttributeValue *string `field:"optional" json:"attributeValue" yaml:"attributeValue"`
	// The value between -1 to 1 by which to boost the score if the attributeValue evaluates to the value specified above.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#boost_amount GoogleDiscoveryEngineControl#boost_amount}
	BoostAmount *float64 `field:"optional" json:"boostAmount" yaml:"boostAmount"`
}

