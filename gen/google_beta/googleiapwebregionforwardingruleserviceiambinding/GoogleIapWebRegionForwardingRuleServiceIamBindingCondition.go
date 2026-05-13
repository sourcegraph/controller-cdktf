package googleiapwebregionforwardingruleserviceiambinding


type GoogleIapWebRegionForwardingRuleServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_region_forwarding_rule_service_iam_binding#expression GoogleIapWebRegionForwardingRuleServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_region_forwarding_rule_service_iam_binding#title GoogleIapWebRegionForwardingRuleServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_region_forwarding_rule_service_iam_binding#description GoogleIapWebRegionForwardingRuleServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

