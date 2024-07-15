package googlevertexaifeaturestoreiammember


type GoogleVertexAiFeaturestoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_iam_member#expression GoogleVertexAiFeaturestoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_iam_member#title GoogleVertexAiFeaturestoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore_iam_member#description GoogleVertexAiFeaturestoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

