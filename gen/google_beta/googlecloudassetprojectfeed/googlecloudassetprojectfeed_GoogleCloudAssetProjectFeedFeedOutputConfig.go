package googlecloudassetprojectfeed


type GoogleCloudAssetProjectFeedFeedOutputConfig struct {
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_asset_project_feed#pubsub_destination GoogleCloudAssetProjectFeed#pubsub_destination}
	PubsubDestination *GoogleCloudAssetProjectFeedFeedOutputConfigPubsubDestination `field:"required" json:"pubsubDestination" yaml:"pubsubDestination"`
}

