package developerconnectconnection


type DeveloperConnectConnectionHttpConfigBearerTokenAuthentication struct {
	// The token SecretManager secret version to authenticate as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/developer_connect_connection#token_secret_version DeveloperConnectConnection#token_secret_version}
	TokenSecretVersion *string `field:"optional" json:"tokenSecretVersion" yaml:"tokenSecretVersion"`
}

