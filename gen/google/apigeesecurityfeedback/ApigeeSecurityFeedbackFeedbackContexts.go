package apigeesecurityfeedback


type ApigeeSecurityFeedbackFeedbackContexts struct {
	// The attribute the user is providing feedback about. Possible values: ["ATTRIBUTE_ENVIRONMENTS", "ATTRIBUTE_IP_ADDRESS_RANGES"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_security_feedback#attribute ApigeeSecurityFeedback#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The values of the attribute the user is providing feedback about, separated by commas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/apigee_security_feedback#values ApigeeSecurityFeedback#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

