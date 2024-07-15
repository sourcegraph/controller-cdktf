package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials struct {
	// Secret version of Password for Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#client_id GoogleIntegrationConnectorsConnection#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// client_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#client_secret GoogleIntegrationConnectorsConnection#client_secret}
	ClientSecret *GoogleIntegrationConnectorsConnectionAuthConfigOauth2ClientCredentialsClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

