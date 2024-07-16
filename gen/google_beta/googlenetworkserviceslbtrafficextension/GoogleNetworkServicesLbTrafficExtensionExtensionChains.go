package googlenetworkserviceslbtrafficextension


type GoogleNetworkServicesLbTrafficExtensionExtensionChains struct {
	// extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_lb_traffic_extension#extensions GoogleNetworkServicesLbTrafficExtension#extensions}
	Extensions interface{} `field:"required" json:"extensions" yaml:"extensions"`
	// match_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_lb_traffic_extension#match_condition GoogleNetworkServicesLbTrafficExtension#match_condition}
	MatchCondition *GoogleNetworkServicesLbTrafficExtensionExtensionChainsMatchCondition `field:"required" json:"matchCondition" yaml:"matchCondition"`
	// The name for this extension chain.
	//
	// The name is logged as part of the HTTP request logs.
	// The name must conform with RFC-1034, is restricted to lower-cased letters, numbers and hyphens,
	// and can have a maximum length of 63 characters. Additionally, the first character must be a letter
	// and the last a letter or a number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_lb_traffic_extension#name GoogleNetworkServicesLbTrafficExtension#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

