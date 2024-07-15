package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleHeaderMatch struct {
	// The header name to match on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#header_name GoogleNetworkServicesEdgeCacheService#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The value of the header should exactly match contents of exactMatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#exact_match GoogleNetworkServicesEdgeCacheService#exact_match}
	ExactMatch *string `field:"optional" json:"exactMatch" yaml:"exactMatch"`
	// If set to false (default), the headerMatch is considered a match if the match criteria above are met.
	//
	// If set to true, the headerMatch is considered a match if the match criteria above are NOT met.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#invert_match GoogleNetworkServicesEdgeCacheService#invert_match}
	InvertMatch interface{} `field:"optional" json:"invertMatch" yaml:"invertMatch"`
	// The value of the header must start with the contents of prefixMatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#prefix_match GoogleNetworkServicesEdgeCacheService#prefix_match}
	PrefixMatch *string `field:"optional" json:"prefixMatch" yaml:"prefixMatch"`
	// A header with the contents of headerName must exist.
	//
	// The match takes place whether or not the request's header has a value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#present_match GoogleNetworkServicesEdgeCacheService#present_match}
	PresentMatch interface{} `field:"optional" json:"presentMatch" yaml:"presentMatch"`
	// The value of the header must end with the contents of suffixMatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#suffix_match GoogleNetworkServicesEdgeCacheService#suffix_match}
	SuffixMatch *string `field:"optional" json:"suffixMatch" yaml:"suffixMatch"`
}

