package googlelookerinstance


type GoogleLookerInstanceOauthConfig struct {
	// The client ID for the Oauth config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#client_id GoogleLookerInstance#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret for the Oauth config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_looker_instance#client_secret GoogleLookerInstance#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
}

