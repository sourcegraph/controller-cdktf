package googlecloudsecuritycompliancecloudcontrol


type GoogleCloudSecurityComplianceCloudControlRulesCelExpression struct {
	// Logic expression in CEL language. The max length of the condition is 1000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#expression GoogleCloudSecurityComplianceCloudControl#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// resource_types_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#resource_types_values GoogleCloudSecurityComplianceCloudControl#resource_types_values}
	ResourceTypesValues *GoogleCloudSecurityComplianceCloudControlRulesCelExpressionResourceTypesValues `field:"optional" json:"resourceTypesValues" yaml:"resourceTypesValues"`
}

