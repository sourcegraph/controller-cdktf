package googlenetworkservicesgrpcroute


type GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort struct {
	// The HTTP status code used to abort the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#http_status GoogleNetworkServicesGrpcRoute#http_status}
	HttpStatus *float64 `field:"optional" json:"httpStatus" yaml:"httpStatus"`
	// The percentage of traffic which will be aborted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#percentage GoogleNetworkServicesGrpcRoute#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

