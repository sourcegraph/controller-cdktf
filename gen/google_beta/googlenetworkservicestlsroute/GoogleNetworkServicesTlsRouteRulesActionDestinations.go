package googlenetworkservicestlsroute


type GoogleNetworkServicesTlsRouteRulesActionDestinations struct {
	// The URL of a BackendService to route traffic to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tls_route#service_name GoogleNetworkServicesTlsRoute#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_tls_route#weight GoogleNetworkServicesTlsRoute#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

