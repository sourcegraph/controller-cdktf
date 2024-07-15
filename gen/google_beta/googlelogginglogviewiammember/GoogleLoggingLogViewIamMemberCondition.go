package googlelogginglogviewiammember


type GoogleLoggingLogViewIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_log_view_iam_member#expression GoogleLoggingLogViewIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_log_view_iam_member#title GoogleLoggingLogViewIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_logging_log_view_iam_member#description GoogleLoggingLogViewIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

