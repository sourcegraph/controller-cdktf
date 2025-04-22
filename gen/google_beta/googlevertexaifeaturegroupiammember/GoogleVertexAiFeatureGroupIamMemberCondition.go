package googlevertexaifeaturegroupiammember


type GoogleVertexAiFeatureGroupIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_group_iam_member#expression GoogleVertexAiFeatureGroupIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_group_iam_member#title GoogleVertexAiFeatureGroupIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_group_iam_member#description GoogleVertexAiFeatureGroupIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

