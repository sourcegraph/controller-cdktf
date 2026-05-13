package iapwebregionforwardingruleserviceiammember


type IapWebRegionForwardingRuleServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iap_web_region_forwarding_rule_service_iam_member#expression IapWebRegionForwardingRuleServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iap_web_region_forwarding_rule_service_iam_member#title IapWebRegionForwardingRuleServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/iap_web_region_forwarding_rule_service_iam_member#description IapWebRegionForwardingRuleServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

