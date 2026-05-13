package discoveryenginecontrol


type DiscoveryEngineControlFilterAction struct {
	// The data store to filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#data_store DiscoveryEngineControl#data_store}
	DataStore *string `field:"required" json:"dataStore" yaml:"dataStore"`
	// The filter to apply to the search results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#filter DiscoveryEngineControl#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
}

