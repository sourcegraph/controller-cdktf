package googlechroniclefeed


type GoogleChronicleFeedDetailsWorkdaySettingsAuthentication struct {
	// Client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#client_id GoogleChronicleFeed#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Client Secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#client_secret GoogleChronicleFeed#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Refresh Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#refresh_token GoogleChronicleFeed#refresh_token}
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
	// The access token used to authenticate against Workday.
	//
	// This field is called
	// "secret" to maintain backwards compatibility. Workday was (only) configured
	// using username (which was unused) and secret (which is used as the access
	// token). Either this field or all of the other OAuth fields below must be
	// specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#secret GoogleChronicleFeed#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
	// Token endpoint to get the OAuth token from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#token_endpoint GoogleChronicleFeed#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// Username.
	//
	// This is unused: Workday feeds were originally configured using a
	// username and secret authentication method, but only the secret field was
	// used, and it was used to supply the OAuth access token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#user GoogleChronicleFeed#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

