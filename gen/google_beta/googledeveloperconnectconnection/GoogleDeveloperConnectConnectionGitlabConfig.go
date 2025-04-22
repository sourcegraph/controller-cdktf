package googledeveloperconnectconnection


type GoogleDeveloperConnectConnectionGitlabConfig struct {
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#authorizer_credential GoogleDeveloperConnectConnection#authorizer_credential}
	AuthorizerCredential *GoogleDeveloperConnectConnectionGitlabConfigAuthorizerCredential `field:"required" json:"authorizerCredential" yaml:"authorizerCredential"`
	// read_authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#read_authorizer_credential GoogleDeveloperConnectConnection#read_authorizer_credential}
	ReadAuthorizerCredential *GoogleDeveloperConnectConnectionGitlabConfigReadAuthorizerCredential `field:"required" json:"readAuthorizerCredential" yaml:"readAuthorizerCredential"`
	// Required.
	//
	// Immutable. SecretManager resource containing the webhook secret of a GitLab project,
	// formatted as 'projects/* /secrets/* /versions/*'. This is used to validate
	// webhooks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#webhook_secret_secret_version GoogleDeveloperConnectConnection#webhook_secret_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	WebhookSecretSecretVersion *string `field:"required" json:"webhookSecretSecretVersion" yaml:"webhookSecretSecretVersion"`
}

