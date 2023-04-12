package googlecontainercluster


type GoogleContainerClusterNotificationConfig struct {
	// pubsub block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#pubsub GoogleContainerCluster#pubsub}
	Pubsub *GoogleContainerClusterNotificationConfigPubsub `field:"required" json:"pubsub" yaml:"pubsub"`
}

