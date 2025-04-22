package googledeveloperconnectconnection


type GoogleDeveloperConnectConnectionGithubConfig struct {
	// Required. Immutable. The GitHub Application that was installed to the GitHub user or organization. Possible values: GIT_HUB_APP_UNSPECIFIED DEVELOPER_CONNECT FIREBASE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#github_app GoogleDeveloperConnectConnection#github_app}
	GithubApp *string `field:"required" json:"githubApp" yaml:"githubApp"`
	// Optional. GitHub App installation id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#app_installation_id GoogleDeveloperConnectConnection#app_installation_id}
	AppInstallationId *string `field:"optional" json:"appInstallationId" yaml:"appInstallationId"`
	// authorizer_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_developer_connect_connection#authorizer_credential GoogleDeveloperConnectConnection#authorizer_credential}
	AuthorizerCredential *GoogleDeveloperConnectConnectionGithubConfigAuthorizerCredential `field:"optional" json:"authorizerCredential" yaml:"authorizerCredential"`
}

