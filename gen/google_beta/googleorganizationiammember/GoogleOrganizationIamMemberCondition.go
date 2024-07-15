package googleorganizationiammember


type GoogleOrganizationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_organization_iam_member#expression GoogleOrganizationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_organization_iam_member#title GoogleOrganizationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_organization_iam_member#description GoogleOrganizationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

