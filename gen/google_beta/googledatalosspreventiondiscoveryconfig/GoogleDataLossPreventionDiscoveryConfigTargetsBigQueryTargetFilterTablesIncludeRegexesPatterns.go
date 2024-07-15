package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTablesIncludeRegexesPatterns struct {
	// if unset, this property matches all datasets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#dataset_id_regex GoogleDataLossPreventionDiscoveryConfig#dataset_id_regex}
	DatasetIdRegex *string `field:"optional" json:"datasetIdRegex" yaml:"datasetIdRegex"`
	// For organizations, if unset, will match all projects. Has no effect for data profile configurations created within a project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#project_id_regex GoogleDataLossPreventionDiscoveryConfig#project_id_regex}
	ProjectIdRegex *string `field:"optional" json:"projectIdRegex" yaml:"projectIdRegex"`
	// if unset, this property matches all tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#table_id_regex GoogleDataLossPreventionDiscoveryConfig#table_id_regex}
	TableIdRegex *string `field:"optional" json:"tableIdRegex" yaml:"tableIdRegex"`
}

