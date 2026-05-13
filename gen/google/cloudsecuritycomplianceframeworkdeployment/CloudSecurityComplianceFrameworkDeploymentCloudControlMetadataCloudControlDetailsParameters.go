package cloudsecuritycomplianceframeworkdeployment


type CloudSecurityComplianceFrameworkDeploymentCloudControlMetadataCloudControlDetailsParameters struct {
	// The name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#name CloudSecurityComplianceFrameworkDeployment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// parameter_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#parameter_value CloudSecurityComplianceFrameworkDeployment#parameter_value}
	ParameterValue *CloudSecurityComplianceFrameworkDeploymentCloudControlMetadataCloudControlDetailsParametersParameterValue `field:"required" json:"parameterValue" yaml:"parameterValue"`
}

