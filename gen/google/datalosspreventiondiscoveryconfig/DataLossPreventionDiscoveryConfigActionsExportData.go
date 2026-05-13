package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActionsExportData struct {
	// profile_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#profile_table DataLossPreventionDiscoveryConfig#profile_table}
	ProfileTable *DataLossPreventionDiscoveryConfigActionsExportDataProfileTable `field:"optional" json:"profileTable" yaml:"profileTable"`
	// sample_findings_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#sample_findings_table DataLossPreventionDiscoveryConfig#sample_findings_table}
	SampleFindingsTable *DataLossPreventionDiscoveryConfigActionsExportDataSampleFindingsTable `field:"optional" json:"sampleFindingsTable" yaml:"sampleFindingsTable"`
}

