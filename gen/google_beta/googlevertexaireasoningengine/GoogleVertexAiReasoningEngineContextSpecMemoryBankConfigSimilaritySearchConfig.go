package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigSimilaritySearchConfig struct {
	// The model used to generate embeddings to lookup similar memories. Format: projects/{project}/locations/{location}/publishers/google/models/{model}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#embedding_model GoogleVertexAiReasoningEngine#embedding_model}
	EmbeddingModel *string `field:"required" json:"embeddingModel" yaml:"embeddingModel"`
}

