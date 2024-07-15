package googlecloudbuildtrigger


type GoogleCloudbuildTriggerGithubPush struct {
	// Regex of branches to match.  Specify only one of branch or tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#branch GoogleCloudbuildTrigger#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// When true, only trigger a build if the revision regex does NOT match the git_ref regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#invert_regex GoogleCloudbuildTrigger#invert_regex}
	InvertRegex interface{} `field:"optional" json:"invertRegex" yaml:"invertRegex"`
	// Regex of tags to match.  Specify only one of branch or tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudbuild_trigger#tag GoogleCloudbuildTrigger#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

