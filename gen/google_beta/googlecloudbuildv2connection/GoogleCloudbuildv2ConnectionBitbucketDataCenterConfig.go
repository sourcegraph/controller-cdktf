package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionBitbucketDataCenterConfig struct {
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#authorizer_credential GoogleCloudbuildv2Connection#authorizer_credential}
	AuthorizerCredential *GoogleCloudbuildv2ConnectionBitbucketDataCenterConfigAuthorizerCredential `field:"required" json:"authorizerCredential" yaml:"authorizerCredential"`
	// The URI of the Bitbucket Data Center host this connection is for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#host_uri GoogleCloudbuildv2Connection#host_uri}
	HostUri *string `field:"required" json:"hostUri" yaml:"hostUri"`
	// read_authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#read_authorizer_credential GoogleCloudbuildv2Connection#read_authorizer_credential}
	ReadAuthorizerCredential *GoogleCloudbuildv2ConnectionBitbucketDataCenterConfigReadAuthorizerCredential `field:"required" json:"readAuthorizerCredential" yaml:"readAuthorizerCredential"`
	// Required. Immutable. SecretManager resource containing the webhook secret used to verify webhook events, formatted as 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#webhook_secret_secret_version GoogleCloudbuildv2Connection#webhook_secret_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	WebhookSecretSecretVersion *string `field:"required" json:"webhookSecretSecretVersion" yaml:"webhookSecretSecretVersion"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#service_directory_config GoogleCloudbuildv2Connection#service_directory_config}
	ServiceDirectoryConfig *GoogleCloudbuildv2ConnectionBitbucketDataCenterConfigServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// SSL certificate to use for requests to the Bitbucket Data Center.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#ssl_ca GoogleCloudbuildv2Connection#ssl_ca}
	SslCa *string `field:"optional" json:"sslCa" yaml:"sslCa"`
}

