package wafv2rulegroup


type Wafv2RuleGroupRuleActionCaptchaCustomRequestHandlingInsertHeader struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#name Wafv2RuleGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_rule_group#value Wafv2RuleGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

