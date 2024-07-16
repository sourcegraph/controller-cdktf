package googleappengineapplication


type GoogleAppEngineApplicationIap struct {
	// OAuth2 client ID to use for the authentication flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#oauth2_client_id GoogleAppEngineApplication#oauth2_client_id}
	Oauth2ClientId *string `field:"required" json:"oauth2ClientId" yaml:"oauth2ClientId"`
	// OAuth2 client secret to use for the authentication flow.
	//
	// The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#oauth2_client_secret GoogleAppEngineApplication#oauth2_client_secret}
	Oauth2ClientSecret *string `field:"required" json:"oauth2ClientSecret" yaml:"oauth2ClientSecret"`
	// Adapted for use with the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_application#enabled GoogleAppEngineApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

