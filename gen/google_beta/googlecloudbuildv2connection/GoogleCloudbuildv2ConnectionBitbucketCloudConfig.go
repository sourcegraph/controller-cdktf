package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionBitbucketCloudConfig struct {
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#authorizer_credential GoogleCloudbuildv2Connection#authorizer_credential}
	AuthorizerCredential *GoogleCloudbuildv2ConnectionBitbucketCloudConfigAuthorizerCredential `field:"required" json:"authorizerCredential" yaml:"authorizerCredential"`
	// read_authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#read_authorizer_credential GoogleCloudbuildv2Connection#read_authorizer_credential}
	ReadAuthorizerCredential *GoogleCloudbuildv2ConnectionBitbucketCloudConfigReadAuthorizerCredential `field:"required" json:"readAuthorizerCredential" yaml:"readAuthorizerCredential"`
	// Required. Immutable. SecretManager resource containing the webhook secret used to verify webhook events, formatted as 'projects/* /secrets/* /versions/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#webhook_secret_secret_version GoogleCloudbuildv2Connection#webhook_secret_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	WebhookSecretSecretVersion *string `field:"required" json:"webhookSecretSecretVersion" yaml:"webhookSecretSecretVersion"`
	// The Bitbucket Cloud Workspace ID to be connected to Google Cloud Platform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#workspace GoogleCloudbuildv2Connection#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
}

