package vertexaiendpoint


type VertexAiEndpointPrivateServiceConnectConfig struct {
	// Required. If true, expose the IndexEndpoint via private service connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_endpoint#enable_private_service_connect VertexAiEndpoint#enable_private_service_connect}
	EnablePrivateServiceConnect interface{} `field:"required" json:"enablePrivateServiceConnect" yaml:"enablePrivateServiceConnect"`
	// A list of Projects from which the forwarding rule will target the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_endpoint#project_allowlist VertexAiEndpoint#project_allowlist}
	ProjectAllowlist *[]*string `field:"optional" json:"projectAllowlist" yaml:"projectAllowlist"`
	// psc_automation_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_endpoint#psc_automation_configs VertexAiEndpoint#psc_automation_configs}
	PscAutomationConfigs interface{} `field:"optional" json:"pscAutomationConfigs" yaml:"pscAutomationConfigs"`
}

