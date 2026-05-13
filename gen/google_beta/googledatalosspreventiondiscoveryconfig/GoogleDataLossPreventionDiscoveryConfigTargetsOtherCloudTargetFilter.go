package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilter struct {
	// collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#collection GoogleDataLossPreventionDiscoveryConfig#collection}
	Collection *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterCollection `field:"optional" json:"collection" yaml:"collection"`
	// others block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#others GoogleDataLossPreventionDiscoveryConfig#others}
	Others *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterOthers `field:"optional" json:"others" yaml:"others"`
	// single_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#single_resource GoogleDataLossPreventionDiscoveryConfig#single_resource}
	SingleResource *GoogleDataLossPreventionDiscoveryConfigTargetsOtherCloudTargetFilterSingleResource `field:"optional" json:"singleResource" yaml:"singleResource"`
}

