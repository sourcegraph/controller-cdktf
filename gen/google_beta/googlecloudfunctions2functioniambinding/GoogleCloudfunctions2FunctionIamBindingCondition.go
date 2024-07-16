package googlecloudfunctions2functioniambinding


type GoogleCloudfunctions2FunctionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function_iam_binding#expression GoogleCloudfunctions2FunctionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function_iam_binding#title GoogleCloudfunctions2FunctionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions2_function_iam_binding#description GoogleCloudfunctions2FunctionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

