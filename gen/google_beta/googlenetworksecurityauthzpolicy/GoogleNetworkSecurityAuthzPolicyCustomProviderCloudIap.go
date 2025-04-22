package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyCustomProviderCloudIap struct {
	// Enable Cloud IAP at the AuthzPolicy level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#enabled GoogleNetworkSecurityAuthzPolicy#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

