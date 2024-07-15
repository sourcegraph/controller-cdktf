package googlenetworkservicesgrpcroute


type GoogleNetworkServicesGrpcRouteRulesMatchesHeaders struct {
	// Required. The key of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#key GoogleNetworkServicesGrpcRoute#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Required. The value of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#value GoogleNetworkServicesGrpcRoute#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The type of match. Default value: "EXACT" Possible values: ["TYPE_UNSPECIFIED", "EXACT", "REGULAR_EXPRESSION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_network_services_grpc_route#type GoogleNetworkServicesGrpcRoute#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

