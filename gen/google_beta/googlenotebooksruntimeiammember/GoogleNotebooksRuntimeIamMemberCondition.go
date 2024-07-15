package googlenotebooksruntimeiammember


type GoogleNotebooksRuntimeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime_iam_member#expression GoogleNotebooksRuntimeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime_iam_member#title GoogleNotebooksRuntimeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_runtime_iam_member#description GoogleNotebooksRuntimeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

