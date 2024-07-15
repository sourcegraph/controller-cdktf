package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigActions struct {
	// export_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#export_data GoogleDataLossPreventionDiscoveryConfig#export_data}
	ExportData *GoogleDataLossPreventionDiscoveryConfigActionsExportData `field:"optional" json:"exportData" yaml:"exportData"`
	// pub_sub_notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#pub_sub_notification GoogleDataLossPreventionDiscoveryConfig#pub_sub_notification}
	PubSubNotification *GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotification `field:"optional" json:"pubSubNotification" yaml:"pubSubNotification"`
}

