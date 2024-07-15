package googlevertexaifeaturestore


type GoogleVertexAiFeaturestoreOnlineServingConfig struct {
	// The number of nodes for each cluster.
	//
	// The number of nodes will not scale automatically but can be scaled manually by providing different values when updating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore#fixed_node_count GoogleVertexAiFeaturestore#fixed_node_count}
	FixedNodeCount *float64 `field:"optional" json:"fixedNodeCount" yaml:"fixedNodeCount"`
	// scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_vertex_ai_featurestore#scaling GoogleVertexAiFeaturestore#scaling}
	Scaling *GoogleVertexAiFeaturestoreOnlineServingConfigScaling `field:"optional" json:"scaling" yaml:"scaling"`
}

