package googlecestool


type GoogleCesToolDataStoreToolEngineSource struct {
	// Full resource name of the Engine. Format: 'projects/{project}/locations/{location}/collections/{collection}/engines/{engine}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#engine GoogleCesTool#engine}
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// data_store_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#data_store_sources GoogleCesTool#data_store_sources}
	DataStoreSources interface{} `field:"optional" json:"dataStoreSources" yaml:"dataStoreSources"`
	// A filter applied to the search across the Engine. Not relevant and not used if 'data_store_sources' is provided. See: https://cloud.google.com/generative-ai-app-builder/docs/filter-search-metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_ces_tool#filter GoogleCesTool#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

