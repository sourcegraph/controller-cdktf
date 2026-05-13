package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsOtherCloudTarget struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#filter DataLossPreventionDiscoveryConfig#filter}
	Filter *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter `field:"required" json:"filter" yaml:"filter"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#conditions DataLossPreventionDiscoveryConfig#conditions}
	Conditions *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// data_source_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#data_source_type DataLossPreventionDiscoveryConfig#data_source_type}
	DataSourceType *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDataSourceType `field:"optional" json:"dataSourceType" yaml:"dataSourceType"`
	// disabled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#disabled DataLossPreventionDiscoveryConfig#disabled}
	Disabled *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetDisabled `field:"optional" json:"disabled" yaml:"disabled"`
	// generation_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#generation_cadence DataLossPreventionDiscoveryConfig#generation_cadence}
	GenerationCadence *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetGenerationCadence `field:"optional" json:"generationCadence" yaml:"generationCadence"`
}

