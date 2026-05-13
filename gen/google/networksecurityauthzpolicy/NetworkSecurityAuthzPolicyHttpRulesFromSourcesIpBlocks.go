package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesFromSourcesIpBlocks struct {
	// The length of the address range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#length NetworkSecurityAuthzPolicy#length}
	Length *float64 `field:"required" json:"length" yaml:"length"`
	// The address prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#prefix NetworkSecurityAuthzPolicy#prefix}
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

