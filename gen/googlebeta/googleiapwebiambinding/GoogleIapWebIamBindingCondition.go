package googleiapwebiambinding


type GoogleIapWebIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_iam_binding#expression GoogleIapWebIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_iam_binding#title GoogleIapWebIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_web_iam_binding#description GoogleIapWebIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

