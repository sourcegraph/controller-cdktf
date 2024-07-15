package googleintegrationsauthconfig


type GoogleIntegrationsAuthConfigDecryptedCredentialOidcToken struct {
	// Audience to be used when generating OIDC token.
	//
	// The audience claim identifies the recipients that the JWT is intended for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#audience GoogleIntegrationsAuthConfig#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// The service account email to be used as the identity for the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#service_account_email GoogleIntegrationsAuthConfig#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

