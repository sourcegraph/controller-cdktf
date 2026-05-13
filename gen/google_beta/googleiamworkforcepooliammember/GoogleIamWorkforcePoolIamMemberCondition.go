package googleiamworkforcepooliammember


type GoogleIamWorkforcePoolIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_iam_member#expression GoogleIamWorkforcePoolIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_iam_member#title GoogleIamWorkforcePoolIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_iam_workforce_pool_iam_member#description GoogleIamWorkforcePoolIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

