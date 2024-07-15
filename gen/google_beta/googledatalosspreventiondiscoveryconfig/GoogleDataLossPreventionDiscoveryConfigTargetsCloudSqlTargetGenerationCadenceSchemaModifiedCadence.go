package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudSqlTargetGenerationCadenceSchemaModifiedCadence struct {
	// Frequency to regenerate data profiles when the schema is modified. Defaults to monthly. Possible values: ["UPDATE_FREQUENCY_NEVER", "UPDATE_FREQUENCY_DAILY", "UPDATE_FREQUENCY_MONTHLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#frequency GoogleDataLossPreventionDiscoveryConfig#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// The types of schema modifications to consider. Defaults to NEW_COLUMNS. Possible values: ["NEW_COLUMNS", "REMOVED_COLUMNS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#types GoogleDataLossPreventionDiscoveryConfig#types}
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

