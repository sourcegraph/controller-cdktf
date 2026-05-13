package googlevertexaiendpointwithmodelgardendeployment


type GoogleVertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbeHttpGetHttpHeaders struct {
	// The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#name GoogleVertexAiEndpointWithModelGardenDeployment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The header field value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#value GoogleVertexAiEndpointWithModelGardenDeployment#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

