package googlechroniclefeed


type GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettingsAuthentication struct {
	// claims block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#claims GoogleChronicleFeed#claims}
	Claims *GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettingsAuthenticationClaims `field:"optional" json:"claims" yaml:"claims"`
	// rs_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#rs_credentials GoogleChronicleFeed#rs_credentials}
	RsCredentials *GoogleChronicleFeedDetailsGoogleCloudIdentityDeviceUsersSettingsAuthenticationRsCredentials `field:"optional" json:"rsCredentials" yaml:"rsCredentials"`
	// Token endpoint to get the OAuth token from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#token_endpoint GoogleChronicleFeed#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

