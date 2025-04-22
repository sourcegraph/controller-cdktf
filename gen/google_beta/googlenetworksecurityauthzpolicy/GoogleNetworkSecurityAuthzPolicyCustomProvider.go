package googlenetworksecurityauthzpolicy


type GoogleNetworkSecurityAuthzPolicyCustomProvider struct {
	// authz_extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#authz_extension GoogleNetworkSecurityAuthzPolicy#authz_extension}
	AuthzExtension *GoogleNetworkSecurityAuthzPolicyCustomProviderAuthzExtension `field:"optional" json:"authzExtension" yaml:"authzExtension"`
	// cloud_iap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_security_authz_policy#cloud_iap GoogleNetworkSecurityAuthzPolicy#cloud_iap}
	CloudIap *GoogleNetworkSecurityAuthzPolicyCustomProviderCloudIap `field:"optional" json:"cloudIap" yaml:"cloudIap"`
}

