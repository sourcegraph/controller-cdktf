package googlenotebooksruntimeiammember


type GoogleNotebooksRuntimeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime_iam_member#expression GoogleNotebooksRuntimeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime_iam_member#title GoogleNotebooksRuntimeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime_iam_member#description GoogleNotebooksRuntimeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

