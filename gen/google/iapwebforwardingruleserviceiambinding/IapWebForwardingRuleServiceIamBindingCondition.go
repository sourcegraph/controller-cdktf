package iapwebforwardingruleserviceiambinding


type IapWebForwardingRuleServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iap_web_forwarding_rule_service_iam_binding#expression IapWebForwardingRuleServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iap_web_forwarding_rule_service_iam_binding#title IapWebForwardingRuleServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iap_web_forwarding_rule_service_iam_binding#description IapWebForwardingRuleServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

