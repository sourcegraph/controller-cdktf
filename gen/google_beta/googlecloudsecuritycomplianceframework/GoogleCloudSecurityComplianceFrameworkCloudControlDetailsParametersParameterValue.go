package googlecloudsecuritycomplianceframework


type GoogleCloudSecurityComplianceFrameworkCloudControlDetailsParametersParameterValue struct {
	// Represents a boolean value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework#bool_value GoogleCloudSecurityComplianceFramework#bool_value}
	BoolValue interface{} `field:"optional" json:"boolValue" yaml:"boolValue"`
	// Represents a double value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework#number_value GoogleCloudSecurityComplianceFramework#number_value}
	NumberValue *float64 `field:"optional" json:"numberValue" yaml:"numberValue"`
	// oneof_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework#oneof_value GoogleCloudSecurityComplianceFramework#oneof_value}
	OneofValue *GoogleCloudSecurityComplianceFrameworkCloudControlDetailsParametersParameterValueOneofValue `field:"optional" json:"oneofValue" yaml:"oneofValue"`
	// string_list_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework#string_list_value GoogleCloudSecurityComplianceFramework#string_list_value}
	StringListValue *GoogleCloudSecurityComplianceFrameworkCloudControlDetailsParametersParameterValueStringListValue `field:"optional" json:"stringListValue" yaml:"stringListValue"`
	// Represents a string value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_framework#string_value GoogleCloudSecurityComplianceFramework#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

