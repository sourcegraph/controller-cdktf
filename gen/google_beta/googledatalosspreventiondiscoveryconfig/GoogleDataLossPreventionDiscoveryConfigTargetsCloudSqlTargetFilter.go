package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilter struct {
	// collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_data_loss_prevention_discovery_config#collection GoogleDataLossPreventionDiscoveryConfig#collection}
	Collection *GoogleDataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterCollection `field:"optional" json:"collection" yaml:"collection"`
	// others block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.30.0/docs/resources/google_data_loss_prevention_discovery_config#others GoogleDataLossPreventionDiscoveryConfig#others}
	Others *GoogleDataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterOthers `field:"optional" json:"others" yaml:"others"`
}

