package googledeveloperconnectconnection


type GoogleDeveloperConnectConnectionHttpConfigBearerTokenAuthentication struct {
	// The token SecretManager secret version to authenticate as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_connection#token_secret_version GoogleDeveloperConnectConnection#token_secret_version}
	TokenSecretVersion *string `field:"optional" json:"tokenSecretVersion" yaml:"tokenSecretVersion"`
}

