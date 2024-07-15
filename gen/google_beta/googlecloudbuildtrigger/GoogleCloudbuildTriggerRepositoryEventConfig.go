package googlecloudbuildtrigger


type GoogleCloudbuildTriggerRepositoryEventConfig struct {
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#pull_request GoogleCloudbuildTrigger#pull_request}
	PullRequest *GoogleCloudbuildTriggerRepositoryEventConfigPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// push block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#push GoogleCloudbuildTrigger#push}
	Push *GoogleCloudbuildTriggerRepositoryEventConfigPush `field:"optional" json:"push" yaml:"push"`
	// The resource name of the Repo API resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#repository GoogleCloudbuildTrigger#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

