package networkservicesedgecacheservice

type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction struct {
	// request_header_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_services_edge_cache_service#request_header_to_add NetworkServicesEdgeCacheService#request_header_to_add}
	RequestHeaderToAdd any `field:"optional" json:"requestHeaderToAdd" yaml:"requestHeaderToAdd"`
	// request_header_to_remove block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_services_edge_cache_service#request_header_to_remove NetworkServicesEdgeCacheService#request_header_to_remove}
	RequestHeaderToRemove any `field:"optional" json:"requestHeaderToRemove" yaml:"requestHeaderToRemove"`
	// response_header_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_services_edge_cache_service#response_header_to_add NetworkServicesEdgeCacheService#response_header_to_add}
	ResponseHeaderToAdd any `field:"optional" json:"responseHeaderToAdd" yaml:"responseHeaderToAdd"`
	// response_header_to_remove block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_services_edge_cache_service#response_header_to_remove NetworkServicesEdgeCacheService#response_header_to_remove}
	ResponseHeaderToRemove any `field:"optional" json:"responseHeaderToRemove" yaml:"responseHeaderToRemove"`
}
