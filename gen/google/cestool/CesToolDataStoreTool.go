package cestool


type CesToolDataStoreTool struct {
	// The data store tool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#name CesTool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// boost_specs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#boost_specs CesTool#boost_specs}
	BoostSpecs interface{} `field:"optional" json:"boostSpecs" yaml:"boostSpecs"`
	// The tool description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#description CesTool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// engine_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#engine_source CesTool#engine_source}
	EngineSource *CesToolDataStoreToolEngineSource `field:"optional" json:"engineSource" yaml:"engineSource"`
	// Number of search results to return per query. The default value is 10. The maximum allowed value is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#max_results CesTool#max_results}
	MaxResults *float64 `field:"optional" json:"maxResults" yaml:"maxResults"`
	// modality_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/ces_tool#modality_configs CesTool#modality_configs}
	ModalityConfigs interface{} `field:"optional" json:"modalityConfigs" yaml:"modalityConfigs"`
}

