package googleintegrationsauthconfig


type GoogleIntegrationsAuthConfigDecryptedCredentialAuthToken struct {
	// The token for the auth type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#token GoogleIntegrationsAuthConfig#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Authentication type, e.g. "Basic", "Bearer", etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#type GoogleIntegrationsAuthConfig#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

