package googlekmskeyringiambinding


type GoogleKmsKeyRingIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_key_ring_iam_binding#expression GoogleKmsKeyRingIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_key_ring_iam_binding#title GoogleKmsKeyRingIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_key_ring_iam_binding#description GoogleKmsKeyRingIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

