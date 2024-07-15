package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceRouting struct {
	// host_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#host_rule GoogleNetworkServicesEdgeCacheService#host_rule}
	HostRule interface{} `field:"required" json:"hostRule" yaml:"hostRule"`
	// path_matcher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#path_matcher GoogleNetworkServicesEdgeCacheService#path_matcher}
	PathMatcher interface{} `field:"required" json:"pathMatcher" yaml:"pathMatcher"`
}

