package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfigGranularTtlConfig struct {
	// The TTL duration for memories uploaded via CreateMemory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#create_ttl GoogleVertexAiReasoningEngine#create_ttl}
	CreateTtl *string `field:"optional" json:"createTtl" yaml:"createTtl"`
	// The TTL duration for memories newly generated via GenerateMemories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#generate_created_ttl GoogleVertexAiReasoningEngine#generate_created_ttl}
	GenerateCreatedTtl *string `field:"optional" json:"generateCreatedTtl" yaml:"generateCreatedTtl"`
	// The TTL duration for memories updated via GenerateMemories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#generate_updated_ttl GoogleVertexAiReasoningEngine#generate_updated_ttl}
	GenerateUpdatedTtl *string `field:"optional" json:"generateUpdatedTtl" yaml:"generateUpdatedTtl"`
}

