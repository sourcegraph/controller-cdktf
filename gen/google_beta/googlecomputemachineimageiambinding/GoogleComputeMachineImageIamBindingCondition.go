package googlecomputemachineimageiambinding


type GoogleComputeMachineImageIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image_iam_binding#expression GoogleComputeMachineImageIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image_iam_binding#title GoogleComputeMachineImageIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_machine_image_iam_binding#description GoogleComputeMachineImageIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

