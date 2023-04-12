package googlecomputemachineimageiammember


type GoogleComputeMachineImageIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_machine_image_iam_member#expression GoogleComputeMachineImageIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_machine_image_iam_member#title GoogleComputeMachineImageIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_machine_image_iam_member#description GoogleComputeMachineImageIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

