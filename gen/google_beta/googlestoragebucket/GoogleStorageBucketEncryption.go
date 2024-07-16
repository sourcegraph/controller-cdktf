package googlestoragebucket


type GoogleStorageBucketEncryption struct {
	// A Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified.
	//
	// You must pay attention to whether the crypto key is available in the location that this bucket is created in. See the docs for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_bucket#default_kms_key_name GoogleStorageBucket#default_kms_key_name}
	DefaultKmsKeyName *string `field:"required" json:"defaultKmsKeyName" yaml:"defaultKmsKeyName"`
}

