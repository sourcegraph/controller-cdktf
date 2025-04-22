package googlegeminirepositorygroupiammember


type GoogleGeminiRepositoryGroupIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gemini_repository_group_iam_member#expression GoogleGeminiRepositoryGroupIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gemini_repository_group_iam_member#title GoogleGeminiRepositoryGroupIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_gemini_repository_group_iam_member#description GoogleGeminiRepositoryGroupIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

