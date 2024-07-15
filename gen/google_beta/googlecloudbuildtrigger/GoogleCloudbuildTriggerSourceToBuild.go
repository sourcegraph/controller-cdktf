package googlecloudbuildtrigger


type GoogleCloudbuildTriggerSourceToBuild struct {
	// The branch or tag to use. Must start with "refs/" (required).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#ref GoogleCloudbuildTrigger#ref}
	Ref *string `field:"required" json:"ref" yaml:"ref"`
	// The type of the repo, since it may not be explicit from the repo field (e.g from a URL). Values can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER Possible values: ["UNKNOWN", "CLOUD_SOURCE_REPOSITORIES", "GITHUB", "BITBUCKET_SERVER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#repo_type GoogleCloudbuildTrigger#repo_type}
	RepoType *string `field:"required" json:"repoType" yaml:"repoType"`
	// The full resource name of the bitbucket server config. Format: projects/{project}/locations/{location}/bitbucketServerConfigs/{id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#bitbucket_server_config GoogleCloudbuildTrigger#bitbucket_server_config}
	BitbucketServerConfig *string `field:"optional" json:"bitbucketServerConfig" yaml:"bitbucketServerConfig"`
	// The full resource name of the github enterprise config. Format: projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}. projects/{project}/githubEnterpriseConfigs/{id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#github_enterprise_config GoogleCloudbuildTrigger#github_enterprise_config}
	GithubEnterpriseConfig *string `field:"optional" json:"githubEnterpriseConfig" yaml:"githubEnterpriseConfig"`
	// The qualified resource name of the Repo API repository. Either uri or repository can be specified and is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#repository GoogleCloudbuildTrigger#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// The URI of the repo.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#uri GoogleCloudbuildTrigger#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

