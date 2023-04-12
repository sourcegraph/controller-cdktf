package googlenotebooksruntimeiambinding


type GoogleNotebooksRuntimeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime_iam_binding#expression GoogleNotebooksRuntimeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime_iam_binding#title GoogleNotebooksRuntimeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_runtime_iam_binding#description GoogleNotebooksRuntimeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

