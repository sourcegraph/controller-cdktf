package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigActionsTagResourcesTagConditions struct {
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_data_loss_prevention_discovery_config#sensitivity_score GoogleDataLossPreventionDiscoveryConfig#sensitivity_score}
	SensitivityScore *GoogleDataLossPreventionDiscoveryConfigActionsTagResourcesTagConditionsSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_data_loss_prevention_discovery_config#tag GoogleDataLossPreventionDiscoveryConfig#tag}
	Tag *GoogleDataLossPreventionDiscoveryConfigActionsTagResourcesTagConditionsTag `field:"optional" json:"tag" yaml:"tag"`
}

