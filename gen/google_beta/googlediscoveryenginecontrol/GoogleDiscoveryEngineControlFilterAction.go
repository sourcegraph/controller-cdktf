package googlediscoveryenginecontrol


type GoogleDiscoveryEngineControlFilterAction struct {
	// The data store to filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#data_store GoogleDiscoveryEngineControl#data_store}
	DataStore *string `field:"required" json:"dataStore" yaml:"dataStore"`
	// The filter to apply to the search results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#filter GoogleDiscoveryEngineControl#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
}

