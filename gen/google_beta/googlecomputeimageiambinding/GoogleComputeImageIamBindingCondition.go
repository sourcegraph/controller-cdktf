package googlecomputeimageiambinding


type GoogleComputeImageIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_image_iam_binding#expression GoogleComputeImageIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_image_iam_binding#title GoogleComputeImageIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_compute_image_iam_binding#description GoogleComputeImageIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

