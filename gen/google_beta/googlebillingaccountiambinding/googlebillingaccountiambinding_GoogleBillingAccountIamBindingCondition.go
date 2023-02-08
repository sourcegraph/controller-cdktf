package googlebillingaccountiambinding


type GoogleBillingAccountIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_billing_account_iam_binding#expression GoogleBillingAccountIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_billing_account_iam_binding#title GoogleBillingAccountIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_billing_account_iam_binding#description GoogleBillingAccountIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

