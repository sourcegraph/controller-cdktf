package googlechroniclefeed


type GoogleChronicleFeedDetailsNetskopeAlertSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsNetskopeAlertSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Content type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#content_type GoogleChronicleFeed#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Feedname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#feedname GoogleChronicleFeed#feedname}
	Feedname *string `field:"optional" json:"feedname" yaml:"feedname"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

