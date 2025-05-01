package googlegeminirepositorygroupiambinding


type GoogleGeminiRepositoryGroupIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group_iam_binding#expression GoogleGeminiRepositoryGroupIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group_iam_binding#title GoogleGeminiRepositoryGroupIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_gemini_repository_group_iam_binding#description GoogleGeminiRepositoryGroupIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

