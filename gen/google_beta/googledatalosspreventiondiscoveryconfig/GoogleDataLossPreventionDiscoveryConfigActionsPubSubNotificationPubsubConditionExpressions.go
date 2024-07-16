package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotificationPubsubConditionExpressions struct {
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#conditions GoogleDataLossPreventionDiscoveryConfig#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// The operator to apply to the collection of conditions Possible values: ["OR", "AND"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#logical_operator GoogleDataLossPreventionDiscoveryConfig#logical_operator}
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
}

