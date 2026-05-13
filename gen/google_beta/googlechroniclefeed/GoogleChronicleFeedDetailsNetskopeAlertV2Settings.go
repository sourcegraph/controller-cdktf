package googlechroniclefeed


type GoogleChronicleFeedDetailsNetskopeAlertV2Settings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsNetskopeAlertV2SettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Content Category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#content_category GoogleChronicleFeed#content_category}
	ContentCategory *string `field:"optional" json:"contentCategory" yaml:"contentCategory"`
	// Content type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#content_types GoogleChronicleFeed#content_types}
	ContentTypes *[]*string `field:"optional" json:"contentTypes" yaml:"contentTypes"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

