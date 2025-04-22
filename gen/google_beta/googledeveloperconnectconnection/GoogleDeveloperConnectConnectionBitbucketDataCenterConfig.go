package googledeveloperconnectconnection


type GoogleDeveloperConnectConnectionBitbucketDataCenterConfig struct {
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#authorizer_credential GoogleDeveloperConnectConnection#authorizer_credential}
	AuthorizerCredential *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigAuthorizerCredential `field:"required" json:"authorizerCredential" yaml:"authorizerCredential"`
	// Required. The URI of the Bitbucket Data Center host this connection is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#host_uri GoogleDeveloperConnectConnection#host_uri}
	HostUri *string `field:"required" json:"hostUri" yaml:"hostUri"`
	// read_authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#read_authorizer_credential GoogleDeveloperConnectConnection#read_authorizer_credential}
	ReadAuthorizerCredential *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigReadAuthorizerCredential `field:"required" json:"readAuthorizerCredential" yaml:"readAuthorizerCredential"`
	// Required.
	//
	// Immutable. SecretManager resource containing the webhook secret used to verify webhook
	// events, formatted as 'projects/* /secrets/* /versions/*'. This is used to
	// validate webhooks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#webhook_secret_secret_version GoogleDeveloperConnectConnection#webhook_secret_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	WebhookSecretSecretVersion *string `field:"required" json:"webhookSecretSecretVersion" yaml:"webhookSecretSecretVersion"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#service_directory_config GoogleDeveloperConnectConnection#service_directory_config}
	ServiceDirectoryConfig *GoogleDeveloperConnectConnectionBitbucketDataCenterConfigServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// Optional. SSL certificate authority to trust when making requests to Bitbucket Data Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#ssl_ca_certificate GoogleDeveloperConnectConnection#ssl_ca_certificate}
	SslCaCertificate *string `field:"optional" json:"sslCaCertificate" yaml:"sslCaCertificate"`
}

