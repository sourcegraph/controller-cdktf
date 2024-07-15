package googlenotebooksruntimeiambinding


type GoogleNotebooksRuntimeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime_iam_binding#expression GoogleNotebooksRuntimeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime_iam_binding#title GoogleNotebooksRuntimeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime_iam_binding#description GoogleNotebooksRuntimeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

