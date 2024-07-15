package googleidentityplatformconfig


type GoogleIdentityPlatformConfigQuotaSignUpQuotaConfig struct {
	// A sign up APIs quota that customers can override temporarily.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#quota GoogleIdentityPlatformConfig#quota}
	Quota *float64 `field:"optional" json:"quota" yaml:"quota"`
	// How long this quota will be active for. It is measurred in seconds, e.g., Example: "9.615s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#quota_duration GoogleIdentityPlatformConfig#quota_duration}
	QuotaDuration *string `field:"optional" json:"quotaDuration" yaml:"quotaDuration"`
	// When this quota will take affect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#start_time GoogleIdentityPlatformConfig#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

