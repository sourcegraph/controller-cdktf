package googlestoragebucket


type GoogleStorageBucketLifecycleRuleAction struct {
	// The type of the action of this Lifecycle Rule. Supported values include: Delete and SetStorageClass.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#type GoogleStorageBucket#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The target Storage Class of objects affected by this Lifecycle Rule. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#storage_class GoogleStorageBucket#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
}

