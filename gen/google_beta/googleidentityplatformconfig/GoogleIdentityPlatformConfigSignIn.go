package googleidentityplatformconfig


type GoogleIdentityPlatformConfigSignIn struct {
	// Whether to allow more than one account to have the same email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#allow_duplicate_emails GoogleIdentityPlatformConfig#allow_duplicate_emails}
	AllowDuplicateEmails interface{} `field:"optional" json:"allowDuplicateEmails" yaml:"allowDuplicateEmails"`
	// anonymous block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#anonymous GoogleIdentityPlatformConfig#anonymous}
	Anonymous *GoogleIdentityPlatformConfigSignInAnonymous `field:"optional" json:"anonymous" yaml:"anonymous"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#email GoogleIdentityPlatformConfig#email}
	Email *GoogleIdentityPlatformConfigSignInEmail `field:"optional" json:"email" yaml:"email"`
	// phone_number block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#phone_number GoogleIdentityPlatformConfig#phone_number}
	PhoneNumber *GoogleIdentityPlatformConfigSignInPhoneNumber `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

