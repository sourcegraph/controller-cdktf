package googlecomputeinstanceiambinding


type GoogleComputeInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_iam_binding#expression GoogleComputeInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_iam_binding#title GoogleComputeInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_instance_iam_binding#description GoogleComputeInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

