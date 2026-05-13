package cloudsecuritycompliancecloudcontrol


type CloudSecurityComplianceCloudControlParameterSpec struct {
	// if the parameter is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#is_required CloudSecurityComplianceCloudControl#is_required}
	IsRequired interface{} `field:"required" json:"isRequired" yaml:"isRequired"`
	// The name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#name CloudSecurityComplianceCloudControl#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Parameter value type. Possible values: STRING BOOLEAN STRINGLIST NUMBER ONEOF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#value_type CloudSecurityComplianceCloudControl#value_type}
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// default_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#default_value CloudSecurityComplianceCloudControl#default_value}
	DefaultValue *CloudSecurityComplianceCloudControlParameterSpecDefaultValue `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The description of the parameter. The maximum length is 2000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#description CloudSecurityComplianceCloudControl#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the parameter. The maximum length is 200 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#display_name CloudSecurityComplianceCloudControl#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// sub_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#sub_parameters CloudSecurityComplianceCloudControl#sub_parameters}
	SubParameters interface{} `field:"optional" json:"subParameters" yaml:"subParameters"`
	// substitution_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#substitution_rules CloudSecurityComplianceCloudControl#substitution_rules}
	SubstitutionRules interface{} `field:"optional" json:"substitutionRules" yaml:"substitutionRules"`
	// validation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#validation CloudSecurityComplianceCloudControl#validation}
	Validation *CloudSecurityComplianceCloudControlParameterSpecValidation `field:"optional" json:"validation" yaml:"validation"`
}

