package cloudbuildtrigger


type CloudbuildTriggerDeveloperConnectEventConfigPullRequest struct {
	// Regex of branches to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudbuild_trigger#branch CloudbuildTrigger#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Configure builds to run whether a repository owner or collaborator need to comment '/gcbrun'. Possible values: ["COMMENTS_DISABLED", "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudbuild_trigger#comment_control CloudbuildTrigger#comment_control}
	CommentControl *string `field:"optional" json:"commentControl" yaml:"commentControl"`
	// If true, branches that do NOT match the git_ref will trigger a build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/cloudbuild_trigger#invert_regex CloudbuildTrigger#invert_regex}
	InvertRegex interface{} `field:"optional" json:"invertRegex" yaml:"invertRegex"`
}

