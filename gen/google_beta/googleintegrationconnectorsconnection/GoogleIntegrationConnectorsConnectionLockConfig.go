package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionLockConfig struct {
	// Indicates whether or not the connection is locked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#locked GoogleIntegrationConnectorsConnection#locked}
	Locked interface{} `field:"required" json:"locked" yaml:"locked"`
	// Describes why a connection is locked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#reason GoogleIntegrationConnectorsConnection#reason}
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

