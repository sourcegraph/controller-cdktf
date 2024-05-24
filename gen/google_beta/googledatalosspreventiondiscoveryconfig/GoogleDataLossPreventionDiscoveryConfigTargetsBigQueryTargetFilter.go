package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilter struct {
	// other_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_data_loss_prevention_discovery_config#other_tables GoogleDataLossPreventionDiscoveryConfig#other_tables}
	OtherTables *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTables `field:"optional" json:"otherTables" yaml:"otherTables"`
	// tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_data_loss_prevention_discovery_config#tables GoogleDataLossPreventionDiscoveryConfig#tables}
	Tables *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTables `field:"optional" json:"tables" yaml:"tables"`
}

