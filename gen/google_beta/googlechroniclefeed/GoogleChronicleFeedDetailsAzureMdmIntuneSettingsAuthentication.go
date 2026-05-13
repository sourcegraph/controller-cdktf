package googlechroniclefeed


type GoogleChronicleFeedDetailsAzureMdmIntuneSettingsAuthentication struct {
	// Client ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#client_id GoogleChronicleFeed#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Client secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#client_secret GoogleChronicleFeed#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

