package googlecloudbuildbitbucketserverconfig


type GoogleCloudbuildBitbucketServerConfigConnectedRepositories struct {
	// Identifier for the project storing the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_bitbucket_server_config#project_key GoogleCloudbuildBitbucketServerConfig#project_key}
	ProjectKey *string `field:"required" json:"projectKey" yaml:"projectKey"`
	// Identifier for the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_bitbucket_server_config#repo_slug GoogleCloudbuildBitbucketServerConfig#repo_slug}
	RepoSlug *string `field:"required" json:"repoSlug" yaml:"repoSlug"`
}

