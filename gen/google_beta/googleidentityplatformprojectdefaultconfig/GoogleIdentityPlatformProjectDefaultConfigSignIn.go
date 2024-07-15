package googleidentityplatformprojectdefaultconfig


type GoogleIdentityPlatformProjectDefaultConfigSignIn struct {
	// Whether to allow more than one account to have the same email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_project_default_config#allow_duplicate_emails GoogleIdentityPlatformProjectDefaultConfig#allow_duplicate_emails}
	AllowDuplicateEmails interface{} `field:"optional" json:"allowDuplicateEmails" yaml:"allowDuplicateEmails"`
	// anonymous block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_project_default_config#anonymous GoogleIdentityPlatformProjectDefaultConfig#anonymous}
	Anonymous *GoogleIdentityPlatformProjectDefaultConfigSignInAnonymous `field:"optional" json:"anonymous" yaml:"anonymous"`
	// email block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_project_default_config#email GoogleIdentityPlatformProjectDefaultConfig#email}
	Email *GoogleIdentityPlatformProjectDefaultConfigSignInEmail `field:"optional" json:"email" yaml:"email"`
	// phone_number block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_project_default_config#phone_number GoogleIdentityPlatformProjectDefaultConfig#phone_number}
	PhoneNumber *GoogleIdentityPlatformProjectDefaultConfigSignInPhoneNumber `field:"optional" json:"phoneNumber" yaml:"phoneNumber"`
}

