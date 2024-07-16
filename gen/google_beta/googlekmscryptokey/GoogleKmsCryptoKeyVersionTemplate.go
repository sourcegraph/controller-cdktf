package googlekmscryptokey


type GoogleKmsCryptoKeyVersionTemplate struct {
	// The algorithm to use when creating a version based on this template. See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_crypto_key#algorithm GoogleKmsCryptoKey#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// The protection level to use when creating a version based on this template.
	//
	// Possible values include "SOFTWARE", "HSM", "EXTERNAL", "EXTERNAL_VPC". Defaults to "SOFTWARE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_crypto_key#protection_level GoogleKmsCryptoKey#protection_level}
	ProtectionLevel *string `field:"optional" json:"protectionLevel" yaml:"protectionLevel"`
}

