package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence struct {
	// inspect_template_modified_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#inspect_template_modified_cadence GoogleDataLossPreventionDiscoveryConfig#inspect_template_modified_cadence}
	InspectTemplateModifiedCadence *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadenceInspectTemplateModifiedCadence `field:"optional" json:"inspectTemplateModifiedCadence" yaml:"inspectTemplateModifiedCadence"`
	// Frequency to update profiles regardless of whether the underlying resource has changes.
	//
	// Defaults to never. Possible values: ["UPDATE_FREQUENCY_NEVER", "UPDATE_FREQUENCY_DAILY", "UPDATE_FREQUENCY_MONTHLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#refresh_frequency GoogleDataLossPreventionDiscoveryConfig#refresh_frequency}
	RefreshFrequency *string `field:"optional" json:"refreshFrequency" yaml:"refreshFrequency"`
}

