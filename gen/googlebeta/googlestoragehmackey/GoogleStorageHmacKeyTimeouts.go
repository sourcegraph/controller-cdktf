package googlestoragehmackey


type GoogleStorageHmacKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_hmac_key#create GoogleStorageHmacKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_hmac_key#delete GoogleStorageHmacKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_hmac_key#update GoogleStorageHmacKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

