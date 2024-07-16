package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionEventingConfig struct {
	// registration_destination_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#registration_destination_config GoogleIntegrationConnectorsConnection#registration_destination_config}
	RegistrationDestinationConfig *GoogleIntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfig `field:"required" json:"registrationDestinationConfig" yaml:"registrationDestinationConfig"`
	// additional_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#additional_variable GoogleIntegrationConnectorsConnection#additional_variable}
	AdditionalVariable interface{} `field:"optional" json:"additionalVariable" yaml:"additionalVariable"`
	// auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#auth_config GoogleIntegrationConnectorsConnection#auth_config}
	AuthConfig *GoogleIntegrationConnectorsConnectionEventingConfigAuthConfig `field:"optional" json:"authConfig" yaml:"authConfig"`
	// Enrichment Enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#enrichment_enabled GoogleIntegrationConnectorsConnection#enrichment_enabled}
	EnrichmentEnabled interface{} `field:"optional" json:"enrichmentEnabled" yaml:"enrichmentEnabled"`
}

