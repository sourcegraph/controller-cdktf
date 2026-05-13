package googlecloudsecuritycomplianceframework


type GoogleCloudSecurityComplianceFrameworkCloudControlDetailsParametersParameterValueOneofValue struct {
	// The name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework#name GoogleCloudSecurityComplianceFramework#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// parameter_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework#parameter_value GoogleCloudSecurityComplianceFramework#parameter_value}
	ParameterValue *GoogleCloudSecurityComplianceFrameworkCloudControlDetailsParametersParameterValueOneofValueParameterValue `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

