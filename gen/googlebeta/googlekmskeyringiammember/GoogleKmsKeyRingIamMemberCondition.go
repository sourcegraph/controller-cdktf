package googlekmskeyringiammember


type GoogleKmsKeyRingIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_key_ring_iam_member#expression GoogleKmsKeyRingIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_key_ring_iam_member#title GoogleKmsKeyRingIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_key_ring_iam_member#description GoogleKmsKeyRingIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

