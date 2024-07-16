package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterDatabaseResourceReference struct {
	// Required. Name of a database within the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#database GoogleDataLossPreventionDiscoveryConfig#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// Required. Name of a database resource, for example, a table within the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#database_resource GoogleDataLossPreventionDiscoveryConfig#database_resource}
	DatabaseResource *string `field:"required" json:"databaseResource" yaml:"databaseResource"`
	// Required. The instance where this resource is located. For example: Cloud SQL instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#instance GoogleDataLossPreventionDiscoveryConfig#instance}
	Instance *string `field:"required" json:"instance" yaml:"instance"`
	// Required. If within a project-level config, then this must match the config's project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#project_id GoogleDataLossPreventionDiscoveryConfig#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

