package googlevertexaifeatureonlinestoreiambinding


type GoogleVertexAiFeatureOnlineStoreIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_online_store_iam_binding#expression GoogleVertexAiFeatureOnlineStoreIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_online_store_iam_binding#title GoogleVertexAiFeatureOnlineStoreIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_online_store_iam_binding#description GoogleVertexAiFeatureOnlineStoreIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

