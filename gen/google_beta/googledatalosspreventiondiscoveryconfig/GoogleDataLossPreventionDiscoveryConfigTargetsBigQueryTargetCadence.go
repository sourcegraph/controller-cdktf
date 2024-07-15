package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadence struct {
	// schema_modified_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#schema_modified_cadence GoogleDataLossPreventionDiscoveryConfig#schema_modified_cadence}
	SchemaModifiedCadence *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadenceSchemaModifiedCadence `field:"optional" json:"schemaModifiedCadence" yaml:"schemaModifiedCadence"`
	// table_modified_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#table_modified_cadence GoogleDataLossPreventionDiscoveryConfig#table_modified_cadence}
	TableModifiedCadence *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadenceTableModifiedCadence `field:"optional" json:"tableModifiedCadence" yaml:"tableModifiedCadence"`
}

