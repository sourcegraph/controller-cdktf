package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter struct {
	// collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#collection DataLossPreventionDiscoveryConfig#collection}
	Collection *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollection `field:"optional" json:"collection" yaml:"collection"`
	// others block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#others DataLossPreventionDiscoveryConfig#others}
	Others *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOthers `field:"optional" json:"others" yaml:"others"`
	// single_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#single_resource DataLossPreventionDiscoveryConfig#single_resource}
	SingleResource *DataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResource `field:"optional" json:"singleResource" yaml:"singleResource"`
}

