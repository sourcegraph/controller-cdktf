package googleiapsettings


type GoogleIapSettingsAccessSettings struct {
	// allowed_domains_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#allowed_domains_settings GoogleIapSettings#allowed_domains_settings}
	AllowedDomainsSettings *GoogleIapSettingsAccessSettingsAllowedDomainsSettings `field:"optional" json:"allowedDomainsSettings" yaml:"allowedDomainsSettings"`
	// cors_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#cors_settings GoogleIapSettings#cors_settings}
	CorsSettings *GoogleIapSettingsAccessSettingsCorsSettings `field:"optional" json:"corsSettings" yaml:"corsSettings"`
	// gcip_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#gcip_settings GoogleIapSettings#gcip_settings}
	GcipSettings *GoogleIapSettingsAccessSettingsGcipSettings `field:"optional" json:"gcipSettings" yaml:"gcipSettings"`
	// Identity sources that IAP can use to authenticate the end user.
	//
	// Only one identity source
	// can be configured. The possible values are:
	//
	// * 'WORKFORCE_IDENTITY_FEDERATION': Use external identities set up on Google Cloud Workforce
	//   				     Identity Federation. Possible values: ["WORKFORCE_IDENTITY_FEDERATION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#identity_sources GoogleIapSettings#identity_sources}
	IdentitySources *[]*string `field:"optional" json:"identitySources" yaml:"identitySources"`
	// oauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#oauth_settings GoogleIapSettings#oauth_settings}
	OauthSettings *GoogleIapSettingsAccessSettingsOauthSettings `field:"optional" json:"oauthSettings" yaml:"oauthSettings"`
	// reauth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#reauth_settings GoogleIapSettings#reauth_settings}
	ReauthSettings *GoogleIapSettingsAccessSettingsReauthSettings `field:"optional" json:"reauthSettings" yaml:"reauthSettings"`
	// workforce_identity_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_iap_settings#workforce_identity_settings GoogleIapSettings#workforce_identity_settings}
	WorkforceIdentitySettings *GoogleIapSettingsAccessSettingsWorkforceIdentitySettings `field:"optional" json:"workforceIdentitySettings" yaml:"workforceIdentitySettings"`
}

