package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsHeaderSetHeaders struct {
	// Specifies the name of the header in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#name GoogleNetworkSecurityAuthzPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#value GoogleNetworkSecurityAuthzPolicy#value}
	Value *GoogleNetworkSecurityAuthzPolicyHttpRulesToOperationsHeaderSetHeadersValue `field:"optional" json:"value" yaml:"value"`
}

