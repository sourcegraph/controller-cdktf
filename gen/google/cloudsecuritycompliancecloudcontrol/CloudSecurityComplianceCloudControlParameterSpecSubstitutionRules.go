package cloudsecuritycompliancecloudcontrol


type CloudSecurityComplianceCloudControlParameterSpecSubstitutionRules struct {
	// attribute_substitution_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#attribute_substitution_rule CloudSecurityComplianceCloudControl#attribute_substitution_rule}
	AttributeSubstitutionRule *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesAttributeSubstitutionRule `field:"optional" json:"attributeSubstitutionRule" yaml:"attributeSubstitutionRule"`
	// placeholder_substitution_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#placeholder_substitution_rule CloudSecurityComplianceCloudControl#placeholder_substitution_rule}
	PlaceholderSubstitutionRule *CloudSecurityComplianceCloudControlParameterSpecSubstitutionRulesPlaceholderSubstitutionRule `field:"optional" json:"placeholderSubstitutionRule" yaml:"placeholderSubstitutionRule"`
}

