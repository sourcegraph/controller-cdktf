package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionLogConfig struct {
	// Enabled represents whether logging is enabled or not for a connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_integration_connectors_connection#enabled GoogleIntegrationConnectorsConnection#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Log configuration level. Possible values: ["LOG_LEVEL_UNSPECIFIED", "ERROR", "INFO", "DEBUG"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_integration_connectors_connection#level GoogleIntegrationConnectorsConnection#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
}

