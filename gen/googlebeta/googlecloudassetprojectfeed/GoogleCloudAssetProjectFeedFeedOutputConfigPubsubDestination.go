package googlecloudassetprojectfeed


type GoogleCloudAssetProjectFeedFeedOutputConfigPubsubDestination struct {
	// Destination on Cloud Pubsub topic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_asset_project_feed#topic GoogleCloudAssetProjectFeed#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

