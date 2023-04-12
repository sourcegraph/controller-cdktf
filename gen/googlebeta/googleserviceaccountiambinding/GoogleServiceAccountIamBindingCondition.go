package googleserviceaccountiambinding


type GoogleServiceAccountIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_account_iam_binding#expression GoogleServiceAccountIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_account_iam_binding#title GoogleServiceAccountIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_service_account_iam_binding#description GoogleServiceAccountIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

