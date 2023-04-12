package googlekmscryptokeyiammember


type GoogleKmsCryptoKeyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_crypto_key_iam_member#expression GoogleKmsCryptoKeyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_crypto_key_iam_member#title GoogleKmsCryptoKeyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_crypto_key_iam_member#description GoogleKmsCryptoKeyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

