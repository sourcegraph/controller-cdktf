package googlecloudsecuritycompliancecloudcontrol


type GoogleCloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesAttributeSubstitutionRule struct {
	// Fully qualified proto attribute path (in dot notation). Example: rules[0].cel_expression.resource_types_values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#attribute GoogleCloudSecurityComplianceCloudControl#attribute}
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

