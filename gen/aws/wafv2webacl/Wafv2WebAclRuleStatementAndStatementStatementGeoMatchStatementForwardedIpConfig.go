package wafv2webacl


type Wafv2WebAclRuleStatementAndStatementStatementGeoMatchStatementForwardedIpConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_web_acl#fallback_behavior Wafv2WebAcl#fallback_behavior}.
	FallbackBehavior *string `field:"required" json:"fallbackBehavior" yaml:"fallbackBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/wafv2_web_acl#header_name Wafv2WebAcl#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

