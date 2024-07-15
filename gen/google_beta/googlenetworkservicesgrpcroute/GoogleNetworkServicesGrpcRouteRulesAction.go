package googlenetworkservicesgrpcroute


type GoogleNetworkServicesGrpcRouteRulesAction struct {
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#destinations GoogleNetworkServicesGrpcRoute#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#fault_injection_policy GoogleNetworkServicesGrpcRoute#fault_injection_policy}
	FaultInjectionPolicy *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#retry_policy GoogleNetworkServicesGrpcRoute#retry_policy}
	RetryPolicy *GoogleNetworkServicesGrpcRouteRulesActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Specifies the timeout for selected route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#timeout GoogleNetworkServicesGrpcRoute#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

