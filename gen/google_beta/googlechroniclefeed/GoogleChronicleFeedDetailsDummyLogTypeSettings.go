package googlechroniclefeed


type GoogleChronicleFeedDetailsDummyLogTypeSettings struct {
	// Full API Endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#api_endpoint GoogleChronicleFeed#api_endpoint}
	ApiEndpoint *string `field:"optional" json:"apiEndpoint" yaml:"apiEndpoint"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsDummyLogTypeSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
}

