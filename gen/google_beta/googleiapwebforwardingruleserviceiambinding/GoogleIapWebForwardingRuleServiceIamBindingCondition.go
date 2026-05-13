package googleiapwebforwardingruleserviceiambinding


type GoogleIapWebForwardingRuleServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_forwarding_rule_service_iam_binding#expression GoogleIapWebForwardingRuleServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_forwarding_rule_service_iam_binding#title GoogleIapWebForwardingRuleServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_forwarding_rule_service_iam_binding#description GoogleIapWebForwardingRuleServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

