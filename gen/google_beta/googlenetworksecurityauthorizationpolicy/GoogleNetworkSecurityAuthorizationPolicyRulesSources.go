package googlenetworksecurityauthorizationpolicy


type GoogleNetworkSecurityAuthorizationPolicyRulesSources struct {
	// List of CIDR ranges to match based on source IP address.
	//
	// At least one IP block should match. Single IP (e.g., "1.2.3.4") and CIDR (e.g., "1.2.3.0/24") are supported. Authorization based on source IP alone should be avoided.
	// The IP addresses of any load balancers or proxies should be considered untrusted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_authorization_policy#ip_blocks GoogleNetworkSecurityAuthorizationPolicy#ip_blocks}
	IpBlocks *[]*string `field:"optional" json:"ipBlocks" yaml:"ipBlocks"`
	// List of peer identities to match for authorization.
	//
	// At least one principal should match. Each peer can be an exact match, or a prefix match (example, "namespace/*") or a suffix match (example, "* /service-account") or a presence match "*".
	// Authorization based on the principal name without certificate validation (configured by ServerTlsPolicy resource) is considered insecure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_authorization_policy#principals GoogleNetworkSecurityAuthorizationPolicy#principals}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Principals *[]*string `field:"optional" json:"principals" yaml:"principals"`
}

