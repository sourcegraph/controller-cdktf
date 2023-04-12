package googlecloudfunctions2functioniammember


type GoogleCloudfunctions2FunctionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudfunctions2_function_iam_member#expression GoogleCloudfunctions2FunctionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudfunctions2_function_iam_member#title GoogleCloudfunctions2FunctionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudfunctions2_function_iam_member#description GoogleCloudfunctions2FunctionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

