package googlecloudassetfolderfeed


type GoogleCloudAssetFolderFeedFeedOutputConfig struct {
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_asset_folder_feed#pubsub_destination GoogleCloudAssetFolderFeed#pubsub_destination}
	PubsubDestination *GoogleCloudAssetFolderFeedFeedOutputConfigPubsubDestination `field:"required" json:"pubsubDestination" yaml:"pubsubDestination"`
}

