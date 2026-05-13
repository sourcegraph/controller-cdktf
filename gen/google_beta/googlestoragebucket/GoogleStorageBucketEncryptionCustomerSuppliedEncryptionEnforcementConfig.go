package googlestoragebucket


type GoogleStorageBucketEncryptionCustomerSuppliedEncryptionEnforcementConfig struct {
	// Whether CSEK is restricted for new objects within the bucket.
	//
	// If FullyRestricted, new objects can't be created using CSEK encryption. If NotRestricted or unset, creation of new objects with CSEK encryption is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_storage_bucket#restriction_mode GoogleStorageBucket#restriction_mode}
	RestrictionMode *string `field:"required" json:"restrictionMode" yaml:"restrictionMode"`
}

