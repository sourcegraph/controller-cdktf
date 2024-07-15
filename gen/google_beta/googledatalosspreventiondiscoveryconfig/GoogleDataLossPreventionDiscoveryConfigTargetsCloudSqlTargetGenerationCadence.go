package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudSqlTargetGenerationCadence struct {
	// Data changes (non-schema changes) in Cloud SQL tables can't trigger reprofiling.
	//
	// If you set this field, profiles are refreshed at this frequency regardless of whether the underlying tables have changes. Defaults to never. Possible values: ["UPDATE_FREQUENCY_NEVER", "UPDATE_FREQUENCY_DAILY", "UPDATE_FREQUENCY_MONTHLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#refresh_frequency GoogleDataLossPreventionDiscoveryConfig#refresh_frequency}
	RefreshFrequency *string `field:"optional" json:"refreshFrequency" yaml:"refreshFrequency"`
	// schema_modified_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#schema_modified_cadence GoogleDataLossPreventionDiscoveryConfig#schema_modified_cadence}
	SchemaModifiedCadence *GoogleDataLossPreventionDiscoveryConfigTargetsCloudSqlTargetGenerationCadenceSchemaModifiedCadence `field:"optional" json:"schemaModifiedCadence" yaml:"schemaModifiedCadence"`
}

