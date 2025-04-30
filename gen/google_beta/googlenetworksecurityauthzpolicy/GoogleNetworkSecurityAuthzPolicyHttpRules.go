package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyHttpRules struct {
	// from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_authz_policy#from GoogleNetworkSecurityAuthzPolicy#from}
	From *GoogleNetworkSecurityAuthzPolicyHttpRulesFrom `field:"optional" json:"from" yaml:"from"`
	// to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_authz_policy#to GoogleNetworkSecurityAuthzPolicy#to}
	To *GoogleNetworkSecurityAuthzPolicyHttpRulesTo `field:"optional" json:"to" yaml:"to"`
	// CEL expression that describes the conditions to be satisfied for the action.
	//
	// The result of the CEL expression is ANDed with the from and to. Refer to the CEL language reference for a list of available attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_authz_policy#when GoogleNetworkSecurityAuthzPolicy#when}
	When *string `field:"optional" json:"when" yaml:"when"`
}

