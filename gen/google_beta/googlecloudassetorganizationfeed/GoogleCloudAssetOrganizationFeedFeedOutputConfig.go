package googlecloudassetorganizationfeed


type GoogleCloudAssetOrganizationFeedFeedOutputConfig struct {
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloud_asset_organization_feed#pubsub_destination GoogleCloudAssetOrganizationFeed#pubsub_destination}
	PubsubDestination *GoogleCloudAssetOrganizationFeedFeedOutputConfigPubsubDestination `field:"required" json:"pubsubDestination" yaml:"pubsubDestination"`
}

