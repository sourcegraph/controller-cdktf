package googleidentityplatformconfig


type GoogleIdentityPlatformConfigSmsRegionConfig struct {
	// allow_by_default block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#allow_by_default GoogleIdentityPlatformConfig#allow_by_default}
	AllowByDefault *GoogleIdentityPlatformConfigSmsRegionConfigAllowByDefault `field:"optional" json:"allowByDefault" yaml:"allowByDefault"`
	// allowlist_only block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#allowlist_only GoogleIdentityPlatformConfig#allowlist_only}
	AllowlistOnly *GoogleIdentityPlatformConfigSmsRegionConfigAllowlistOnly `field:"optional" json:"allowlistOnly" yaml:"allowlistOnly"`
}

