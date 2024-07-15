package googlenetworkservicesgrpcroute


type GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay struct {
	// Specify a fixed delay before forwarding the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#fixed_delay GoogleNetworkServicesGrpcRoute#fixed_delay}
	FixedDelay *string `field:"optional" json:"fixedDelay" yaml:"fixedDelay"`
	// The percentage of traffic on which delay will be injected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#percentage GoogleNetworkServicesGrpcRoute#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

