package googlesecretmanagersecretiammember


type GoogleSecretManagerSecretIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret_iam_member#expression GoogleSecretManagerSecretIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret_iam_member#title GoogleSecretManagerSecretIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_secret_manager_secret_iam_member#description GoogleSecretManagerSecretIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

