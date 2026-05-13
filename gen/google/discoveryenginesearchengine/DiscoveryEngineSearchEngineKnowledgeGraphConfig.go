package discoveryenginesearchengine


type DiscoveryEngineSearchEngineKnowledgeGraphConfig struct {
	// Specify entity types to support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_search_engine#cloud_knowledge_graph_types DiscoveryEngineSearchEngine#cloud_knowledge_graph_types}
	CloudKnowledgeGraphTypes *[]*string `field:"optional" json:"cloudKnowledgeGraphTypes" yaml:"cloudKnowledgeGraphTypes"`
	// Whether to enable the Cloud Knowledge Graph for the engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_search_engine#enable_cloud_knowledge_graph DiscoveryEngineSearchEngine#enable_cloud_knowledge_graph}
	EnableCloudKnowledgeGraph interface{} `field:"optional" json:"enableCloudKnowledgeGraph" yaml:"enableCloudKnowledgeGraph"`
	// Whether to enable the Private Knowledge Graph for the engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_search_engine#enable_private_knowledge_graph DiscoveryEngineSearchEngine#enable_private_knowledge_graph}
	EnablePrivateKnowledgeGraph interface{} `field:"optional" json:"enablePrivateKnowledgeGraph" yaml:"enablePrivateKnowledgeGraph"`
	// feature_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_search_engine#feature_config DiscoveryEngineSearchEngine#feature_config}
	FeatureConfig *DiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig `field:"optional" json:"featureConfig" yaml:"featureConfig"`
}

