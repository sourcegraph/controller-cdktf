package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfigDestination struct {
	// Host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#host GoogleIntegrationConnectorsConnection#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#port GoogleIntegrationConnectorsConnection#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Service Attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#service_attachment GoogleIntegrationConnectorsConnection#service_attachment}
	ServiceAttachment *string `field:"optional" json:"serviceAttachment" yaml:"serviceAttachment"`
}

