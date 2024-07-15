package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionGitlabConfig struct {
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#authorizer_credential GoogleCloudbuildv2Connection#authorizer_credential}
	AuthorizerCredential *GoogleCloudbuildv2ConnectionGitlabConfigAuthorizerCredential `field:"required" json:"authorizerCredential" yaml:"authorizerCredential"`
	// read_authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#read_authorizer_credential GoogleCloudbuildv2Connection#read_authorizer_credential}
	ReadAuthorizerCredential *GoogleCloudbuildv2ConnectionGitlabConfigReadAuthorizerCredential `field:"required" json:"readAuthorizerCredential" yaml:"readAuthorizerCredential"`
	// Required. Immutable. SecretManager resource containing the webhook secret of a GitLab Enterprise project, formatted as 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#webhook_secret_secret_version GoogleCloudbuildv2Connection#webhook_secret_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	WebhookSecretSecretVersion *string `field:"required" json:"webhookSecretSecretVersion" yaml:"webhookSecretSecretVersion"`
	// The URI of the GitLab Enterprise host this connection is for. If not specified, the default value is https://gitlab.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#host_uri GoogleCloudbuildv2Connection#host_uri}
	HostUri *string `field:"optional" json:"hostUri" yaml:"hostUri"`
	// service_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#service_directory_config GoogleCloudbuildv2Connection#service_directory_config}
	ServiceDirectoryConfig *GoogleCloudbuildv2ConnectionGitlabConfigServiceDirectoryConfig `field:"optional" json:"serviceDirectoryConfig" yaml:"serviceDirectoryConfig"`
	// SSL certificate to use for requests to GitLab Enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#ssl_ca GoogleCloudbuildv2Connection#ssl_ca}
	SslCa *string `field:"optional" json:"sslCa" yaml:"sslCa"`
}

