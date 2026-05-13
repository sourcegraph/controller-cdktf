package cloudsecuritycompliancecloudcontrol


type CloudSecurityComplianceCloudControlRulesCelExpression struct {
	// Logic expression in CEL language. The max length of the condition is 1000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#expression CloudSecurityComplianceCloudControl#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// resource_types_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloud_security_compliance_cloud_control#resource_types_values CloudSecurityComplianceCloudControl#resource_types_values}
	ResourceTypesValues *CloudSecurityComplianceCloudControlRulesCelExpressionResourceTypesValues `field:"optional" json:"resourceTypesValues" yaml:"resourceTypesValues"`
}

