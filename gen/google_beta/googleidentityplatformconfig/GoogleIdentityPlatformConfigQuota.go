package googleidentityplatformconfig


type GoogleIdentityPlatformConfigQuota struct {
	// sign_up_quota_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#sign_up_quota_config GoogleIdentityPlatformConfig#sign_up_quota_config}
	SignUpQuotaConfig *GoogleIdentityPlatformConfigQuotaSignUpQuotaConfig `field:"optional" json:"signUpQuotaConfig" yaml:"signUpQuotaConfig"`
}

