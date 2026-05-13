package googlevertexaiendpointwithmodelgardendeployment


type GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigs struct {
	// Required. The full name of the Google Compute Engine network. Format: projects/{project}/global/networks/{network}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#network GoogleVertexAiEndpointWithModelGardenDeployment#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Required. Project id used to create forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#project_id GoogleVertexAiEndpointWithModelGardenDeployment#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

