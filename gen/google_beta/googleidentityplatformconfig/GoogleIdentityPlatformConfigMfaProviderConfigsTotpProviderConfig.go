package googleidentityplatformconfig


type GoogleIdentityPlatformConfigMfaProviderConfigsTotpProviderConfig struct {
	// The allowed number of adjacent intervals that will be used for verification to avoid clock skew.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#adjacent_intervals GoogleIdentityPlatformConfig#adjacent_intervals}
	AdjacentIntervals *float64 `field:"optional" json:"adjacentIntervals" yaml:"adjacentIntervals"`
}

