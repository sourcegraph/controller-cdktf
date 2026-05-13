package googleidentityplatformoauthidpconfig


type GoogleIdentityPlatformOauthIdpConfigResponseType struct {
	// If true, authorization code is returned from IdP's authorization endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_identity_platform_oauth_idp_config#code GoogleIdentityPlatformOauthIdpConfig#code}
	Code interface{} `field:"optional" json:"code" yaml:"code"`
	// If true, ID token is returned from IdP's authorization endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_identity_platform_oauth_idp_config#id_token GoogleIdentityPlatformOauthIdpConfig#id_token}
	IdToken interface{} `field:"optional" json:"idToken" yaml:"idToken"`
}

