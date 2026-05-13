package googlevertexaifeatureonlinestore


type GoogleVertexAiFeatureOnlineStoreBigtable struct {
	// auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_feature_online_store#auto_scaling GoogleVertexAiFeatureOnlineStore#auto_scaling}
	AutoScaling *GoogleVertexAiFeatureOnlineStoreBigtableAutoScaling `field:"required" json:"autoScaling" yaml:"autoScaling"`
	// Optional. If true, enable direct access to the Bigtable instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_feature_online_store#enable_direct_bigtable_access GoogleVertexAiFeatureOnlineStore#enable_direct_bigtable_access}
	EnableDirectBigtableAccess interface{} `field:"optional" json:"enableDirectBigtableAccess" yaml:"enableDirectBigtableAccess"`
	// The zone where the Bigtable instance will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_vertex_ai_feature_online_store#zone GoogleVertexAiFeatureOnlineStore#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

