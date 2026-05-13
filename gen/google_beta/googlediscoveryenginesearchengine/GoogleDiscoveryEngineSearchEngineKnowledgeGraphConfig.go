package googlediscoveryenginesearchengine


type GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfig struct {
	// Specify entity types to support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#cloud_knowledge_graph_types GoogleDiscoveryEngineSearchEngine#cloud_knowledge_graph_types}
	CloudKnowledgeGraphTypes *[]*string `field:"optional" json:"cloudKnowledgeGraphTypes" yaml:"cloudKnowledgeGraphTypes"`
	// Whether to enable the Cloud Knowledge Graph for the engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#enable_cloud_knowledge_graph GoogleDiscoveryEngineSearchEngine#enable_cloud_knowledge_graph}
	EnableCloudKnowledgeGraph interface{} `field:"optional" json:"enableCloudKnowledgeGraph" yaml:"enableCloudKnowledgeGraph"`
	// Whether to enable the Private Knowledge Graph for the engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#enable_private_knowledge_graph GoogleDiscoveryEngineSearchEngine#enable_private_knowledge_graph}
	EnablePrivateKnowledgeGraph interface{} `field:"optional" json:"enablePrivateKnowledgeGraph" yaml:"enablePrivateKnowledgeGraph"`
	// feature_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#feature_config GoogleDiscoveryEngineSearchEngine#feature_config}
	FeatureConfig *GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig `field:"optional" json:"featureConfig" yaml:"featureConfig"`
}

