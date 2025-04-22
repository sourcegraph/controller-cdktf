package googledeveloperconnectconnection


type GoogleDeveloperConnectConnectionBitbucketCloudConfigAuthorizerCredential struct {
	// Required. A SecretManager resource containing the user token that authorizes the Developer Connect connection. Format: 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#user_token_secret_version GoogleDeveloperConnectConnection#user_token_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	UserTokenSecretVersion *string `field:"required" json:"userTokenSecretVersion" yaml:"userTokenSecretVersion"`
}

