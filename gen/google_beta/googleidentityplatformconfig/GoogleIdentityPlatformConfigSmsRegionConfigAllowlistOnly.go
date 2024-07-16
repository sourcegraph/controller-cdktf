package googleidentityplatformconfig


type GoogleIdentityPlatformConfigSmsRegionConfigAllowlistOnly struct {
	// Two letter unicode region codes to allow as defined by https://cldr.unicode.org/ The full list of these region codes is here: https://github.com/unicode-cldr/cldr-localenames-full/blob/master/main/en/territories.json.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#allowed_regions GoogleIdentityPlatformConfig#allowed_regions}
	AllowedRegions *[]*string `field:"optional" json:"allowedRegions" yaml:"allowedRegions"`
}

