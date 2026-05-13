package cloudsecuritycompliancecloudcontrol


type CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValue struct {
	// The name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#name CloudSecurityComplianceCloudControl#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// parameter_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#parameter_value CloudSecurityComplianceCloudControl#parameter_value}
	ParameterValue *CloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

