package googlevertexaifeatureonlinestoreiammember


type GoogleVertexAiFeatureOnlineStoreIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_online_store_iam_member#expression GoogleVertexAiFeatureOnlineStoreIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_online_store_iam_member#title GoogleVertexAiFeatureOnlineStoreIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_online_store_iam_member#description GoogleVertexAiFeatureOnlineStoreIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

