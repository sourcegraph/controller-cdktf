package wafv2rulegroup


type Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatch struct {
	// all_query_arguments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#all_query_arguments Wafv2RuleGroup#all_query_arguments}
	AllQueryArguments *Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchAllQueryArguments `field:"optional" json:"allQueryArguments" yaml:"allQueryArguments"`
	// body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#body Wafv2RuleGroup#body}
	Body *Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchBody `field:"optional" json:"body" yaml:"body"`
	// cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#cookies Wafv2RuleGroup#cookies}
	Cookies *Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchCookies `field:"optional" json:"cookies" yaml:"cookies"`
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#headers Wafv2RuleGroup#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// json_body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#json_body Wafv2RuleGroup#json_body}
	JsonBody *Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchJsonBody `field:"optional" json:"jsonBody" yaml:"jsonBody"`
	// method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#method Wafv2RuleGroup#method}
	Method *Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchMethod `field:"optional" json:"method" yaml:"method"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#query_string Wafv2RuleGroup#query_string}
	QueryString *Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchQueryString `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#single_header Wafv2RuleGroup#single_header}
	SingleHeader *Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleHeader `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// single_query_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#single_query_argument Wafv2RuleGroup#single_query_argument}
	SingleQueryArgument *Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchSingleQueryArgument `field:"optional" json:"singleQueryArgument" yaml:"singleQueryArgument"`
	// uri_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_rule_group#uri_path Wafv2RuleGroup#uri_path}
	UriPath *Wafv2RuleGroupRuleStatementOrStatementStatementSizeConstraintStatementFieldToMatchUriPath `field:"optional" json:"uriPath" yaml:"uriPath"`
}

