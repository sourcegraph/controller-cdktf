package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfig struct {
	// If true, no memory revisions will be created for any requests to the Memory Bank.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#disable_memory_revisions GoogleVertexAiReasoningEngine#disable_memory_revisions}
	DisableMemoryRevisions interface{} `field:"optional" json:"disableMemoryRevisions" yaml:"disableMemoryRevisions"`
	// generation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#generation_config GoogleVertexAiReasoningEngine#generation_config}
	GenerationConfig *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigGenerationConfig `field:"optional" json:"generationConfig" yaml:"generationConfig"`
	// similarity_search_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#similarity_search_config GoogleVertexAiReasoningEngine#similarity_search_config}
	SimilaritySearchConfig *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig `field:"optional" json:"similaritySearchConfig" yaml:"similaritySearchConfig"`
	// ttl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#ttl_config GoogleVertexAiReasoningEngine#ttl_config}
	TtlConfig *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig `field:"optional" json:"ttlConfig" yaml:"ttlConfig"`
}

