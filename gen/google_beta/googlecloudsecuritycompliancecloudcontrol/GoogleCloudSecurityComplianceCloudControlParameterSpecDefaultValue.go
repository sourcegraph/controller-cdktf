package googlecloudsecuritycompliancecloudcontrol


type GoogleCloudSecurityComplianceCloudControlParameterSpecDefaultValue struct {
	// Represents a boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#bool_value GoogleCloudSecurityComplianceCloudControl#bool_value}
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Represents a double value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#number_value GoogleCloudSecurityComplianceCloudControl#number_value}
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// oneof_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#oneof_value GoogleCloudSecurityComplianceCloudControl#oneof_value}
	OneofValue *GoogleCloudSecurityComplianceCloudControlParameterSpecDefaultValueOneofValue `field:"optional" json:"oneofValue" yaml:"oneofValue"`
	// string_list_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#string_list_value GoogleCloudSecurityComplianceCloudControl#string_list_value}
	StringListValue *GoogleCloudSecurityComplianceCloudControlParameterSpecDefaultValueStringListValue `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// Represents a string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#string_value GoogleCloudSecurityComplianceCloudControl#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

