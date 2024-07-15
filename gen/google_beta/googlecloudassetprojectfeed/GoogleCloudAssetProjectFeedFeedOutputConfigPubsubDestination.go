package googlecloudassetprojectfeed


type GoogleCloudAssetProjectFeedFeedOutputConfigPubsubDestination struct {
	// Destination on Cloud Pubsub topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_asset_project_feed#topic GoogleCloudAssetProjectFeed#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

