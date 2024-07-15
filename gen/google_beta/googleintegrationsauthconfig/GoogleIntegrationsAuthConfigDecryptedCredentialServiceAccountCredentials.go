package googleintegrationsauthconfig


type GoogleIntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials struct {
	// A space-delimited list of requested scope permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#scope GoogleIntegrationsAuthConfig#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Name of the service account that has the permission to make the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#service_account GoogleIntegrationsAuthConfig#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

