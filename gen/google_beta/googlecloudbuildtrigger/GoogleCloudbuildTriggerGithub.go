package googlecloudbuildtrigger


type GoogleCloudbuildTriggerGithub struct {
	// The resource name of the github enterprise config that should be applied to this installation. For example: "projects/{$projectId}/locations/{$locationId}/githubEnterpriseConfigs/{$configId}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#enterprise_config_resource_name GoogleCloudbuildTrigger#enterprise_config_resource_name}
	EnterpriseConfigResourceName *string `field:"optional" json:"enterpriseConfigResourceName" yaml:"enterpriseConfigResourceName"`
	// Name of the repository. For example: The name for https://github.com/googlecloudplatform/cloud-builders is "cloud-builders".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#name GoogleCloudbuildTrigger#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Owner of the repository. For example: The owner for https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#owner GoogleCloudbuildTrigger#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#pull_request GoogleCloudbuildTrigger#pull_request}
	PullRequest *GoogleCloudbuildTriggerGithubPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// push block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#push GoogleCloudbuildTrigger#push}
	Push *GoogleCloudbuildTriggerGithubPush `field:"optional" json:"push" yaml:"push"`
}

