package googlevertexaiendpointwithmodelgardendeployment


type GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecGrpcPorts struct {
	// The number of the port to expose on the pod's IP address.
	//
	// Must be a valid port number, between 1 and 65535 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#container_port GoogleVertexAiEndpointWithModelGardenDeployment#container_port}
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
}

