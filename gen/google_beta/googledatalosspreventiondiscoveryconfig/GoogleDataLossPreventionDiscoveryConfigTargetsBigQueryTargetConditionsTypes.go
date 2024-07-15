package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetConditionsTypes struct {
	// A set of BiqQuery table types Possible values: ["BIG_QUERY_TABLE_TYPE_TABLE", "BIG_QUERY_TABLE_TYPE_EXTERNAL_BIG_LAKE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#types GoogleDataLossPreventionDiscoveryConfig#types}
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

