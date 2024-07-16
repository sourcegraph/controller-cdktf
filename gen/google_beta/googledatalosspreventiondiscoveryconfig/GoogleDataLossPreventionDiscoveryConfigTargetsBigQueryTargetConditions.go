package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetConditions struct {
	// A timestamp in RFC3339 UTC "Zulu" format with nanosecond resolution and upto nine fractional digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#created_after GoogleDataLossPreventionDiscoveryConfig#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// or_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#or_conditions GoogleDataLossPreventionDiscoveryConfig#or_conditions}
	OrConditions *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetConditionsOrConditions `field:"optional" json:"orConditions" yaml:"orConditions"`
	// Restrict discovery to categories of table types.
	//
	// Currently view, materialized view, snapshot and non-biglake external tables are supported. Possible values: ["BIG_QUERY_COLLECTION_ALL_TYPES", "BIG_QUERY_COLLECTION_ONLY_SUPPORTED_TYPES"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#type_collection GoogleDataLossPreventionDiscoveryConfig#type_collection}
	TypeCollection *string `field:"optional" json:"typeCollection" yaml:"typeCollection"`
	// types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_data_loss_prevention_discovery_config#types GoogleDataLossPreventionDiscoveryConfig#types}
	Types *GoogleDataLossPreventionDiscoveryConfigTargetsBigQueryTargetConditionsTypes `field:"optional" json:"types" yaml:"types"`
}

