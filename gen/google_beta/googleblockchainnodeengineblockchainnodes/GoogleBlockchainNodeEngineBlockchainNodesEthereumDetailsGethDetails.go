package googleblockchainnodeengineblockchainnodes


type GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails struct {
	// Blockchain garbage collection modes. Only applicable when NodeType is FULL or ARCHIVE. Possible values: ["FULL", "ARCHIVE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#garbage_collection_mode GoogleBlockchainNodeEngineBlockchainNodes#garbage_collection_mode}
	GarbageCollectionMode *string `field:"optional" json:"garbageCollectionMode" yaml:"garbageCollectionMode"`
}

