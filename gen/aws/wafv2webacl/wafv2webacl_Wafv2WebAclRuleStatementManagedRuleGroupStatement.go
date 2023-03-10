package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#name Wafv2WebAcl#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#vendor_name Wafv2WebAcl#vendor_name}.
	VendorName *string `field:"required" json:"vendorName" yaml:"vendorName"`
	// excluded_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#excluded_rule Wafv2WebAcl#excluded_rule}
	ExcludedRule interface{} `field:"optional" json:"excludedRule" yaml:"excludedRule"`
	// managed_rule_group_configs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#managed_rule_group_configs Wafv2WebAcl#managed_rule_group_configs}
	ManagedRuleGroupConfigs interface{} `field:"optional" json:"managedRuleGroupConfigs" yaml:"managedRuleGroupConfigs"`
	// rule_action_override block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#rule_action_override Wafv2WebAcl#rule_action_override}
	RuleActionOverride interface{} `field:"optional" json:"ruleActionOverride" yaml:"ruleActionOverride"`
	// scope_down_statement block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#scope_down_statement Wafv2WebAcl#scope_down_statement}
	ScopeDownStatement *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatement `field:"optional" json:"scopeDownStatement" yaml:"scopeDownStatement"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#version Wafv2WebAcl#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

