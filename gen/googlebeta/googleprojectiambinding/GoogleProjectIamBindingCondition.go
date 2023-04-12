package googleprojectiambinding


type GoogleProjectIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_project_iam_binding#expression GoogleProjectIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_project_iam_binding#title GoogleProjectIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_project_iam_binding#description GoogleProjectIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

