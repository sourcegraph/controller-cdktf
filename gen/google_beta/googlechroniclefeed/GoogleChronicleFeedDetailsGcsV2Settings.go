package googlechroniclefeed


type GoogleChronicleFeedDetailsGcsV2Settings struct {
	// Google Cloud Storage Bucket URI for the feed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#bucket_uri GoogleChronicleFeed#bucket_uri}
	BucketUri *string `field:"required" json:"bucketUri" yaml:"bucketUri"`
	// Maximum File Age to ingest in days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#max_lookback_days GoogleChronicleFeed#max_lookback_days}
	MaxLookbackDays *float64 `field:"optional" json:"maxLookbackDays" yaml:"maxLookbackDays"`
	// Possible values: NEVER ON_SUCCESS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_chronicle_feed#source_deletion_option GoogleChronicleFeed#source_deletion_option}
	SourceDeletionOption *string `field:"optional" json:"sourceDeletionOption" yaml:"sourceDeletionOption"`
}

