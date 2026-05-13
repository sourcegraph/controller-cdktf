package googlechroniclefeed


type GoogleChronicleFeedDetailsAzureBlobStoreV2Settings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsAzureBlobStoreV2SettingsAuthentication `field:"required" json:"authentication" yaml:"authentication"`
	// Azure URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#azure_uri GoogleChronicleFeed#azure_uri}
	AzureUri *string `field:"required" json:"azureUri" yaml:"azureUri"`
	// Maximum File Age to ingest in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#max_lookback_days GoogleChronicleFeed#max_lookback_days}
	MaxLookbackDays *float64 `field:"optional" json:"maxLookbackDays" yaml:"maxLookbackDays"`
	// Possible values: NEVER ON_SUCCESS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#source_deletion_option GoogleChronicleFeed#source_deletion_option}
	SourceDeletionOption *string `field:"optional" json:"sourceDeletionOption" yaml:"sourceDeletionOption"`
}

