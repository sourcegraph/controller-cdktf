package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfig struct {
	// Required. If true, expose the IndexEndpoint via private service connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#enable_private_service_connect VertexAiEndpointWithModelGardenDeployment#enable_private_service_connect}
	EnablePrivateServiceConnect interface{} `field:"required" json:"enablePrivateServiceConnect" yaml:"enablePrivateServiceConnect"`
	// A list of Projects from which the forwarding rule will target the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#project_allowlist VertexAiEndpointWithModelGardenDeployment#project_allowlist}
	ProjectAllowlist *[]*string `field:"optional" json:"projectAllowlist" yaml:"projectAllowlist"`
	// psc_automation_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#psc_automation_configs VertexAiEndpointWithModelGardenDeployment#psc_automation_configs}
	PscAutomationConfigs *VertexAiEndpointWithModelGardenDeploymentEndpointConfigPrivateServiceConnectConfigPscAutomationConfigs `field:"optional" json:"pscAutomationConfigs" yaml:"pscAutomationConfigs"`
}

