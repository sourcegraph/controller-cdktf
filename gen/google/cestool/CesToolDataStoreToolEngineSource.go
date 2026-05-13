package cestool


type CesToolDataStoreToolEngineSource struct {
	// Full resource name of the Engine. Format: 'projects/{project}/locations/{location}/collections/{collection}/engines/{engine}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#engine CesTool#engine}
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// data_store_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#data_store_sources CesTool#data_store_sources}
	DataStoreSources interface{} `field:"optional" json:"dataStoreSources" yaml:"dataStoreSources"`
	// A filter applied to the search across the Engine. Not relevant and not used if 'data_store_sources' is provided. See: https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#filter CesTool#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

