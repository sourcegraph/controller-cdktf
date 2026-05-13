package vertexaireasoningengine


type VertexAiReasoningEngineSpecSourceCodeSpecImageSpec struct {
	// Build arguments to be used. They will be passed through --build-arg flags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine#build_args VertexAiReasoningEngine#build_args}
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
}

