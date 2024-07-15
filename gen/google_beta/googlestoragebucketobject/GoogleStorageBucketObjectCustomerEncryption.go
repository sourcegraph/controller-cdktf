package googlestoragebucketobject


type GoogleStorageBucketObjectCustomerEncryption struct {
	// Base64 encoded customer supplied encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket_object#encryption_key GoogleStorageBucketObject#encryption_key}
	EncryptionKey *string `field:"required" json:"encryptionKey" yaml:"encryptionKey"`
	// The encryption algorithm. Default: AES256.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket_object#encryption_algorithm GoogleStorageBucketObject#encryption_algorithm}
	EncryptionAlgorithm *string `field:"optional" json:"encryptionAlgorithm" yaml:"encryptionAlgorithm"`
}

