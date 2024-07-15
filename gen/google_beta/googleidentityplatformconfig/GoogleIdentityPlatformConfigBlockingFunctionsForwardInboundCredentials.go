package googleidentityplatformconfig


type GoogleIdentityPlatformConfigBlockingFunctionsForwardInboundCredentials struct {
	// Whether to pass the user's OAuth identity provider's access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#access_token GoogleIdentityPlatformConfig#access_token}
	AccessToken interface{} `field:"optional" json:"accessToken" yaml:"accessToken"`
	// Whether to pass the user's OIDC identity provider's ID token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#id_token GoogleIdentityPlatformConfig#id_token}
	IdToken interface{} `field:"optional" json:"idToken" yaml:"idToken"`
	// Whether to pass the user's OAuth identity provider's refresh token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_identity_platform_config#refresh_token GoogleIdentityPlatformConfig#refresh_token}
	RefreshToken interface{} `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

