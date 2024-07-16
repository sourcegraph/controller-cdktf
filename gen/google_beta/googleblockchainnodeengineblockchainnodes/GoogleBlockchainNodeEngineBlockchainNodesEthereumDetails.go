package googleblockchainnodeengineblockchainnodes


type GoogleBlockchainNodeEngineBlockchainNodesEthereumDetails struct {
	// Enables JSON-RPC access to functions in the admin namespace. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#api_enable_admin GoogleBlockchainNodeEngineBlockchainNodes#api_enable_admin}
	ApiEnableAdmin interface{} `field:"optional" json:"apiEnableAdmin" yaml:"apiEnableAdmin"`
	// Enables JSON-RPC access to functions in the debug namespace. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#api_enable_debug GoogleBlockchainNodeEngineBlockchainNodes#api_enable_debug}
	ApiEnableDebug interface{} `field:"optional" json:"apiEnableDebug" yaml:"apiEnableDebug"`
	// The consensus client Possible values: ["CONSENSUS_CLIENT_UNSPECIFIED", "LIGHTHOUSE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#consensus_client GoogleBlockchainNodeEngineBlockchainNodes#consensus_client}
	ConsensusClient *string `field:"optional" json:"consensusClient" yaml:"consensusClient"`
	// The execution client Possible values: ["EXECUTION_CLIENT_UNSPECIFIED", "GETH", "ERIGON"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#execution_client GoogleBlockchainNodeEngineBlockchainNodes#execution_client}
	ExecutionClient *string `field:"optional" json:"executionClient" yaml:"executionClient"`
	// geth_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#geth_details GoogleBlockchainNodeEngineBlockchainNodes#geth_details}
	FetchhDetails *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails `field:"optional" json:"fetchhDetails" yaml:"fetchhDetails"`
	// The Ethereum environment being accessed. Possible values: ["MAINNET", "TESTNET_GOERLI_PRATER", "TESTNET_SEPOLIA"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#network GoogleBlockchainNodeEngineBlockchainNodes#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The type of Ethereum node. Possible values: ["LIGHT", "FULL", "ARCHIVE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#node_type GoogleBlockchainNodeEngineBlockchainNodes#node_type}
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// validator_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_blockchain_node_engine_blockchain_nodes#validator_config GoogleBlockchainNodeEngineBlockchainNodes#validator_config}
	ValidatorConfig *GoogleBlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig `field:"optional" json:"validatorConfig" yaml:"validatorConfig"`
}

