package googlechroniclefeed


type GoogleChronicleFeedDetailsAmazonS3V2Settings struct {
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#authentication GoogleChronicleFeed#authentication}
	Authentication *GoogleChronicleFeedDetailsAmazonS3V2SettingsAuthentication `field:"required" json:"authentication" yaml:"authentication"`
	// S3 URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#s3_uri GoogleChronicleFeed#s3_uri}
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Maximum File Age to ingest in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#max_lookback_days GoogleChronicleFeed#max_lookback_days}
	MaxLookbackDays *float64 `field:"optional" json:"maxLookbackDays" yaml:"maxLookbackDays"`
	// Possible values: NEVER ON_SUCCESS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#source_deletion_option GoogleChronicleFeed#source_deletion_option}
	SourceDeletionOption *string `field:"optional" json:"sourceDeletionOption" yaml:"sourceDeletionOption"`
}

