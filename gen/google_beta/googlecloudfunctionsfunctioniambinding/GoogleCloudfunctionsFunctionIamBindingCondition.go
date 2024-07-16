package googlecloudfunctionsfunctioniambinding


type GoogleCloudfunctionsFunctionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions_function_iam_binding#expression GoogleCloudfunctionsFunctionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions_function_iam_binding#title GoogleCloudfunctionsFunctionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_cloudfunctions_function_iam_binding#description GoogleCloudfunctionsFunctionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

