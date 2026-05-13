package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTarget struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#filter GoogleDataLossPreventionDiscoveryConfig#filter}
	Filter *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter `field:"required" json:"filter" yaml:"filter"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#conditions GoogleDataLossPreventionDiscoveryConfig#conditions}
	Conditions *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// data_source_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#data_source_type GoogleDataLossPreventionDiscoveryConfig#data_source_type}
	DataSourceType *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType `field:"optional" json:"dataSourceType" yaml:"dataSourceType"`
	// disabled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#disabled GoogleDataLossPreventionDiscoveryConfig#disabled}
	Disabled *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled `field:"optional" json:"disabled" yaml:"disabled"`
	// generation_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#generation_cadence GoogleDataLossPreventionDiscoveryConfig#generation_cadence}
	GenerationCadence *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence `field:"optional" json:"generationCadence" yaml:"generationCadence"`
}

