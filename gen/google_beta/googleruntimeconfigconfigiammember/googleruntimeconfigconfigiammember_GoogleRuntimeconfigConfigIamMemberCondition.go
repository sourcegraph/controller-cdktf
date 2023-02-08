package googleruntimeconfigconfigiammember


type GoogleRuntimeconfigConfigIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_runtimeconfig_config_iam_member#expression GoogleRuntimeconfigConfigIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_runtimeconfig_config_iam_member#title GoogleRuntimeconfigConfigIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_runtimeconfig_config_iam_member#description GoogleRuntimeconfigConfigIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

