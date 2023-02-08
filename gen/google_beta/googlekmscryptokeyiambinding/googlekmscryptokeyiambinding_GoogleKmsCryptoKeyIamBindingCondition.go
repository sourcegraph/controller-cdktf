package googlekmscryptokeyiambinding


type GoogleKmsCryptoKeyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_crypto_key_iam_binding#expression GoogleKmsCryptoKeyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_crypto_key_iam_binding#title GoogleKmsCryptoKeyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_crypto_key_iam_binding#description GoogleKmsCryptoKeyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

