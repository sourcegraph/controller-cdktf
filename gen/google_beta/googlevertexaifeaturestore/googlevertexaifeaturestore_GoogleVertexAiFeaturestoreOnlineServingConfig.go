package googlevertexaifeaturestore


type GoogleVertexAiFeaturestoreOnlineServingConfig struct {
	// The number of nodes for each cluster.
	//
	// The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore#fixed_node_count GoogleVertexAiFeaturestore#fixed_node_count}
	FixedNodeCount *float64 `field:"required" json:"fixedNodeCount" yaml:"fixedNodeCount"`
}

