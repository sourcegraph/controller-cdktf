package googlecloudbuildtrigger


type GoogleCloudbuildTriggerDeveloperConnectEventConfig struct {
	// The Developer Connect Git repository link, formatted as 'projects/* /locations/* /connections/* /gitRepositoryLink/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloudbuild_trigger#git_repository_link GoogleCloudbuildTrigger#git_repository_link}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	GitRepositoryLink *string `field:"required" json:"gitRepositoryLink" yaml:"gitRepositoryLink"`
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloudbuild_trigger#pull_request GoogleCloudbuildTrigger#pull_request}
	PullRequest *GoogleCloudbuildTriggerDeveloperConnectEventConfigPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// push block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_cloudbuild_trigger#push GoogleCloudbuildTrigger#push}
	Push *GoogleCloudbuildTriggerDeveloperConnectEventConfigPush `field:"optional" json:"push" yaml:"push"`
}

