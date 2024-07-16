package googlevertexaifeatureonlinestorefeatureview


type GoogleVertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceFeatureGroups struct {
	// Identifier of the feature group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#feature_group_id GoogleVertexAiFeatureOnlineStoreFeatureview#feature_group_id}
	FeatureGroupId *string `field:"required" json:"featureGroupId" yaml:"featureGroupId"`
	// Identifiers of features under the feature group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#feature_ids GoogleVertexAiFeatureOnlineStoreFeatureview#feature_ids}
	FeatureIds *[]*string `field:"required" json:"featureIds" yaml:"featureIds"`
}

