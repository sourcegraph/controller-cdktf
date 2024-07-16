package googleintegrationsauthconfig


type GoogleIntegrationsAuthConfigDecryptedCredentialUsernameAndPassword struct {
	// Password to be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#password GoogleIntegrationsAuthConfig#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Username to be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_integrations_auth_config#username GoogleIntegrationsAuthConfig#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

