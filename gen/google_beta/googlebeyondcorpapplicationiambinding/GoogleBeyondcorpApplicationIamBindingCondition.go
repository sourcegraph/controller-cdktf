package googlebeyondcorpapplicationiambinding


type GoogleBeyondcorpApplicationIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_application_iam_binding#expression GoogleBeyondcorpApplicationIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_application_iam_binding#title GoogleBeyondcorpApplicationIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_beyondcorp_application_iam_binding#description GoogleBeyondcorpApplicationIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

