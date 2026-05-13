package storagebucket


type StorageBucketEncryption struct {
	// customer_managed_encryption_enforcement_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_bucket#customer_managed_encryption_enforcement_config StorageBucket#customer_managed_encryption_enforcement_config}
	CustomerManagedEncryptionEnforcementConfig *StorageBucketEncryptionCustomerManagedEncryptionEnforcementConfig `field:"optional" json:"customerManagedEncryptionEnforcementConfig" yaml:"customerManagedEncryptionEnforcementConfig"`
	// customer_supplied_encryption_enforcement_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_bucket#customer_supplied_encryption_enforcement_config StorageBucket#customer_supplied_encryption_enforcement_config}
	CustomerSuppliedEncryptionEnforcementConfig *StorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfig `field:"optional" json:"customerSuppliedEncryptionEnforcementConfig" yaml:"customerSuppliedEncryptionEnforcementConfig"`
	// A Cloud KMS key that will be used to encrypt objects inserted into this bucket, if no encryption method is specified.
	//
	// You must pay attention to whether the crypto key is available in the location that this bucket is created in. See the docs for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_bucket#default_kms_key_name StorageBucket#default_kms_key_name}
	DefaultKmsKeyName *string `field:"optional" json:"defaultKmsKeyName" yaml:"defaultKmsKeyName"`
	// google_managed_encryption_enforcement_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/storage_bucket#google_managed_encryption_enforcement_config StorageBucket#google_managed_encryption_enforcement_config}
	GoogleManagedEncryptionEnforcementConfig *StorageBucketEncryptionGoogleManagedEncryptionEnforcementConfig `field:"optional" json:"googleManagedEncryptionEnforcementConfig" yaml:"googleManagedEncryptionEnforcementConfig"`
}

