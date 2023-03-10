package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRegexMatchStatementFieldToMatchJsonBodyMatchPattern struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#all Wafv2RuleGroup#all}
	All *Wafv2RuleGroupRuleStatementRegexMatchStatementFieldToMatchJsonBodyMatchPatternAll `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#included_paths Wafv2RuleGroup#included_paths}.
	IncludedPaths *[]*string `field:"optional" json:"includedPaths" yaml:"includedPaths"`
}

