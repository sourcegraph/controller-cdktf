package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecDeploymentSpecSecretEnvSecretRef struct {
	// The name of the secret in Cloud Secret Manager. Format: {secret_name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#secret GoogleVertexAiReasoningEngine#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// The Cloud Secret Manager secret version.
	//
	// Can be 'latest'
	// for the latest version, an integer for a specific
	// version, or a version alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#version GoogleVertexAiReasoningEngine#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

