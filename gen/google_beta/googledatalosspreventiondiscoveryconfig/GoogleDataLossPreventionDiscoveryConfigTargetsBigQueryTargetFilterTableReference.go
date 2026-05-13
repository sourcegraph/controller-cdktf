package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReference struct {
	// Dataset ID of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#dataset_id GoogleDataLossPreventionDiscoveryConfig#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// Name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#table_id GoogleDataLossPreventionDiscoveryConfig#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
	// The Google Cloud project ID of the project containing the table.
	//
	// If omitted, the project ID is inferred from the parent project.
	// This field is required if the parent resource is an organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#project_id GoogleDataLossPreventionDiscoveryConfig#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

