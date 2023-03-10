package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUse struct {
	// allow block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#allow Wafv2WebAcl#allow}
	Allow *Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseAllow `field:"optional" json:"allow" yaml:"allow"`
	// block block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#block Wafv2WebAcl#block}
	Block *Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseBlock `field:"optional" json:"block" yaml:"block"`
	// captcha block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#captcha Wafv2WebAcl#captcha}
	Captcha *Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCaptcha `field:"optional" json:"captcha" yaml:"captcha"`
	// count block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#count Wafv2WebAcl#count}
	Count *Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseCount `field:"optional" json:"count" yaml:"count"`
}

