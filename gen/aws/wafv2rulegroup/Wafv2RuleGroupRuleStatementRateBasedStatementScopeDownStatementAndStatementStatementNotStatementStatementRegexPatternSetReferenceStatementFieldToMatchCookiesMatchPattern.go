package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPattern struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#all Wafv2RuleGroup#all}
	All *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementNotStatementStatementRegexPatternSetReferenceStatementFieldToMatchCookiesMatchPatternAll `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#excluded_cookies Wafv2RuleGroup#excluded_cookies}.
	ExcludedCookies *[]*string `field:"optional" json:"excludedCookies" yaml:"excludedCookies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#included_cookies Wafv2RuleGroup#included_cookies}.
	IncludedCookies *[]*string `field:"optional" json:"includedCookies" yaml:"includedCookies"`
}

