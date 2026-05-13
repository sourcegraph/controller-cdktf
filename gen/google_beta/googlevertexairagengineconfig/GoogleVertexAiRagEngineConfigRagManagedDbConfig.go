package googlevertexairagengineconfig


type GoogleVertexAiRagEngineConfigRagManagedDbConfig struct {
	// basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_rag_engine_config#basic GoogleVertexAiRagEngineConfig#basic}
	Basic *GoogleVertexAiRagEngineConfigRagManagedDbConfigBasic `field:"optional" json:"basic" yaml:"basic"`
	// scaled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_rag_engine_config#scaled GoogleVertexAiRagEngineConfig#scaled}
	Scaled *GoogleVertexAiRagEngineConfigRagManagedDbConfigScaled `field:"optional" json:"scaled" yaml:"scaled"`
	// unprovisioned block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_rag_engine_config#unprovisioned GoogleVertexAiRagEngineConfig#unprovisioned}
	Unprovisioned *GoogleVertexAiRagEngineConfigRagManagedDbConfigUnprovisioned `field:"optional" json:"unprovisioned" yaml:"unprovisioned"`
}

