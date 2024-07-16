package googlenetworkservicesgrpcroute


type GoogleNetworkServicesGrpcRouteRulesMatches struct {
	// headers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#headers GoogleNetworkServicesGrpcRoute#headers}
	Headers interface{} `field:"optional" json:"headers" yaml:"headers"`
	// method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#method GoogleNetworkServicesGrpcRoute#method}
	Method *GoogleNetworkServicesGrpcRouteRulesMatchesMethod `field:"optional" json:"method" yaml:"method"`
}

