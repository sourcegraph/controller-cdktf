package googleintegrationsauthconfig


type GoogleIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode struct {
	// The auth url endpoint to send the auth code request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#auth_endpoint GoogleIntegrationsAuthConfig#auth_endpoint}
	AuthEndpoint *string `field:"optional" json:"authEndpoint" yaml:"authEndpoint"`
	// The client's id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#client_id GoogleIntegrationsAuthConfig#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client's secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#client_secret GoogleIntegrationsAuthConfig#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// A space-delimited list of requested scope permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#scope GoogleIntegrationsAuthConfig#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The token url endpoint to send the token request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#token_endpoint GoogleIntegrationsAuthConfig#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
}

