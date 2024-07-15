package googleidentityplatformconfig


type GoogleIdentityPlatformConfigMfa struct {
	// A list of usable second factors for this project. Possible values: ["PHONE_SMS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#enabled_providers GoogleIdentityPlatformConfig#enabled_providers}
	EnabledProviders *[]*string `field:"optional" json:"enabledProviders" yaml:"enabledProviders"`
	// provider_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#provider_configs GoogleIdentityPlatformConfig#provider_configs}
	ProviderConfigs interface{} `field:"optional" json:"providerConfigs" yaml:"providerConfigs"`
	// Whether MultiFactor Authentication has been enabled for this project. Possible values: ["DISABLED", "ENABLED", "MANDATORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#state GoogleIdentityPlatformConfig#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

