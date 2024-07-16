package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionDestinationConfigDestination struct {
	// For publicly routable host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#host GoogleIntegrationConnectorsConnection#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The port is the target port number that is accepted by the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#port GoogleIntegrationConnectorsConnection#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// PSC service attachments. Format: projects/* /regions/* /serviceAttachments/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#service_attachment GoogleIntegrationConnectorsConnection#service_attachment}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	ServiceAttachment *string `field:"optional" json:"serviceAttachment" yaml:"serviceAttachment"`
}

