package googlechroniclefeed


type GoogleChronicleFeedDetailsPanIocSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsPanIocSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// PAN IOC feed name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#feed GoogleChronicleFeed#feed}
	Feed *string `field:"optional" json:"feed" yaml:"feed"`
	// PAN IOC feed ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#feed_id GoogleChronicleFeed#feed_id}
	FeedId *string `field:"optional" json:"feedId" yaml:"feedId"`
}

