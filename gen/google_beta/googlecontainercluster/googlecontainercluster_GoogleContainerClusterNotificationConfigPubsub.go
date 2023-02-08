package googlecontainercluster


type GoogleContainerClusterNotificationConfigPubsub struct {
	// Whether or not the notification config is enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The pubsub topic to push upgrade notifications to.
	//
	// Must be in the same project as the cluster. Must be in the format: projects/{project}/topics/{topic}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#topic GoogleContainerCluster#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

