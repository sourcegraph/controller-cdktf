package googleserviceaccountiammember


type GoogleServiceAccountIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_account_iam_member#expression GoogleServiceAccountIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_account_iam_member#title GoogleServiceAccountIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_service_account_iam_member#description GoogleServiceAccountIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

