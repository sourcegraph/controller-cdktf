package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfig struct {
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#destination GoogleIntegrationConnectorsConnection#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Key for the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#key GoogleIntegrationConnectorsConnection#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
}

