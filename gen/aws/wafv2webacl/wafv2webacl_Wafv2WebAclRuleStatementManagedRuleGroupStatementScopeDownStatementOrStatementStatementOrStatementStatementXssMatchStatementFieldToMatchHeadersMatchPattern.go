package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementOrStatementStatementOrStatementStatementXssMatchStatementFieldToMatchHeadersMatchPattern struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#all Wafv2WebAcl#all}
	All *Wafv2WebAclRuleStatementManagedRuleGroupStatementScopeDownStatementOrStatementStatementOrStatementStatementXssMatchStatementFieldToMatchHeadersMatchPatternAll `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#excluded_headers Wafv2WebAcl#excluded_headers}.
	ExcludedHeaders *[]*string `field:"optional" json:"excludedHeaders" yaml:"excludedHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#included_headers Wafv2WebAcl#included_headers}.
	IncludedHeaders *[]*string `field:"optional" json:"includedHeaders" yaml:"includedHeaders"`
}

