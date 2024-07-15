package googleintegrationconnectorsconnection


type GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow struct {
	// Auth URL for Authorization Code Flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#auth_uri GoogleIntegrationConnectorsConnection#auth_uri}
	AuthUri *string `field:"optional" json:"authUri" yaml:"authUri"`
	// Client ID for user-provided OAuth app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#client_id GoogleIntegrationConnectorsConnection#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// client_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#client_secret GoogleIntegrationConnectorsConnection#client_secret}
	ClientSecret *GoogleIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Whether to enable PKCE when the user performs the auth code flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#enable_pkce GoogleIntegrationConnectorsConnection#enable_pkce}
	EnablePkce interface{} `field:"optional" json:"enablePkce" yaml:"enablePkce"`
	// Scopes the connection will request when the user performs the auth code flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integration_connectors_connection#scopes GoogleIntegrationConnectorsConnection#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

