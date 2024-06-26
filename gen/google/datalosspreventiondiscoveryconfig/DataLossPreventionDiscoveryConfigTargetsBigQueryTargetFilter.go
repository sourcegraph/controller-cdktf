package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilter struct {
	// other_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/data_loss_prevention_discovery_config#other_tables DataLossPreventionDiscoveryConfig#other_tables}
	OtherTables *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTables `field:"optional" json:"otherTables" yaml:"otherTables"`
	// tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/data_loss_prevention_discovery_config#tables DataLossPreventionDiscoveryConfig#tables}
	Tables *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTables `field:"optional" json:"tables" yaml:"tables"`
}

