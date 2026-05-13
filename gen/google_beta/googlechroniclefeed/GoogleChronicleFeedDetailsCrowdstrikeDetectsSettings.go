package googlechroniclefeed


type GoogleChronicleFeedDetailsCrowdstrikeDetectsSettings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsCrowdstrikeDetectsSettingsAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// API Hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#hostname GoogleChronicleFeed#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Ingestion Type. Possible values: BRING_ONLY_NEW_DETECTIONS BRING_ALL_DETECTIONS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#ingestion_type GoogleChronicleFeed#ingestion_type}
	IngestionType *string `field:"optional" json:"ingestionType" yaml:"ingestionType"`
}

