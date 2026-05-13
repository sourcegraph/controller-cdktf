package googleblockchainnodeengineblockchainnodes


type GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig struct {
	// An Ethereum address which the beacon client will send fee rewards to if no recipient is configured in the validator client.
	//
	// See https://lighthouse-book.sigmaprime.io/suggested-fee-recipient.html or https://docs.prylabs.network/docs/execution-node/fee-recipient for examples of how this is used. Note that while this is often described as "suggested", as we run the execution node we can trust the execution node, and therefore this is considered enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#beacon_fee_recipient GoogleBlockchainNodeEngineBlockchainNodes#beacon_fee_recipient}
	BeaconFeeRecipient *string `field:"optional" json:"beaconFeeRecipient" yaml:"beaconFeeRecipient"`
	// URLs for MEV-relay services to use for block building.
	//
	// When set, a managed MEV-boost service is configured on the beacon client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#mev_relay_urls GoogleBlockchainNodeEngineBlockchainNodes#mev_relay_urls}
	MevRelayUrls *[]*string `field:"optional" json:"mevRelayUrls" yaml:"mevRelayUrls"`
}

