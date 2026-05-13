package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecDeploymentSpecSecretEnv struct {
	// The name of the environment variable. Must be a valid C identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#name GoogleVertexAiReasoningEngine#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#secret_ref GoogleVertexAiReasoningEngine#secret_ref}
	SecretRef *GoogleVertexAiReasoningEngineSpecDeploymentSpecSecretEnvSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

