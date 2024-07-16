package googleidentityplatformconfig


type GoogleIdentityPlatformConfigSmsRegionConfigAllowByDefault struct {
	// Two letter unicode region codes to disallow as defined by https://cldr.unicode.org/ The full list of these region codes is here: https://github.com/unicode-cldr/cldr-localenames-full/blob/master/main/en/territories.json.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#disallowed_regions GoogleIdentityPlatformConfig#disallowed_regions}
	DisallowedRegions *[]*string `field:"optional" json:"disallowedRegions" yaml:"disallowedRegions"`
}

