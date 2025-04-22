package googleiapsettings


type GoogleIapSettingsAccessSettingsWorkforceIdentitySettings struct {
	// oauth2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_iap_settings#oauth2 GoogleIapSettings#oauth2}
	Oauth2 *GoogleIapSettingsAccessSettingsWorkforceIdentitySettingsOauth2 `field:"optional" json:"oauth2" yaml:"oauth2"`
	// The workforce pool resources. Only one workforce pool is accepted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_iap_settings#workforce_pools GoogleIapSettings#workforce_pools}
	WorkforcePools *[]*string `field:"optional" json:"workforcePools" yaml:"workforcePools"`
}

