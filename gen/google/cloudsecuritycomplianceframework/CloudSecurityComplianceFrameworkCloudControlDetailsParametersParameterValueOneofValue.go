package cloudsecuritycomplianceframework


type CloudSecurityComplianceFrameworkCloudControlDetailsParametersParameterValueOneofValue struct {
	// The name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#name CloudSecurityComplianceFramework#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// parameter_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#parameter_value CloudSecurityComplianceFramework#parameter_value}
	ParameterValue *CloudSecurityComplianceFrameworkCloudControlDetailsParametersParameterValueOneofValueParameterValue `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

