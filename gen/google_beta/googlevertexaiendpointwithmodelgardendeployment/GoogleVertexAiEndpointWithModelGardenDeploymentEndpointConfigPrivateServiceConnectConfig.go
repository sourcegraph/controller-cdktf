package googlevertexaiendpointwithmodelgardendeployment


type GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfig struct {
	// Required. If true, expose the IndexEndpoint via private service connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#enable_private_service_connect GoogleVertexAiEndpointWithModelGardenDeployment#enable_private_service_connect}
	EnablePrivateServiceConnect interface{} `field:"required" json:"enablePrivateServiceConnect" yaml:"enablePrivateServiceConnect"`
	// A list of Projects from which the forwarding rule will target the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#project_allowlist GoogleVertexAiEndpointWithModelGardenDeployment#project_allowlist}
	ProjectAllowlist *[]*string `field:"optional" json:"projectAllowlist" yaml:"projectAllowlist"`
	// psc_automation_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_endpoint_with_model_garden_deployment#psc_automation_configs GoogleVertexAiEndpointWithModelGardenDeployment#psc_automation_configs}
	PscAutomationConfigs *GoogleVertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigs `field:"optional" json:"pscAutomationConfigs" yaml:"pscAutomationConfigs"`
}

