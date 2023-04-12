package googlecloudfunctionsfunctioniammember


type GoogleCloudfunctionsFunctionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudfunctions_function_iam_member#expression GoogleCloudfunctionsFunctionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudfunctions_function_iam_member#title GoogleCloudfunctionsFunctionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudfunctions_function_iam_member#description GoogleCloudfunctionsFunctionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

