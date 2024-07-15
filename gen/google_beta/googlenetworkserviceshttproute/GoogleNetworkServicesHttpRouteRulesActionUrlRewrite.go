package googlenetworkserviceshttproute


type GoogleNetworkServicesHttpRouteRulesActionUrlRewrite struct {
	// Prior to forwarding the request to the selected destination, the requests host header is replaced by this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#host_rewrite GoogleNetworkServicesHttpRoute#host_rewrite}
	HostRewrite *string `field:"optional" json:"hostRewrite" yaml:"hostRewrite"`
	// Prior to forwarding the request to the selected destination, the matching portion of the requests path is replaced by this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_http_route#path_prefix_rewrite GoogleNetworkServicesHttpRoute#path_prefix_rewrite}
	PathPrefixRewrite *string `field:"optional" json:"pathPrefixRewrite" yaml:"pathPrefixRewrite"`
}

