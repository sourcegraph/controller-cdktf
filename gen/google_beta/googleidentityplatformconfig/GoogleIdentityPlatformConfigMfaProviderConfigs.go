package googleidentityplatformconfig


type GoogleIdentityPlatformConfigMfaProviderConfigs struct {
	// Whether MultiFactor Authentication has been enabled for this project. Possible values: ["DISABLED", "ENABLED", "MANDATORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#state GoogleIdentityPlatformConfig#state}
	State *string `field:"optional" json:"state" yaml:"state"`
	// totp_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#totp_provider_config GoogleIdentityPlatformConfig#totp_provider_config}
	TotpProviderConfig *GoogleIdentityPlatformConfigMfaProviderConfigsTotpProviderConfig `field:"optional" json:"totpProviderConfig" yaml:"totpProviderConfig"`
}

