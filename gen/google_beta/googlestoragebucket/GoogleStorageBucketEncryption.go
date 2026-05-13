package googlestoragebucket


type GoogleStorageBucketEncryption struct {
	// customer_managed_encryption_enforcement_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket#customer_managed_encryption_enforcement_config GoogleStorageBucket#customer_managed_encryption_enforcement_config}
	CustomerManagedEncryptionEnforcementConfig *GoogleStorageBucketEncryptionCustomerManagedEncryptionEnforcementConfig `field:"optional" json:"customerManagedEncryptionEnforcementConfig" yaml:"customerManagedEncryptionEnforcementConfig"`
	// customer_supplied_encryption_enforcement_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket#customer_supplied_encryption_enforcement_config GoogleStorageBucket#customer_supplied_encryption_enforcement_config}
	CustomerSuppliedEncryptionEnforcementConfig *GoogleStorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfig `field:"optional" json:"customerSuppliedEncryptionEnforcementConfig" yaml:"customerSuppliedEncryptionEnforcementConfig"`
	// A Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified.
	//
	// You must pay attention to whether the crypto key is available in the location that this bucket is created in. See the docs for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket#default_kms_key_name GoogleStorageBucket#default_kms_key_name}
	DefaultKmsKeyName *string `field:"optional" json:"defaultKmsKeyName" yaml:"defaultKmsKeyName"`
	// google_managed_encryption_enforcement_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket#google_managed_encryption_enforcement_config GoogleStorageBucket#google_managed_encryption_enforcement_config}
	GoogleManagedEncryptionEnforcementConfig *GoogleStorageBucketEncryptionGoogleManagedEncryptionEnforcementConfig `field:"optional" json:"googleManagedEncryptionEnforcementConfig" yaml:"googleManagedEncryptionEnforcementConfig"`
}

