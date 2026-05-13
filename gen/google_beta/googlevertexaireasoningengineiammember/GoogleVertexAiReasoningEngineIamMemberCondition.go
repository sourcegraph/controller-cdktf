package googlevertexaireasoningengineiammember


type GoogleVertexAiReasoningEngineIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine_iam_member#expression GoogleVertexAiReasoningEngineIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine_iam_member#title GoogleVertexAiReasoningEngineIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_reasoning_engine_iam_member#description GoogleVertexAiReasoningEngineIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

