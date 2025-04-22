package googlecontainercluster


type GoogleContainerClusterNotificationConfigPubsubFilter struct {
	// Can be used to filter what notifications are sent. Valid values include include UPGRADE_AVAILABLE_EVENT, UPGRADE_EVENT, SECURITY_BULLETIN_EVENT, and UPGRADE_INFO_EVENT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_container_cluster#event_type GoogleContainerCluster#event_type}
	EventType *[]*string `field:"required" json:"eventType" yaml:"eventType"`
}

