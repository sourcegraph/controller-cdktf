package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesToOperationsMcpMethodsParams struct {
	// A substring match on the MCP method parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#contains NetworkSecurityAuthzPolicy#contains}
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// An exact match on the MCP method parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#exact NetworkSecurityAuthzPolicy#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Specifies that the string match should be case insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#ignore_case NetworkSecurityAuthzPolicy#ignore_case}
	IgnoreCase interface{} `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// A prefix match on the MCP method parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#prefix NetworkSecurityAuthzPolicy#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// A suffix match on the MCP method parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#suffix NetworkSecurityAuthzPolicy#suffix}
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

