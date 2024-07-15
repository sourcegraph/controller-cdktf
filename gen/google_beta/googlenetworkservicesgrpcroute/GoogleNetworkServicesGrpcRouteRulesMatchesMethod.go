package googlenetworkservicesgrpcroute


type GoogleNetworkServicesGrpcRouteRulesMatchesMethod struct {
	// Required. Name of the method to match against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#grpc_method GoogleNetworkServicesGrpcRoute#grpc_method}
	GrpcMethod *string `field:"required" json:"grpcMethod" yaml:"grpcMethod"`
	// Required. Name of the service to match against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#grpc_service GoogleNetworkServicesGrpcRoute#grpc_service}
	GrpcService *string `field:"required" json:"grpcService" yaml:"grpcService"`
	// Specifies that matches are case sensitive. The default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#case_sensitive GoogleNetworkServicesGrpcRoute#case_sensitive}
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
}

