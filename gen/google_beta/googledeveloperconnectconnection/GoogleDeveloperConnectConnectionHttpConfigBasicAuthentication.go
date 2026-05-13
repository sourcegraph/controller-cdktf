package googledeveloperconnectconnection


type GoogleDeveloperConnectConnectionHttpConfigBasicAuthentication struct {
	// The username to authenticate as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_connection#username GoogleDeveloperConnectConnection#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The password SecretManager secret version to authenticate as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_developer_connect_connection#password_secret_version GoogleDeveloperConnectConnection#password_secret_version}
	PasswordSecretVersion *string `field:"optional" json:"passwordSecretVersion" yaml:"passwordSecretVersion"`
}

