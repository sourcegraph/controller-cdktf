package googlechroniclefeed


type GoogleChronicleFeedDetailsTrellixHxAlertsSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsTrellixHxAlertsSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Trellix HX Device URL.
	//
	// This must be a valid URL with an http or https scheme. It has no default.
	// Usually a device URL is in the form of either:
	// https://xxx.trellix.com/hx/id//
	// - or -
	// https://htapdeviceproxy.md.mandiant.net/dphb/hx//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#endpoint GoogleChronicleFeed#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

