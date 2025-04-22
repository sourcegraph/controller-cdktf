package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetGenerationCadence struct {
	// inspect_template_modified_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_data_loss_prevention_discovery_config#inspect_template_modified_cadence GoogleDataLossPreventionDiscoveryConfig#inspect_template_modified_cadence}
	InspectTemplateModifiedCadence *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetGenerationCadenceInspectTemplateModifiedCadence `field:"optional" json:"inspectTemplateModifiedCadence" yaml:"inspectTemplateModifiedCadence"`
	// Data changes in Cloud Storage can't trigger reprofiling.
	//
	// If you set this field, profiles are refreshed at this frequency regardless of whether the underlying buckets have changes. Defaults to never. Possible values: ["UPDATE_FREQUENCY_NEVER", "UPDATE_FREQUENCY_DAILY", "UPDATE_FREQUENCY_MONTHLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_data_loss_prevention_discovery_config#refresh_frequency GoogleDataLossPreventionDiscoveryConfig#refresh_frequency}
	RefreshFrequency *string `field:"optional" json:"refreshFrequency" yaml:"refreshFrequency"`
}

