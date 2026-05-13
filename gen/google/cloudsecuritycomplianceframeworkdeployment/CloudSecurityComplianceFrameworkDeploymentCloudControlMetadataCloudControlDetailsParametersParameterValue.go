package cloudsecuritycomplianceframeworkdeployment


type CloudSecurityComplianceFrameworkDeploymentCloudControlMetadataCloudControlDetailsParametersParameterValue struct {
	// Represents a boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#bool_value CloudSecurityComplianceFrameworkDeployment#bool_value}
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Represents a double value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#number_value CloudSecurityComplianceFrameworkDeployment#number_value}
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// oneof_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#oneof_value CloudSecurityComplianceFrameworkDeployment#oneof_value}
	OneofValue *CloudSecurityComplianceFrameworkDeploymentCloudControlMetadataCloudControlDetailsParametersParameterValueOneofValue `field:"optional" json:"oneofValue" yaml:"oneofValue"`
	// string_list_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#string_list_value CloudSecurityComplianceFrameworkDeployment#string_list_value}
	StringListValue *CloudSecurityComplianceFrameworkDeploymentCloudControlMetadataCloudControlDetailsParametersParameterValueStringListValue `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// Represents a string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework_deployment#string_value CloudSecurityComplianceFrameworkDeployment#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

