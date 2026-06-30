package networkservicesgrpcroute

type NetworkServicesGrpcRouteRules struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_services_grpc_route#action NetworkServicesGrpcRoute#action}
	Action *NetworkServicesGrpcRouteRulesAction `field:"optional" json:"action" yaml:"action"`
	// matches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.32.0/docs/resources/network_services_grpc_route#matches NetworkServicesGrpcRoute#matches}
	Matches any `field:"optional" json:"matches" yaml:"matches"`
}
