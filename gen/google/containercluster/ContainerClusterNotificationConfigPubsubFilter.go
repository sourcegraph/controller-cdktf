package containercluster


type ContainerClusterNotificationConfigPubsubFilter struct {
	// Can be used to filter what notifications are sent. Valid values include include UPGRADE_AVAILABLE_EVENT, UPGRADE_EVENT and SECURITY_BULLETIN_EVENT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/container_cluster#event_type ContainerCluster#event_type}
	EventType *[]*string `field:"required" json:"eventType" yaml:"eventType"`
}

