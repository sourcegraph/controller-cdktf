package googlenetworkservicestcproute


type GoogleNetworkServicesTcpRouteRulesActionDestinations struct {
	// The URL of a BackendService to route traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#service_name GoogleNetworkServicesTcpRoute#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.
	//
	// This is computed as: weight/Sum(weights in this destination list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports.
	// If only one serviceName is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend.
	// If weights are specified for any one service name, they need to be specified for all of them.
	// If weights are unspecified for all services, then, traffic is distributed in equal proportions to all of them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tcp_route#weight GoogleNetworkServicesTcpRoute#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

