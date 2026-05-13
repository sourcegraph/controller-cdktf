package googlechroniclefeed


type GoogleChronicleFeedDetailsSentineloneAlertSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsSentineloneAlertSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Hostname of SentinelOne alert settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// initialStartTime from when to fetch the alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#initial_start_time GoogleChronicleFeed#initial_start_time}
	InitialStartTime *string `field:"optional" json:"initialStartTime" yaml:"initialStartTime"`
	// Is the customer subscribed to Alerts Api.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#is_alert_api_subscribed GoogleChronicleFeed#is_alert_api_subscribed}
	IsAlertApiSubscribed interface{} `field:"optional" json:"isAlertApiSubscribed" yaml:"isAlertApiSubscribed"`
}

