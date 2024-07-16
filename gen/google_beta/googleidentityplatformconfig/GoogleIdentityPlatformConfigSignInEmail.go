package googleidentityplatformconfig


type GoogleIdentityPlatformConfigSignInEmail struct {
	// Whether email auth is enabled for the project or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#enabled GoogleIdentityPlatformConfig#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Whether a password is required for email auth or not.
	//
	// If true, both an email and
	// password must be provided to sign in. If false, a user may sign in via either
	// email/password or email link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#password_required GoogleIdentityPlatformConfig#password_required}
	PasswordRequired interface{} `field:"optional" json:"passwordRequired" yaml:"passwordRequired"`
}

