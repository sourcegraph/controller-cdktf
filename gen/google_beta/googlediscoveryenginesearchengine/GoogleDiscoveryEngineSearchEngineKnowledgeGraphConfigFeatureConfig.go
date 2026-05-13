package googlediscoveryenginesearchengine


type GoogleDiscoveryEngineSearchEngineKnowledgeGraphConfigFeatureConfig struct {
	// Whether to disable the private KG auto complete for the engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#disable_private_kg_auto_complete GoogleDiscoveryEngineSearchEngine#disable_private_kg_auto_complete}
	DisablePrivateKgAutoComplete interface{} `field:"optional" json:"disablePrivateKgAutoComplete" yaml:"disablePrivateKgAutoComplete"`
	// Whether to disable the private KG enrichment for the engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#disable_private_kg_enrichment GoogleDiscoveryEngineSearchEngine#disable_private_kg_enrichment}
	DisablePrivateKgEnrichment interface{} `field:"optional" json:"disablePrivateKgEnrichment" yaml:"disablePrivateKgEnrichment"`
	// Whether to disable the private KG for query UI chips.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#disable_private_kg_query_ui_chips GoogleDiscoveryEngineSearchEngine#disable_private_kg_query_ui_chips}
	DisablePrivateKgQueryUiChips interface{} `field:"optional" json:"disablePrivateKgQueryUiChips" yaml:"disablePrivateKgQueryUiChips"`
	// Whether to disable the private KG query understanding for the engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_search_engine#disable_private_kg_query_understanding GoogleDiscoveryEngineSearchEngine#disable_private_kg_query_understanding}
	DisablePrivateKgQueryUnderstanding interface{} `field:"optional" json:"disablePrivateKgQueryUnderstanding" yaml:"disablePrivateKgQueryUnderstanding"`
}

