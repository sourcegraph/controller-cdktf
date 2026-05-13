package googlevertexaireasoningengine


type GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfig struct {
	// The default TTL duration of the memories in the Memory Bank.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#default_ttl GoogleVertexAiReasoningEngine#default_ttl}
	DefaultTtl *string `field:"optional" json:"defaultTtl" yaml:"defaultTtl"`
	// granular_ttl_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#granular_ttl_config GoogleVertexAiReasoningEngine#granular_ttl_config}
	GranularTtlConfig *GoogleVertexAiReasoningEngineContextSpecMemoryBankConfigTtlConfigGranularTtlConfig `field:"optional" json:"granularTtlConfig" yaml:"granularTtlConfig"`
	// The default TTL duration of the memory revisions in the Memory Bank.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine#memory_revision_default_ttl GoogleVertexAiReasoningEngine#memory_revision_default_ttl}
	MemoryRevisionDefaultTtl *string `field:"optional" json:"memoryRevisionDefaultTtl" yaml:"memoryRevisionDefaultTtl"`
}

