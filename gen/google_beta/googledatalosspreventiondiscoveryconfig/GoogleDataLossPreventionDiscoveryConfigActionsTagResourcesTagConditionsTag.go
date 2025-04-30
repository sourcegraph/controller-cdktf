package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigActionsTagResourcesTagConditionsTag struct {
	// The namespaced name for the tag value to attach to resources.
	//
	// Must be in the format '{parent_id}/{tag_key_short_name}/{short_name}', for example, "123456/environment/prod".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_data_loss_prevention_discovery_config#namespaced_value GoogleDataLossPreventionDiscoveryConfig#namespaced_value}
	NamespacedValue *string `field:"optional" json:"namespacedValue" yaml:"namespacedValue"`
}

