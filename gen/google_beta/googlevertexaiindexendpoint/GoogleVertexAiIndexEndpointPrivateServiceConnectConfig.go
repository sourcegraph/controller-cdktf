package googlevertexaiindexendpoint


type GoogleVertexAiIndexEndpointPrivateServiceConnectConfig struct {
	// If set to true, the IndexEndpoint is created without private service access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_index_endpoint#enable_private_service_connect GoogleVertexAiIndexEndpoint#enable_private_service_connect}
	EnablePrivateServiceConnect interface{} `field:"required" json:"enablePrivateServiceConnect" yaml:"enablePrivateServiceConnect"`
	// A list of Projects from which the forwarding rule will target the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_index_endpoint#project_allowlist GoogleVertexAiIndexEndpoint#project_allowlist}
	ProjectAllowlist *[]*string `field:"optional" json:"projectAllowlist" yaml:"projectAllowlist"`
}

