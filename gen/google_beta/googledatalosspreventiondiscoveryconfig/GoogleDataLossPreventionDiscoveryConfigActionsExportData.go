package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigActionsExportData struct {
	// profile_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#profile_table GoogleDataLossPreventionDiscoveryConfig#profile_table}
	ProfileTable *GoogleDataLossPreventionDiscoveryConfigActionsExportDataProfileTable `field:"optional" json:"profileTable" yaml:"profileTable"`
	// sample_findings_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#sample_findings_table GoogleDataLossPreventionDiscoveryConfig#sample_findings_table}
	SampleFindingsTable *GoogleDataLossPreventionDiscoveryConfigActionsExportDataSampleFindingsTable `field:"optional" json:"sampleFindingsTable" yaml:"sampleFindingsTable"`
}

