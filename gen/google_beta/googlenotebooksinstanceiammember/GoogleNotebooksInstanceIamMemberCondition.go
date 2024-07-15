package googlenotebooksinstanceiammember


type GoogleNotebooksInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance_iam_member#expression GoogleNotebooksInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance_iam_member#title GoogleNotebooksInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_notebooks_instance_iam_member#description GoogleNotebooksInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

