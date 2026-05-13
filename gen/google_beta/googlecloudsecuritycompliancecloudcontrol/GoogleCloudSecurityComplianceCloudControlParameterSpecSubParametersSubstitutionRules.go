package googlecloudsecuritycompliancecloudcontrol


type GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersSubstitutionRules struct {
	// attribute_substitution_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#attribute_substitution_rule GoogleCloudSecurityComplianceCloudControl#attribute_substitution_rule}
	AttributeSubstitutionRule *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersSubstitutionRulesAttributeSubstitutionRule `field:"optional" json:"attributeSubstitutionRule" yaml:"attributeSubstitutionRule"`
	// placeholder_substitution_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#placeholder_substitution_rule GoogleCloudSecurityComplianceCloudControl#placeholder_substitution_rule}
	PlaceholderSubstitutionRule *GoogleCloudSecurityComplianceCloudControlParameterSpecSubParametersSubstitutionRulesPlaceholderSubstitutionRule `field:"optional" json:"placeholderSubstitutionRule" yaml:"placeholderSubstitutionRule"`
}

