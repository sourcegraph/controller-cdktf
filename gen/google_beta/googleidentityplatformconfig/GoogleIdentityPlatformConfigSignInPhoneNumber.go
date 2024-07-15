package googleidentityplatformconfig


type GoogleIdentityPlatformConfigSignInPhoneNumber struct {
	// Whether phone number auth is enabled for the project or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#enabled GoogleIdentityPlatformConfig#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// A map of <test phone number, fake code> that can be used for phone auth testing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#test_phone_numbers GoogleIdentityPlatformConfig#test_phone_numbers}
	TestPhoneNumbers *map[string]*string `field:"optional" json:"testPhoneNumbers" yaml:"testPhoneNumbers"`
}

