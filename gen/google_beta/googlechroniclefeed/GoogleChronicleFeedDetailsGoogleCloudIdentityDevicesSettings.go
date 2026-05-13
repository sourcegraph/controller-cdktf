package googlechroniclefeed


type GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettings struct {
	// API Version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#api_version GoogleChronicleFeed#api_version}
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsGoogleCloudIdentityDevicesSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
}

