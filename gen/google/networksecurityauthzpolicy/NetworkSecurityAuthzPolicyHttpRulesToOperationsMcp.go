package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesToOperationsMcp struct {
	// If specified, matches on the MCP protocol’s non-access specific methods namely: * initialize/ * completion/ * logging/ * notifications/ * ping Default value: "SKIP_BASE_PROTOCOL_METHODS" Possible values: ["SKIP_BASE_PROTOCOL_METHODS", "MATCH_BASE_PROTOCOL_METHODS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#base_protocol_methods_option NetworkSecurityAuthzPolicy#base_protocol_methods_option}
	BaseProtocolMethodsOption *string `field:"optional" json:"baseProtocolMethodsOption" yaml:"baseProtocolMethodsOption"`
	// methods block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#methods NetworkSecurityAuthzPolicy#methods}
	Methods interface{} `field:"optional" json:"methods" yaml:"methods"`
}

