package googlecomputeregionurlmap


type GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFilters struct {
	// filter_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#filter_labels GoogleComputeRegionUrlMap#filter_labels}
	FilterLabels interface{} `field:"required" json:"filterLabels" yaml:"filterLabels"`
	// Specifies how individual filterLabel matches within the list of filterLabels contribute towards the overall metadataFilter match. Supported values are:.
	//
	// * MATCH_ANY: At least one of the filterLabels must have a matching label in the
	// provided metadata.
	// * MATCH_ALL: All filterLabels must have matching labels in
	// the provided metadata. Possible values: ["MATCH_ALL", "MATCH_ANY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_region_url_map#filter_match_criteria GoogleComputeRegionUrlMap#filter_match_criteria}
	FilterMatchCriteria *string `field:"required" json:"filterMatchCriteria" yaml:"filterMatchCriteria"`
}

