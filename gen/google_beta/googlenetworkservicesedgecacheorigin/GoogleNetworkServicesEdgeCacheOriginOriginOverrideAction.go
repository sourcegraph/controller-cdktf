package googlenetworkservicesedgecacheorigin


type GoogleNetworkServicesEdgeCacheOriginOriginOverrideAction struct {
	// header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_origin#header_action GoogleNetworkServicesEdgeCacheOrigin#header_action}
	HeaderAction *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionHeaderAction `field:"optional" json:"headerAction" yaml:"headerAction"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_edge_cache_origin#url_rewrite GoogleNetworkServicesEdgeCacheOrigin#url_rewrite}
	UrlRewrite *GoogleNetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
}

