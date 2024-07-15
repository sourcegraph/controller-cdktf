package googlecloudfunctions2functioniammember


type GoogleCloudfunctions2FunctionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function_iam_member#expression GoogleCloudfunctions2FunctionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function_iam_member#title GoogleCloudfunctions2FunctionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function_iam_member#description GoogleCloudfunctions2FunctionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

