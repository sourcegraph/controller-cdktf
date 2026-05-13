package vertexaireasoningengine


type VertexAiReasoningEngineSpecDeploymentSpecSecretEnv struct {
	// The name of the environment variable. Must be a valid C identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#name VertexAiReasoningEngine#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// secret_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#secret_ref VertexAiReasoningEngine#secret_ref}
	SecretRef *VertexAiReasoningEngineSpecDeploymentSpecSecretEnvSecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

