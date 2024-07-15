package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction struct {
	// cdn_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#cdn_policy GoogleNetworkServicesEdgeCacheService#cdn_policy}
	CdnPolicy *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicy `field:"optional" json:"cdnPolicy" yaml:"cdnPolicy"`
	// cors_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#cors_policy GoogleNetworkServicesEdgeCacheService#cors_policy}
	CorsPolicy *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCorsPolicy `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_service#url_rewrite GoogleNetworkServicesEdgeCacheService#url_rewrite}
	UrlRewrite *GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
}

