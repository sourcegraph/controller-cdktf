package googlevertexaifeatureonlinestorefeatureviewiambinding


type GoogleVertexAiFeatureOnlineStoreFeatureviewIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_binding#expression GoogleVertexAiFeatureOnlineStoreFeatureviewIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_binding#title GoogleVertexAiFeatureOnlineStoreFeatureviewIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_binding#description GoogleVertexAiFeatureOnlineStoreFeatureviewIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

