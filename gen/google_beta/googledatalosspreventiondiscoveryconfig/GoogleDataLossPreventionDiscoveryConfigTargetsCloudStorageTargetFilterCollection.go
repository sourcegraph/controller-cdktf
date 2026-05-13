package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollection struct {
	// include_regexes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#include_regexes GoogleDataLossPreventionDiscoveryConfig#include_regexes}
	IncludeRegexes *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollectionIncludeRegexes `field:"optional" json:"includeRegexes" yaml:"includeRegexes"`
	// include_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_data_loss_prevention_discovery_config#include_tags GoogleDataLossPreventionDiscoveryConfig#include_tags}
	IncludeTags *GoogleDataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollectionIncludeTags `field:"optional" json:"includeTags" yaml:"includeTags"`
}

