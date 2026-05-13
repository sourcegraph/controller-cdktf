package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigActions struct {
	// export_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#export_data GoogleDataLossPreventionDiscoveryConfig#export_data}
	ExportData *GoogleDataLossPreventionDiscoveryConfigActionsExportData `field:"optional" json:"exportData" yaml:"exportData"`
	// publish_to_chronicle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#publish_to_chronicle GoogleDataLossPreventionDiscoveryConfig#publish_to_chronicle}
	PublishToChronicle *GoogleDataLossPreventionDiscoveryConfigActionsPublishToChronicle `field:"optional" json:"publishToChronicle" yaml:"publishToChronicle"`
	// publish_to_dataplex_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#publish_to_dataplex_catalog GoogleDataLossPreventionDiscoveryConfig#publish_to_dataplex_catalog}
	PublishToDataplexCatalog *GoogleDataLossPreventionDiscoveryConfigActionsPublishToDataplexCatalog `field:"optional" json:"publishToDataplexCatalog" yaml:"publishToDataplexCatalog"`
	// publish_to_scc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#publish_to_scc GoogleDataLossPreventionDiscoveryConfig#publish_to_scc}
	PublishToScc *GoogleDataLossPreventionDiscoveryConfigActionsPublishToScc `field:"optional" json:"publishToScc" yaml:"publishToScc"`
	// pub_sub_notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#pub_sub_notification GoogleDataLossPreventionDiscoveryConfig#pub_sub_notification}
	PubSubNotification *GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotification `field:"optional" json:"pubSubNotification" yaml:"pubSubNotification"`
	// tag_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#tag_resources GoogleDataLossPreventionDiscoveryConfig#tag_resources}
	TagResources *GoogleDataLossPreventionDiscoveryConfigActionsTagResources `field:"optional" json:"tagResources" yaml:"tagResources"`
}

