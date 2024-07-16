package googlebillingaccountiammember


type GoogleBillingAccountIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_billing_account_iam_member#expression GoogleBillingAccountIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_billing_account_iam_member#title GoogleBillingAccountIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_billing_account_iam_member#description GoogleBillingAccountIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

