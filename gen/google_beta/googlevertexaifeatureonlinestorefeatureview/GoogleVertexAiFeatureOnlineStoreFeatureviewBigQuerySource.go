package googlevertexaifeatureonlinestorefeatureview


type GoogleVertexAiFeatureOnlineStoreFeatureviewBigQuerySource struct {
	// Columns to construct entityId / row keys. Start by supporting 1 only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#entity_id_columns GoogleVertexAiFeatureOnlineStoreFeatureview#entity_id_columns}
	EntityIdColumns *[]*string `field:"required" json:"entityIdColumns" yaml:"entityIdColumns"`
	// The BigQuery view URI that will be materialized on each sync trigger based on FeatureView.SyncConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#uri GoogleVertexAiFeatureOnlineStoreFeatureview#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

