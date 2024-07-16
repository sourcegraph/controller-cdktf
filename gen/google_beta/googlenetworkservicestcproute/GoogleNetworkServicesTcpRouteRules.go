package googlenetworkservicestcproute


type GoogleNetworkServicesTcpRouteRules struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#action GoogleNetworkServicesTcpRoute#action}
	Action *GoogleNetworkServicesTcpRouteRulesAction `field:"required" json:"action" yaml:"action"`
	// matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#matches GoogleNetworkServicesTcpRoute#matches}
	Matches interface{} `field:"optional" json:"matches" yaml:"matches"`
}

