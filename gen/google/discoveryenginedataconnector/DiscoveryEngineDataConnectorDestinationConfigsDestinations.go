package discoveryenginedataconnector


type DiscoveryEngineDataConnectorDestinationConfigsDestinations struct {
	// The host of the destination, for example 'https://example.atlassian.net'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_data_connector#host DiscoveryEngineDataConnector#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Target port number accepted by the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/discovery_engine_data_connector#port DiscoveryEngineDataConnector#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

