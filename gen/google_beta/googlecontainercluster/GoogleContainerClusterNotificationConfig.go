package googlecontainercluster


type GoogleContainerClusterNotificationConfig struct {
	// pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_container_cluster#pubsub GoogleContainerCluster#pubsub}
	Pubsub *GoogleContainerClusterNotificationConfigPubsub `field:"required" json:"pubsub" yaml:"pubsub"`
}

