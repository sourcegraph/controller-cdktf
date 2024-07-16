package googleworkbenchinstanceiammember


type GoogleWorkbenchInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance_iam_member#expression GoogleWorkbenchInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance_iam_member#title GoogleWorkbenchInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_workbench_instance_iam_member#description GoogleWorkbenchInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

