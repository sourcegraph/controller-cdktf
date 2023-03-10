package wafv2webacl


type Wafv2WebAclRuleStatementRuleGroupReferenceStatement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#arn Wafv2WebAcl#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// excluded_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#excluded_rule Wafv2WebAcl#excluded_rule}
	ExcludedRule interface{} `field:"optional" json:"excludedRule" yaml:"excludedRule"`
}

