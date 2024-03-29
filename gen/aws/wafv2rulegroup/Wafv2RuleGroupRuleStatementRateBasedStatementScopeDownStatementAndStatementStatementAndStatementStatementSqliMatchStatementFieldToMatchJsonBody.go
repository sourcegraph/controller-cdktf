package wafv2rulegroup


type Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchJsonBody struct {
	// match_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#match_pattern Wafv2RuleGroup#match_pattern}
	MatchPattern *Wafv2RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatementStatementAndStatementStatementSqliMatchStatementFieldToMatchJsonBodyMatchPattern `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#match_scope Wafv2RuleGroup#match_scope}.
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#invalid_fallback_behavior Wafv2RuleGroup#invalid_fallback_behavior}.
	InvalidFallbackBehavior *string `field:"optional" json:"invalidFallbackBehavior" yaml:"invalidFallbackBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#oversize_handling Wafv2RuleGroup#oversize_handling}.
	OversizeHandling *string `field:"optional" json:"oversizeHandling" yaml:"oversizeHandling"`
}

