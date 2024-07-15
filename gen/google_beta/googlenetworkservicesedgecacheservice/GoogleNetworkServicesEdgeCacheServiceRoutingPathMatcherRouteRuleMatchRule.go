package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRule struct {
	// For satisfying the matchRule condition, the path of the request must exactly match the value specified in fullPathMatch after removing any query parameters and anchor that may be part of the original URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#full_path_match GoogleNetworkServicesEdgeCacheService#full_path_match}
	FullPathMatch *string `field:"optional" json:"fullPathMatch" yaml:"fullPathMatch"`
	// header_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#header_match GoogleNetworkServicesEdgeCacheService#header_match}
	HeaderMatch interface{} `field:"optional" json:"headerMatch" yaml:"headerMatch"`
	// Specifies that prefixMatch and fullPathMatch matches are case sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#ignore_case GoogleNetworkServicesEdgeCacheService#ignore_case}
	IgnoreCase interface{} `field:"optional" json:"ignoreCase" yaml:"ignoreCase"`
	// For satisfying the matchRule condition, the path of the request must match the wildcard pattern specified in pathTemplateMatch after removing any query parameters and anchor that may be part of the original URL.
	//
	// pathTemplateMatch must be between 1 and 255 characters
	// (inclusive).  The pattern specified by pathTemplateMatch may
	// have at most 5 wildcard operators and at most 5 variable
	// captures in total.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#path_template_match GoogleNetworkServicesEdgeCacheService#path_template_match}
	PathTemplateMatch *string `field:"optional" json:"pathTemplateMatch" yaml:"pathTemplateMatch"`
	// For satisfying the matchRule condition, the request's path must begin with the specified prefixMatch.
	//
	// prefixMatch must begin with a /.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#prefix_match GoogleNetworkServicesEdgeCacheService#prefix_match}
	PrefixMatch *string `field:"optional" json:"prefixMatch" yaml:"prefixMatch"`
	// query_parameter_match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#query_parameter_match GoogleNetworkServicesEdgeCacheService#query_parameter_match}
	QueryParameterMatch interface{} `field:"optional" json:"queryParameterMatch" yaml:"queryParameterMatch"`
}

