package googlecloudsecuritycompliancecloudcontrol


type GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValue struct {
	// The name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#name GoogleCloudSecurityComplianceCloudControl#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// parameter_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#parameter_value GoogleCloudSecurityComplianceCloudControl#parameter_value}
	ParameterValue *GoogleCloudSecurityComplianceCloudControlParameterSpecValidationAllowedValuesValuesOneofValueParameterValue `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

