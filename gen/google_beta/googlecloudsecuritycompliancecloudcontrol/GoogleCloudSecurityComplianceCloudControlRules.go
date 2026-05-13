package googlecloudsecuritycompliancecloudcontrol


type GoogleCloudSecurityComplianceCloudControlRules struct {
	// The functionality enabled by the Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#rule_action_types GoogleCloudSecurityComplianceCloudControl#rule_action_types}
	RuleActionTypes *[]*string `field:"required" json:"ruleActionTypes" yaml:"ruleActionTypes"`
	// cel_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#cel_expression GoogleCloudSecurityComplianceCloudControl#cel_expression}
	CelExpression *GoogleCloudSecurityComplianceCloudControlRulesCelExpression `field:"optional" json:"celExpression" yaml:"celExpression"`
	// Description of the Rule. The maximum length is 2000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloud_security_compliance_cloud_control#description GoogleCloudSecurityComplianceCloudControl#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

