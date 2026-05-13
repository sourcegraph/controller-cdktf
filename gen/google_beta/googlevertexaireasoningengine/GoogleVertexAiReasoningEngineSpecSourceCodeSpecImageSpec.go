package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineSpecSourceCodeSpecImageSpec struct {
	// Build arguments to be used. They will be passed through --build-arg flags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#build_args GoogleVertexAiReasoningEngine#build_args}
	BuildArgs *map[string]*string `field:"optional" json:"buildArgs" yaml:"buildArgs"`
}

