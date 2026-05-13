package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollectionIncludeTagsTagFilters struct {
	// The namespaced name for the tag key.
	//
	// Must be in the format
	// '{parent_id}/{tag_key_short_name}', for example, "123456/sensitive" for
	// an organization parent, or "my-project/sensitive" for a project parent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#namespaced_tag_key DataLossPreventionDiscoveryConfig#namespaced_tag_key}
	NamespacedTagKey *string `field:"optional" json:"namespacedTagKey" yaml:"namespacedTagKey"`
	// The namespaced name for the tag value.
	//
	// Must be in the format
	// '{parent_id}/{tag_key_short_name}/{short_name}', for example,
	// "123456/environment/prod" for an organization parent, or
	// "my-project/environment/prod" for a project parent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/data_loss_prevention_discovery_config#namespaced_tag_value DataLossPreventionDiscoveryConfig#namespaced_tag_value}
	NamespacedTagValue *string `field:"optional" json:"namespacedTagValue" yaml:"namespacedTagValue"`
}

