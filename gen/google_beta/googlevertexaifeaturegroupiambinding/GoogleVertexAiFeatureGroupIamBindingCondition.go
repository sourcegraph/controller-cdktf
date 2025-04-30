package googlevertexaifeaturegroupiambinding


type GoogleVertexAiFeatureGroupIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_group_iam_binding#expression GoogleVertexAiFeatureGroupIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_group_iam_binding#title GoogleVertexAiFeatureGroupIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_vertex_ai_feature_group_iam_binding#description GoogleVertexAiFeatureGroupIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

