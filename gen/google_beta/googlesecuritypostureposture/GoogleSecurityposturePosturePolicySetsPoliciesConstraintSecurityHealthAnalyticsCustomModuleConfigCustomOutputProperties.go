package googlesecuritypostureposture


type GoogleSecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModuleConfigCustomOutputProperties struct {
	// Name of the property for the custom output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#name GoogleSecurityposturePosture#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// value_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_securityposture_posture#value_expression GoogleSecurityposturePosture#value_expression}
	ValueExpression *GoogleSecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModuleConfigCustomOutputPropertiesValueExpression `field:"optional" json:"valueExpression" yaml:"valueExpression"`
}

