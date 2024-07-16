package googledataformrepository


type GoogleDataformRepositoryGitRemoteSettings struct {
	// The Git remote's default branch name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#default_branch GoogleDataformRepository#default_branch}
	DefaultBranch *string `field:"required" json:"defaultBranch" yaml:"defaultBranch"`
	// The Git remote's URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#url GoogleDataformRepository#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The name of the Secret Manager secret version to use as an authentication token for Git operations.
	//
	// This secret is for assigning with HTTPS only(for SSH use 'ssh_authentication_config'). Must be in the format projects/* /secrets/* /versions/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#authentication_token_secret_version GoogleDataformRepository#authentication_token_secret_version}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	AuthenticationTokenSecretVersion *string `field:"optional" json:"authenticationTokenSecretVersion" yaml:"authenticationTokenSecretVersion"`
	// ssh_authentication_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_dataform_repository#ssh_authentication_config GoogleDataformRepository#ssh_authentication_config}
	SshAuthenticationConfig *GoogleDataformRepositoryGitRemoteSettingsSshAuthenticationConfig `field:"optional" json:"sshAuthenticationConfig" yaml:"sshAuthenticationConfig"`
}

