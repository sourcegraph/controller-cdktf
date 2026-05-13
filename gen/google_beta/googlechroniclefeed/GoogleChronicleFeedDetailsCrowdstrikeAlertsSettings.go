package googlechroniclefeed


type GoogleChronicleFeedDetailsCrowdstrikeAlertsSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsCrowdstrikeAlertsSettingsAuthentication `field:"required" json:"authentication" yaml:"authentication"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Ingestion Type. Possible values: BRING_ALL_ALERTS BRING_ONLY_NEW_ALERTS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#ingestion_type GoogleChronicleFeed#ingestion_type}
	IngestionType *string `field:"optional" json:"ingestionType" yaml:"ingestionType"`
}

