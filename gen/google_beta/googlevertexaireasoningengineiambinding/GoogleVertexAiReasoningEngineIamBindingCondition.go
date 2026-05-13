package googlevertexaireasoningengineiambinding


type GoogleVertexAiReasoningEngineIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine_iam_binding#expression GoogleVertexAiReasoningEngineIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine_iam_binding#title GoogleVertexAiReasoningEngineIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine_iam_binding#description GoogleVertexAiReasoningEngineIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

