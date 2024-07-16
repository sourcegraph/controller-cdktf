package googlecloudassetfolderfeed


type GoogleCloudAssetFolderFeedFeedOutputConfigPubsubDestination struct {
	// Destination on Cloud Pubsub topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_asset_folder_feed#topic GoogleCloudAssetFolderFeed#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

