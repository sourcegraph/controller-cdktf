package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTarget struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_data_loss_prevention_discovery_config#filter GoogleDataLossPreventionDiscoveryConfig#filter}
	Filter *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilter `field:"required" json:"filter" yaml:"filter"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_data_loss_prevention_discovery_config#conditions GoogleDataLossPreventionDiscoveryConfig#conditions}
	Conditions *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// disabled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_data_loss_prevention_discovery_config#disabled GoogleDataLossPreventionDiscoveryConfig#disabled}
	Disabled *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetDisabled `field:"optional" json:"disabled" yaml:"disabled"`
	// generation_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_data_loss_prevention_discovery_config#generation_cadence GoogleDataLossPreventionDiscoveryConfig#generation_cadence}
	GenerationCadence *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetGenerationCadence `field:"optional" json:"generationCadence" yaml:"generationCadence"`
}

