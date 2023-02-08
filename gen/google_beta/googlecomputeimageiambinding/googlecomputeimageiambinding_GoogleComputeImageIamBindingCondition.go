package googlecomputeimageiambinding


type GoogleComputeImageIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_image_iam_binding#expression GoogleComputeImageIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_image_iam_binding#title GoogleComputeImageIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_image_iam_binding#description GoogleComputeImageIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

