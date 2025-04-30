package googlediscoveryenginedatastore


type GoogleDiscoveryEngineDataStoreAdvancedSiteSearchConfig struct {
	// If set true, automatic refresh is disabled for the DataStore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_discovery_engine_data_store#disable_automatic_refresh GoogleDiscoveryEngineDataStore#disable_automatic_refresh}
	DisableAutomaticRefresh interface{} `field:"optional" json:"disableAutomaticRefresh" yaml:"disableAutomaticRefresh"`
	// If set true, initial indexing is disabled for the DataStore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_discovery_engine_data_store#disable_initial_index GoogleDiscoveryEngineDataStore#disable_initial_index}
	DisableInitialIndex interface{} `field:"optional" json:"disableInitialIndex" yaml:"disableInitialIndex"`
}

