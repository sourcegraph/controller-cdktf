package networksecurityauthzpolicy


type NetworkSecurityAuthzPolicyHttpRulesFromSources struct {
	// ip_blocks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#ip_blocks NetworkSecurityAuthzPolicy#ip_blocks}
	IpBlocks interface{} `field:"optional" json:"ipBlocks" yaml:"ipBlocks"`
	// principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#principals NetworkSecurityAuthzPolicy#principals}
	Principals interface{} `field:"optional" json:"principals" yaml:"principals"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/network_security_authz_policy#resources NetworkSecurityAuthzPolicy#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
}

