package googlecomputeimageiammember


type GoogleComputeImageIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_image_iam_member#expression GoogleComputeImageIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_image_iam_member#title GoogleComputeImageIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_image_iam_member#description GoogleComputeImageIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

