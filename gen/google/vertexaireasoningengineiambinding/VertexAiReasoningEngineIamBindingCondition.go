package vertexaireasoningengineiambinding


type VertexAiReasoningEngineIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine_iam_binding#expression VertexAiReasoningEngineIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine_iam_binding#title VertexAiReasoningEngineIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_reasoning_engine_iam_binding#description VertexAiReasoningEngineIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

