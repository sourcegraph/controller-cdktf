package googleblockchainnodeengineblockchainnodes


type GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig struct {
	// URLs for MEV-relay services to use for block building.
	//
	// When set, a managed MEV-boost service is configured on the beacon client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#mev_relay_urls GoogleBlockchainNodeEngineBlockchainNodes#mev_relay_urls}
	MevRelayUrls *[]*string `field:"optional" json:"mevRelayUrls" yaml:"mevRelayUrls"`
}

