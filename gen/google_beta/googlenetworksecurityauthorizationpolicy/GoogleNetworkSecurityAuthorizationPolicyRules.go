package googlenetworksecurityauthorizationpolicy


type GoogleNetworkSecurityAuthorizationPolicyRules struct {
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_authorization_policy#destinations GoogleNetworkSecurityAuthorizationPolicy#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_security_authorization_policy#sources GoogleNetworkSecurityAuthorizationPolicy#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

