package googlechroniclefeed


type GoogleChronicleFeedDetailsMandiantIocSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsMandiantIocSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// time since when to start fetching the IOCs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#start_time GoogleChronicleFeed#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

