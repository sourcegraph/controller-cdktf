package cloudbuildtrigger


type CloudbuildTriggerDeveloperConnectEventConfig struct {
	// The Developer Connect Git repository link, formatted as 'projects/* /locations/* /connections/* /gitRepositoryLink/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudbuild_trigger#git_repository_link CloudbuildTrigger#git_repository_link}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	GitRepositoryLink *string `field:"required" json:"gitRepositoryLink" yaml:"gitRepositoryLink"`
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudbuild_trigger#pull_request CloudbuildTrigger#pull_request}
	PullRequest *CloudbuildTriggerDeveloperConnectEventConfigPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// push block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudbuild_trigger#push CloudbuildTrigger#push}
	Push *CloudbuildTriggerDeveloperConnectEventConfigPush `field:"optional" json:"push" yaml:"push"`
}

