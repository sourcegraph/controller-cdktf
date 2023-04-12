package googlecloudbuildtrigger


type GoogleCloudbuildTriggerGithub struct {
	// Name of the repository. For example: The name for https://github.com/googlecloudplatform/cloud-builders is "cloud-builders".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger#name GoogleCloudbuildTrigger#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Owner of the repository. For example: The owner for https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger#owner GoogleCloudbuildTrigger#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger#pull_request GoogleCloudbuildTrigger#pull_request}
	PullRequest *GoogleCloudbuildTriggerGithubPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// push block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger#push GoogleCloudbuildTrigger#push}
	Push *GoogleCloudbuildTriggerGithubPush `field:"optional" json:"push" yaml:"push"`
}

