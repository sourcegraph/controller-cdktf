package cloudbuildtrigger


type CloudbuildTriggerDeveloperConnectEventConfigPush struct {
	// Regex of branches to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudbuild_trigger#branch CloudbuildTrigger#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// If true, only trigger a build if the revision regex does NOT match the git_ref regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudbuild_trigger#invert_regex CloudbuildTrigger#invert_regex}
	InvertRegex interface{} `field:"optional" json:"invertRegex" yaml:"invertRegex"`
	// Regex of tags to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudbuild_trigger#tag CloudbuildTrigger#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

