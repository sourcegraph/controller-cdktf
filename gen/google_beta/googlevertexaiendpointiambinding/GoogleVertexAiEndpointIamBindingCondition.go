package googlevertexaiendpointiambinding


type GoogleVertexAiEndpointIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_endpoint_iam_binding#expression GoogleVertexAiEndpointIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_endpoint_iam_binding#title GoogleVertexAiEndpointIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_endpoint_iam_binding#description GoogleVertexAiEndpointIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

