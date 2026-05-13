package googlevertexaiendpointwithmodelgardendeployment


type GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbeGrpc struct {
	// Port number of the gRPC service. Number must be in the range 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#port GoogleVertexAiEndpointWithModelGardenDeployment#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Service is the name of the service to place in the gRPC HealthCheckRequest. See https://github.com/grpc/grpc/blob/master/doc/health-checking.md.
	//
	// If this is not specified, the default behavior is defined by gRPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#service GoogleVertexAiEndpointWithModelGardenDeployment#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

