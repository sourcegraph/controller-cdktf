package cloudsecuritycompliancecloudcontrol


type CloudSecurityComplianceCloudControlParameterSpecSubParametersValidation struct {
	// allowed_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#allowed_values CloudSecurityComplianceCloudControl#allowed_values}
	AllowedValues *CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValues `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// int_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#int_range CloudSecurityComplianceCloudControl#int_range}
	IntRange *CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationIntRange `field:"optional" json:"intRange" yaml:"intRange"`
	// regexp_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#regexp_pattern CloudSecurityComplianceCloudControl#regexp_pattern}
	RegexpPattern *CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationRegexpPattern `field:"optional" json:"regexpPattern" yaml:"regexpPattern"`
}

