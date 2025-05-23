package vertexaifeatureonlinestore


type VertexAiFeatureOnlineStoreBigtable struct {
	// auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/vertex_ai_feature_online_store#auto_scaling VertexAiFeatureOnlineStore#auto_scaling}
	AutoScaling *VertexAiFeatureOnlineStoreBigtableAutoScaling `field:"required" json:"autoScaling" yaml:"autoScaling"`
}

