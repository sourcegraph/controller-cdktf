package wafv2webacl


type Wafv2WebAclRuleStatementManagedRuleGroupStatementRuleActionOverrideActionToUseBlockCustomResponse struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#response_code Wafv2WebAcl#response_code}.
	ResponseCode *float64 `field:"required" json:"responseCode" yaml:"responseCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#custom_response_body_key Wafv2WebAcl#custom_response_body_key}.
	CustomResponseBodyKey *string `field:"optional" json:"customResponseBodyKey" yaml:"customResponseBodyKey"`
	// response_header block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/wafv2_web_acl#response_header Wafv2WebAcl#response_header}
	ResponseHeader interface{} `field:"optional" json:"responseHeader" yaml:"responseHeader"`
}
