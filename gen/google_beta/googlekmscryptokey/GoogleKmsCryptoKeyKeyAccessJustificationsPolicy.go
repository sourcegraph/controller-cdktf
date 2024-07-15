package googlekmscryptokey


type GoogleKmsCryptoKeyKeyAccessJustificationsPolicy struct {
	// The list of allowed reasons for access to this CryptoKey.
	//
	// Zero allowed
	// access reasons means all encrypt, decrypt, and sign operations for
	// this CryptoKey will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_kms_crypto_key#allowed_access_reasons GoogleKmsCryptoKey#allowed_access_reasons}
	AllowedAccessReasons *[]*string `field:"optional" json:"allowedAccessReasons" yaml:"allowedAccessReasons"`
}

