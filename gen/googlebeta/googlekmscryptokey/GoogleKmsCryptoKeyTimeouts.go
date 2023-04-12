package googlekmscryptokey


type GoogleKmsCryptoKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_crypto_key#create GoogleKmsCryptoKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_crypto_key#delete GoogleKmsCryptoKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_kms_crypto_key#update GoogleKmsCryptoKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

