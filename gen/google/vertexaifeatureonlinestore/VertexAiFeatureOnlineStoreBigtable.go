package vertexaifeatureonlinestore


type VertexAiFeatureOnlineStoreBigtable struct {
	// auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_feature_online_store#auto_scaling VertexAiFeatureOnlineStore#auto_scaling}
	AutoScaling *VertexAiFeatureOnlineStoreBigtableAutoScaling `field:"required" json:"autoScaling" yaml:"autoScaling"`
	// Optional. If true, enable direct access to the Bigtable instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_feature_online_store#enable_direct_bigtable_access VertexAiFeatureOnlineStore#enable_direct_bigtable_access}
	EnableDirectBigtableAccess interface{} `field:"optional" json:"enableDirectBigtableAccess" yaml:"enableDirectBigtableAccess"`
	// The zone where the Bigtable instance will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/vertex_ai_feature_online_store#zone VertexAiFeatureOnlineStore#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

