package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActions struct {
	// export_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#export_data DataLossPreventionDiscoveryConfig#export_data}
	ExportData *DataLossPreventionDiscoveryConfigActionsExportData `field:"optional" json:"exportData" yaml:"exportData"`
	// publish_to_chronicle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#publish_to_chronicle DataLossPreventionDiscoveryConfig#publish_to_chronicle}
	PublishToChronicle *DataLossPreventionDiscoveryConfigActionsPublishToChronicle `field:"optional" json:"publishToChronicle" yaml:"publishToChronicle"`
	// publish_to_dataplex_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#publish_to_dataplex_catalog DataLossPreventionDiscoveryConfig#publish_to_dataplex_catalog}
	PublishToDataplexCatalog *DataLossPreventionDiscoveryConfigActionsPublishToDataplexCatalog `field:"optional" json:"publishToDataplexCatalog" yaml:"publishToDataplexCatalog"`
	// publish_to_scc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#publish_to_scc DataLossPreventionDiscoveryConfig#publish_to_scc}
	PublishToScc *DataLossPreventionDiscoveryConfigActionsPublishToScc `field:"optional" json:"publishToScc" yaml:"publishToScc"`
	// pub_sub_notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#pub_sub_notification DataLossPreventionDiscoveryConfig#pub_sub_notification}
	PubSubNotification *DataLossPreventionDiscoveryConfigActionsPubSubNotification `field:"optional" json:"pubSubNotification" yaml:"pubSubNotification"`
	// tag_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#tag_resources DataLossPreventionDiscoveryConfig#tag_resources}
	TagResources *DataLossPreventionDiscoveryConfigActionsTagResources `field:"optional" json:"tagResources" yaml:"tagResources"`
}

