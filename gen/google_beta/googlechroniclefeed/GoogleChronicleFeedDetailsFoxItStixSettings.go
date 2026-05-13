package googlechroniclefeed


type GoogleChronicleFeedDetailsFoxItStixSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsFoxItStixSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Collection available at the poll service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#collection GoogleChronicleFeed#collection}
	Collection *string `field:"optional" json:"collection" yaml:"collection"`
	// TAXII poll service URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#poll_service_uri GoogleChronicleFeed#poll_service_uri}
	PollServiceUri *string `field:"optional" json:"pollServiceUri" yaml:"pollServiceUri"`
	// ssl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#ssl GoogleChronicleFeed#ssl}
	Ssl *GoogleChronicleFeedDetailsFoxItStixSettingsSsl `field:"optional" json:"ssl" yaml:"ssl"`
}

