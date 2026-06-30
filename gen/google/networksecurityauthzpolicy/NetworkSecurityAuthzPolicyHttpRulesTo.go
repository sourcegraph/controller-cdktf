package networksecurityauthzpolicy

type NetworkSecurityAuthzPolicyHttpRulesTo struct {
	// not_operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_authz_policy#not_operations NetworkSecurityAuthzPolicy#not_operations}
	NotOperations any `field:"optional" json:"notOperations" yaml:"notOperations"`
	// operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_authz_policy#operations NetworkSecurityAuthzPolicy#operations}
	Operations any `field:"optional" json:"operations" yaml:"operations"`
}
