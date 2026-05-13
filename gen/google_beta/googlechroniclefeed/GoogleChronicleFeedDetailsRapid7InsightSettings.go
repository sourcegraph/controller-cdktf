package googlechroniclefeed


type GoogleChronicleFeedDetailsRapid7InsightSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsRapid7InsightSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Rapid7 API endpoint. Should be "vulnerabilities" or "assets".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#endpoint GoogleChronicleFeed#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
}

