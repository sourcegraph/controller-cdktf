package googlediscoveryenginedataconnector


type GoogleDiscoveryEngineDataConnectorDestinationConfigsDestinations struct {
	// The host of the destination, for example 'https://example.atlassian.net'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#host GoogleDiscoveryEngineDataConnector#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Target port number accepted by the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_discovery_engine_data_connector#port GoogleDiscoveryEngineDataConnector#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

