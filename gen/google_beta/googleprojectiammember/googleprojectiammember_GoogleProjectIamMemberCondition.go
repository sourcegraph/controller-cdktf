package googleprojectiammember


type GoogleProjectIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_project_iam_member#expression GoogleProjectIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_project_iam_member#title GoogleProjectIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_project_iam_member#description GoogleProjectIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

