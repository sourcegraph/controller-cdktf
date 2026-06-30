package networksecurityauthzpolicy

type NetworkSecurityAuthzPolicyHttpRulesFromNotSources struct {
	// ip_blocks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_authz_policy#ip_blocks NetworkSecurityAuthzPolicy#ip_blocks}
	IpBlocks any `field:"optional" json:"ipBlocks" yaml:"ipBlocks"`
	// principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_authz_policy#principals NetworkSecurityAuthzPolicy#principals}
	Principals any `field:"optional" json:"principals" yaml:"principals"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_security_authz_policy#resources NetworkSecurityAuthzPolicy#resources}
	Resources any `field:"optional" json:"resources" yaml:"resources"`
}
