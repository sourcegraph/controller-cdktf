package discoveryenginedataconnector


type DiscoveryEngineDataConnectorActionConfig struct {
	// Params needed to configure the actions in the format of String-to-String (Key, Value) pairs.
	//
	// Contains connection
	// credentials and configuration for the action connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_data_connector#action_params DiscoveryEngineDataConnector#action_params}
	ActionParams *map[string]*string `field:"optional" json:"actionParams" yaml:"actionParams"`
	// Whether to create a BAP (Business Application Platform) connection for this action connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_data_connector#create_bap_connection DiscoveryEngineDataConnector#create_bap_connection}
	CreateBapConnection interface{} `field:"optional" json:"createBapConnection" yaml:"createBapConnection"`
}

