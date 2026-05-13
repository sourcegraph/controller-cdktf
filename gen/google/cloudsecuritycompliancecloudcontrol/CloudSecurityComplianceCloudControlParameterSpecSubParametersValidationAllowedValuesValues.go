package cloudsecuritycompliancecloudcontrol


type CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValues struct {
	// Represents a boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#bool_value CloudSecurityComplianceCloudControl#bool_value}
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Represents a double value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#number_value CloudSecurityComplianceCloudControl#number_value}
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// oneof_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#oneof_value CloudSecurityComplianceCloudControl#oneof_value}
	OneofValue *CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesOneofValue `field:"optional" json:"oneofValue" yaml:"oneofValue"`
	// string_list_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#string_list_value CloudSecurityComplianceCloudControl#string_list_value}
	StringListValue *CloudSecurityComplianceCloudControlParameterSpecSubParametersValidationAllowedValuesValuesStringListValue `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// Represents a string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#string_value CloudSecurityComplianceCloudControl#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

