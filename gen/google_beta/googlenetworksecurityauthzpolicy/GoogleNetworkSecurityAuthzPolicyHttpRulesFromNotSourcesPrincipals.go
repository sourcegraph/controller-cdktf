package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyHttpRulesFromNotSourcesPrincipals struct {
	// The input string must have the substring specified here.
	//
	// Note: empty contains match is not allowed, please use regex instead.
	// Examples:
	// * abc matches the value xyz.abc.def
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_authz_policy#contains GoogleNetworkSecurityAuthzPolicy#contains}
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// The input string must match exactly the string specified here. Examples: * abc only matches the value abc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_authz_policy#exact GoogleNetworkSecurityAuthzPolicy#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// If true, indicates the exact/prefix/suffix/contains matching should be case insensitive.
	//
	// For example, the matcher data will match both input string Data and data if set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_authz_policy#ignore_case GoogleNetworkSecurityAuthzPolicy#ignore_case}
	IgnoreCase interface{} `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// The input string must have the prefix specified here.
	//
	// Note: empty prefix is not allowed, please use regex instead.
	// Examples:
	// * abc matches the value abc.xyz
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_authz_policy#prefix GoogleNetworkSecurityAuthzPolicy#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The input string must have the suffix specified here.
	//
	// Note: empty prefix is not allowed, please use regex instead.
	// Examples:
	// * abc matches the value xyz.abc
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_authz_policy#suffix GoogleNetworkSecurityAuthzPolicy#suffix}
	Suffix *string `field:"optional" json:"suffix" yaml:"suffix"`
}

