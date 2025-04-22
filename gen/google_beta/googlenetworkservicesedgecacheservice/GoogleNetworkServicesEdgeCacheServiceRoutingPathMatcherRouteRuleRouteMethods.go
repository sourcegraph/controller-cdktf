package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteMethods struct {
	// The non-empty set of HTTP methods that are allowed for this route.
	//
	// Any combination of "GET", "HEAD", "OPTIONS", "PUT", "POST", "DELETE", and "PATCH".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_network_services_edge_cache_service#allowed_methods GoogleNetworkServicesEdgeCacheService#allowed_methods}
	AllowedMethods *[]*string `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
}

