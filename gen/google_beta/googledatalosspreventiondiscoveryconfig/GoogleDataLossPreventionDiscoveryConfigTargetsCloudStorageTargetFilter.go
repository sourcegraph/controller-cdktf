package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilter struct {
	// cloud_storage_resource_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_data_loss_prevention_discovery_config#cloud_storage_resource_reference GoogleDataLossPreventionDiscoveryConfig#cloud_storage_resource_reference}
	CloudStorageResourceReference *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReference `field:"optional" json:"cloudStorageResourceReference" yaml:"cloudStorageResourceReference"`
	// collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_data_loss_prevention_discovery_config#collection GoogleDataLossPreventionDiscoveryConfig#collection}
	Collection *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollection `field:"optional" json:"collection" yaml:"collection"`
	// others block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_data_loss_prevention_discovery_config#others GoogleDataLossPreventionDiscoveryConfig#others}
	Others *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOthers `field:"optional" json:"others" yaml:"others"`
}

