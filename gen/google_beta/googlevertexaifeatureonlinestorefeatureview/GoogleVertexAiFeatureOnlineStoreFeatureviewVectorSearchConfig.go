package googlevertexaifeatureonlinestorefeatureview


type GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfig struct {
	// Column of embedding. This column contains the source data to create index for vector search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#embedding_column GoogleVertexAiFeatureOnlineStoreFeatureview#embedding_column}
	EmbeddingColumn *string `field:"required" json:"embeddingColumn" yaml:"embeddingColumn"`
	// brute_force_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#brute_force_config GoogleVertexAiFeatureOnlineStoreFeatureview#brute_force_config}
	BruteForceConfig *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigBruteForceConfig `field:"optional" json:"bruteForceConfig" yaml:"bruteForceConfig"`
	// Column of crowding.
	//
	// This column contains crowding attribute which is a constraint on a neighbor list produced by nearest neighbor search requiring that no more than some value k' of the k neighbors returned have the same value of crowdingAttribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#crowding_column GoogleVertexAiFeatureOnlineStoreFeatureview#crowding_column}
	CrowdingColumn *string `field:"optional" json:"crowdingColumn" yaml:"crowdingColumn"`
	// The distance measure used in nearest neighbor search.
	//
	// For details on allowed values, see the [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.featureOnlineStores.featureViews#DistanceMeasureType). Possible values: ["SQUARED_L2_DISTANCE", "COSINE_DISTANCE", "DOT_PRODUCT_DISTANCE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#distance_measure_type GoogleVertexAiFeatureOnlineStoreFeatureview#distance_measure_type}
	DistanceMeasureType *string `field:"optional" json:"distanceMeasureType" yaml:"distanceMeasureType"`
	// The number of dimensions of the input embedding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#embedding_dimension GoogleVertexAiFeatureOnlineStoreFeatureview#embedding_dimension}
	EmbeddingDimension *float64 `field:"optional" json:"embeddingDimension" yaml:"embeddingDimension"`
	// Columns of features that are used to filter vector search results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#filter_columns GoogleVertexAiFeatureOnlineStoreFeatureview#filter_columns}
	FilterColumns *[]*string `field:"optional" json:"filterColumns" yaml:"filterColumns"`
	// tree_ah_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_feature_online_store_featureview#tree_ah_config GoogleVertexAiFeatureOnlineStoreFeatureview#tree_ah_config}
	TreeAhConfig *GoogleVertexAiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfig `field:"optional" json:"treeAhConfig" yaml:"treeAhConfig"`
}

