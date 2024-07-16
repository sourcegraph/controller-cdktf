package googleintegrationsauthconfig


type GoogleIntegrationsAuthConfigDecryptedCredential struct {
	// Credential type associated with auth configs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#credential_type GoogleIntegrationsAuthConfig#credential_type}
	CredentialType *string `field:"required" json:"credentialType" yaml:"credentialType"`
	// auth_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#auth_token GoogleIntegrationsAuthConfig#auth_token}
	AuthToken *GoogleIntegrationsAuthConfigDecryptedCredentialAuthToken `field:"optional" json:"authToken" yaml:"authToken"`
	// jwt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#jwt GoogleIntegrationsAuthConfig#jwt}
	Jwt *GoogleIntegrationsAuthConfigDecryptedCredentialJwt `field:"optional" json:"jwt" yaml:"jwt"`
	// oauth2_authorization_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#oauth2_authorization_code GoogleIntegrationsAuthConfig#oauth2_authorization_code}
	Oauth2AuthorizationCode *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode `field:"optional" json:"oauth2AuthorizationCode" yaml:"oauth2AuthorizationCode"`
	// oauth2_client_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#oauth2_client_credentials GoogleIntegrationsAuthConfig#oauth2_client_credentials}
	Oauth2ClientCredentials *GoogleIntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials `field:"optional" json:"oauth2ClientCredentials" yaml:"oauth2ClientCredentials"`
	// oidc_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#oidc_token GoogleIntegrationsAuthConfig#oidc_token}
	OidcToken *GoogleIntegrationsAuthConfigDecryptedCredentialOidcToken `field:"optional" json:"oidcToken" yaml:"oidcToken"`
	// service_account_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#service_account_credentials GoogleIntegrationsAuthConfig#service_account_credentials}
	ServiceAccountCredentials *GoogleIntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials `field:"optional" json:"serviceAccountCredentials" yaml:"serviceAccountCredentials"`
	// username_and_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#username_and_password GoogleIntegrationsAuthConfig#username_and_password}
	UsernameAndPassword *GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPassword `field:"optional" json:"usernameAndPassword" yaml:"usernameAndPassword"`
}

