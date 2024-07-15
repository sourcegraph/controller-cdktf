package googlenetworkservicestlsroute


type GoogleNetworkServicesTlsRouteRules struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tls_route#action GoogleNetworkServicesTlsRoute#action}
	Action *GoogleNetworkServicesTlsRouteRulesAction `field:"required" json:"action" yaml:"action"`
	// matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tls_route#matches GoogleNetworkServicesTlsRoute#matches}
	Matches interface{} `field:"required" json:"matches" yaml:"matches"`
}

