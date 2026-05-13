package discoveryenginedataconnector


type DiscoveryEngineDataConnectorDestinationConfigs struct {
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_data_connector#destinations DiscoveryEngineDataConnector#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// The key of the destination configuration, for example 'url'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_data_connector#key DiscoveryEngineDataConnector#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Additional parameters for this destination config in structured json format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_data_connector#params DiscoveryEngineDataConnector#params}
	Params *string `field:"optional" json:"params" yaml:"params"`
}

