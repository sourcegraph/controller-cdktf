package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigs struct {
	// Required. The full name of the Google Compute Engine network. Format: projects/{project}/global/networks/{network}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#network VertexAiEndpointWithModelGardenDeployment#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Required. Project id used to create forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#project_id VertexAiEndpointWithModelGardenDeployment#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

