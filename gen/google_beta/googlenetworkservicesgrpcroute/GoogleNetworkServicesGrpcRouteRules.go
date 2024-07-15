package googlenetworkservicesgrpcroute


type GoogleNetworkServicesGrpcRouteRules struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#action GoogleNetworkServicesGrpcRoute#action}
	Action *GoogleNetworkServicesGrpcRouteRulesAction `field:"optional" json:"action" yaml:"action"`
	// matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#matches GoogleNetworkServicesGrpcRoute#matches}
	Matches interface{} `field:"optional" json:"matches" yaml:"matches"`
}

