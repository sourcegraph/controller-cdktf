package googleiapwebregionforwardingruleserviceiammember


type GoogleIapWebRegionForwardingRuleServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_region_forwarding_rule_service_iam_member#expression GoogleIapWebRegionForwardingRuleServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_region_forwarding_rule_service_iam_member#title GoogleIapWebRegionForwardingRuleServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iap_web_region_forwarding_rule_service_iam_member#description GoogleIapWebRegionForwardingRuleServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

