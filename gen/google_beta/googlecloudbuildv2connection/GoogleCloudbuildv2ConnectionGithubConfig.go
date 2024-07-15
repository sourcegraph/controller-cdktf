package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionGithubConfig struct {
	// GitHub App installation id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#app_installation_id GoogleCloudbuildv2Connection#app_installation_id}
	AppInstallationId *float64 `field:"optional" json:"appInstallationId" yaml:"appInstallationId"`
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuildv2_connection#authorizer_credential GoogleCloudbuildv2Connection#authorizer_credential}
	AuthorizerCredential *GoogleCloudbuildv2ConnectionGithubConfigAuthorizerCredential `field:"optional" json:"authorizerCredential" yaml:"authorizerCredential"`
}

