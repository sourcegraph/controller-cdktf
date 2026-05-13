package vertexaireasoningengine


type VertexAiReasoningEngineSpecDeploymentSpecEnv struct {
	// The name of the environment variable. Must be a valid C identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#name VertexAiReasoningEngine#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Variables that reference a $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables.
	//
	// If a variable cannot
	// be resolved, the reference in the input string will be
	// unchanged. The $(VAR_NAME) syntax can be escaped with a
	// double $$, ie: $$(VAR_NAME). Escaped references will never
	// be expanded, regardless of whether the variable exists
	// or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#value VertexAiReasoningEngine#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

