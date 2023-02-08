package googleorganizationiammember


type GoogleOrganizationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_organization_iam_member#expression GoogleOrganizationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_organization_iam_member#title GoogleOrganizationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_organization_iam_member#description GoogleOrganizationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

