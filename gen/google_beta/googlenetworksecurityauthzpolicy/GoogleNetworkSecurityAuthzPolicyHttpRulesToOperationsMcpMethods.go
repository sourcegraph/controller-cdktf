package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsMcpMethods struct {
	// The MCP method to match against.
	//
	// Allowed values are as follows:
	// 1) “tools”, “prompts”, “resources” - these will match against all sub methods under the respective methods.
	// 2) “prompts/list”, “tools/list”, “resources/list”, “resources/templates/list”
	// 3) “prompts/get”, “tools/call”, “resources/subscribe”, “resources/unsubscribe”, “resources/read”
	// Params cannot be specified for categories 1) and 2).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_authz_policy#name GoogleNetworkSecurityAuthzPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_security_authz_policy#params GoogleNetworkSecurityAuthzPolicy#params}
	Params interface{} `field:"optional" json:"params" yaml:"params"`
}

