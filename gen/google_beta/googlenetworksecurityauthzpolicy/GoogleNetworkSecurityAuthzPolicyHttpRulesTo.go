package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyHttpRulesTo struct {
	// not_operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#not_operations GoogleNetworkSecurityAuthzPolicy#not_operations}
	NotOperations interface{} `field:"optional" json:"notOperations" yaml:"notOperations"`
	// operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#operations GoogleNetworkSecurityAuthzPolicy#operations}
	Operations interface{} `field:"optional" json:"operations" yaml:"operations"`
}

