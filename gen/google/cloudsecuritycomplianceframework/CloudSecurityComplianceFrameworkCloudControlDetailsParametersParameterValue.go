package cloudsecuritycomplianceframework


type CloudSecurityComplianceFrameworkCloudControlDetailsParametersParameterValue struct {
	// Represents a boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#bool_value CloudSecurityComplianceFramework#bool_value}
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Represents a double value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#number_value CloudSecurityComplianceFramework#number_value}
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// oneof_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#oneof_value CloudSecurityComplianceFramework#oneof_value}
	OneofValue *CloudSecurityComplianceFrameworkCloudControlDetailsParametersParameterValueOneofValue `field:"optional" json:"oneofValue" yaml:"oneofValue"`
	// string_list_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#string_list_value CloudSecurityComplianceFramework#string_list_value}
	StringListValue *CloudSecurityComplianceFrameworkCloudControlDetailsParametersParameterValueStringListValue `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// Represents a string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_framework#string_value CloudSecurityComplianceFramework#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

