package googleorganizationiambinding


type GoogleOrganizationIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_organization_iam_binding#expression GoogleOrganizationIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_organization_iam_binding#title GoogleOrganizationIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_organization_iam_binding#description GoogleOrganizationIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

