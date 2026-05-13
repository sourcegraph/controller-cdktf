package googlediscoveryenginecontrol


type GoogleDiscoveryEngineControlPromoteAction struct {
	// The data store to promote.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#data_store GoogleDiscoveryEngineControl#data_store}
	DataStore *string `field:"required" json:"dataStore" yaml:"dataStore"`
	// search_link_promotion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_control#search_link_promotion GoogleDiscoveryEngineControl#search_link_promotion}
	SearchLinkPromotion *GoogleDiscoveryEngineControlPromoteActionSearchLinkPromotion `field:"required" json:"searchLinkPromotion" yaml:"searchLinkPromotion"`
}

