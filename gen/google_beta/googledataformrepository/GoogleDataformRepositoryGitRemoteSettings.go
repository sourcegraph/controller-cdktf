package googledataformrepository


type GoogleDataformRepositoryGitRemoteSettings struct {
	// The name of the Secret Manager secret version to use as an authentication token for Git operations.
	//
	// Must be in the format projects/*\/secrets/*\/versions/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataform_repository#authentication_token_secret_version GoogleDataformRepository#authentication_token_secret_version}
	AuthenticationTokenSecretVersion *string `field:"required" json:"authenticationTokenSecretVersion" yaml:"authenticationTokenSecretVersion"`
	// The Git remote's default branch name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataform_repository#default_branch GoogleDataformRepository#default_branch}
	DefaultBranch *string `field:"required" json:"defaultBranch" yaml:"defaultBranch"`
	// The Git remote's URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataform_repository#url GoogleDataformRepository#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

