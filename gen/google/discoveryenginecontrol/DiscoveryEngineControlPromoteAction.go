package discoveryenginecontrol


type DiscoveryEngineControlPromoteAction struct {
	// The data store to promote.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#data_store DiscoveryEngineControl#data_store}
	DataStore *string `field:"required" json:"dataStore" yaml:"dataStore"`
	// search_link_promotion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_control#search_link_promotion DiscoveryEngineControl#search_link_promotion}
	SearchLinkPromotion *DiscoveryEngineControlPromoteActionSearchLinkPromotion `field:"required" json:"searchLinkPromotion" yaml:"searchLinkPromotion"`
}

