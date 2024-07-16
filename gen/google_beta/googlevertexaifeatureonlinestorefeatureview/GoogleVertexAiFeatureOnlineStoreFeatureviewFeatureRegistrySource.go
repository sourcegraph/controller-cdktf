package googlevertexaifeatureonlinestorefeatureview


type GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySource struct {
	// feature_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#feature_groups GoogleVertexAiFeatureOnlineStoreFeatureview#feature_groups}
	FeatureGroups interface{} `field:"required" json:"featureGroups" yaml:"featureGroups"`
	// The project number of the parent project of the feature Groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#project_number GoogleVertexAiFeatureOnlineStoreFeatureview#project_number}
	ProjectNumber *string `field:"optional" json:"projectNumber" yaml:"projectNumber"`
}

