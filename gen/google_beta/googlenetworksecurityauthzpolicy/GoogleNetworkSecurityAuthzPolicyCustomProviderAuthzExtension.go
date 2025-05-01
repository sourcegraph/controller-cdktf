package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyCustomProviderAuthzExtension struct {
	// A list of references to authorization extensions that will be invoked for requests matching this policy.
	//
	// Limited to 1 custom provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_network_security_authz_policy#resources GoogleNetworkSecurityAuthzPolicy#resources}
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

