package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction struct {
	// request_header_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#request_header_to_add GoogleNetworkServicesEdgeCacheService#request_header_to_add}
	RequestHeaderToAdd interface{} `field:"optional" json:"requestHeaderToAdd" yaml:"requestHeaderToAdd"`
	// request_header_to_remove block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#request_header_to_remove GoogleNetworkServicesEdgeCacheService#request_header_to_remove}
	RequestHeaderToRemove interface{} `field:"optional" json:"requestHeaderToRemove" yaml:"requestHeaderToRemove"`
	// response_header_to_add block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#response_header_to_add GoogleNetworkServicesEdgeCacheService#response_header_to_add}
	ResponseHeaderToAdd interface{} `field:"optional" json:"responseHeaderToAdd" yaml:"responseHeaderToAdd"`
	// response_header_to_remove block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#response_header_to_remove GoogleNetworkServicesEdgeCacheService#response_header_to_remove}
	ResponseHeaderToRemove interface{} `field:"optional" json:"responseHeaderToRemove" yaml:"responseHeaderToRemove"`
}

