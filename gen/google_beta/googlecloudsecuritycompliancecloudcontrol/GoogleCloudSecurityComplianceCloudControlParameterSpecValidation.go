package googlecloudsecuritycompliancecloudcontrol


type GoogleCloudSecurityComplianceCloudControlParameterSpecValidation struct {
	// allowed_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#allowed_values GoogleCloudSecurityComplianceCloudControl#allowed_values}
	AllowedValues *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValues `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// int_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#int_range GoogleCloudSecurityComplianceCloudControl#int_range}
	IntRange *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationIntRange `field:"optional" json:"intRange" yaml:"intRange"`
	// regexp_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#regexp_pattern GoogleCloudSecurityComplianceCloudControl#regexp_pattern}
	RegexpPattern *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationRegexpPattern `field:"optional" json:"regexpPattern" yaml:"regexpPattern"`
}

